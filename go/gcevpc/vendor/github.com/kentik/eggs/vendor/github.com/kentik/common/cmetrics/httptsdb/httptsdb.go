package httptsdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kentik/go-metrics"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	MIN_FOR_HOST            = 6
	MAX_SEND_TRIES          = 2
	CLIENT_RESPONSE_TIMEOUT = 5 * time.Second
	CLIENT_KEEP_ALIVE       = 60 * time.Second
	CLIENT_TLS_TIMEOUT      = 5 * time.Second
	ContentType             = "application/json"
	API_EMAIL_HEADER        = "X-CH-Auth-Email"
	API_PASSWORD_HEADER     = "X-CH-Auth-API-Token"
)

var shortHostName string = ""

type TSDBMetric struct {
	Metric    string            `json:"metric"`
	Timestamp int64             `json:"timestamp"`
	Value     int64             `json:"value"`
	Tags      map[string]string `json:"tags"`
}

// OpenTSDBConfig provides a container with configuration parameters for
// the OpenTSDB exporter
type OpenTSDBConfig struct {
	Addr               string            // Network address to connect to
	Registry           metrics.Registry  // Registry to be exported
	FlushInterval      time.Duration     // Flush interval
	DurationUnit       time.Duration     // Time conversion unit for durations
	Prefix             string            // Prefix to be prepended to metric names
	Debug              bool              // write to stdout for debug
	Tags               map[string]string // add these tags to each metric writen
	Send               chan []byte       // manage # of outstanding http requests here.
	MaxHttpOutstanding int
	ProxyUrl           string
	Extra              map[string]string
	ApiEmail           *string
	ApiPassword        *string
}

// OpenTSDB is a blocking exporter function which reports metrics in r
// to a TSDB server located at addr, flushing them every d duration
// and prepending metric names with prefix.
func OpenTSDB(r metrics.Registry, d time.Duration, prefix string, addr string, maxOutstanding int) {
	OpenTSDBWithConfig(OpenTSDBConfig{
		Addr:               addr,
		Registry:           r,
		FlushInterval:      d,
		DurationUnit:       time.Nanosecond,
		Prefix:             prefix,
		Debug:              false,
		MaxHttpOutstanding: maxOutstanding,
		Send:               make(chan []byte, maxOutstanding),
		Tags:               map[string]string{},
		Extra:              nil,
	})
}

// OpenTSDBWithConfig is a blocking exporter function just like OpenTSDB,
// but it takes a OpenTSDBConfig instead.
func OpenTSDBWithConfig(c OpenTSDBConfig) {
	go c.runSend()

	for _ = range time.Tick(c.FlushInterval) {
		if err := openTSDB(&c); nil != err {
			log.Println(err)
		}
	}
}

func getShortHostname() string {
	if shortHostName == "" {
		host, _ := os.Hostname()
		strings.Replace(host, ".", "_", -1)
		shortHostName = host
	}
	return shortHostName
}

func addTypeTag(in map[string]string, mtype string) map[string]string {
	out := make(map[string]string)

	// Copy these over as a base.
	for k, v := range in {
		out[k] = v
	}

	// Add in type, and send
	out["type"] = mtype
	return out
}

func (c *OpenTSDBConfig) runSend() {
	tr := &http.Transport{
		DisableCompression: false,
		DisableKeepAlives:  false,
		Dial: (&net.Dialer{
			Timeout:   CLIENT_RESPONSE_TIMEOUT,
			KeepAlive: CLIENT_KEEP_ALIVE,
		}).Dial,
		TLSHandshakeTimeout: CLIENT_TLS_TIMEOUT,
	}

	// Add a proxy if needed.
	if c.ProxyUrl != "" {
		proxyUrl, err := url.Parse(c.ProxyUrl)
		if err != nil {
			fmt.Printf("Error setting proxy: %v\n", err)
		} else {
			tr.Proxy = http.ProxyURL(proxyUrl)
			fmt.Printf("Set outbound proxy: %s\n", c.ProxyUrl)
		}
	}

	client := &http.Client{Transport: tr, Timeout: CLIENT_RESPONSE_TIMEOUT}

	for r := range c.Send {
		if c.Debug {
			fmt.Printf("Metrics: %v", string(r))
		} else {
			for i := 0; i < MAX_SEND_TRIES; i++ {
				req, err := http.NewRequest("POST", c.Addr, bytes.NewBuffer(r))
				if err != nil {
					fmt.Printf("Error Creating Request: %v\n", err)
					continue
				}

				req.Header.Add("Content-Type", ContentType)

				if c.ApiEmail != nil && c.ApiPassword != nil {
					req.Header.Add(API_EMAIL_HEADER, *c.ApiEmail)
					req.Header.Add(API_PASSWORD_HEADER, *c.ApiPassword)
				}

				resp, err := client.Do(req)
				if err != nil {
					if i > 0 {
						fmt.Printf("Error Posting to %s: %v\n", c.Addr, err)
					} else {
						fmt.Printf("Retry Posting to %s: %v\n", c.Addr, err)
					}
					client = &http.Client{Transport: tr, Timeout: CLIENT_RESPONSE_TIMEOUT}
				} else {
					// Fire and forget
					io.Copy(ioutil.Discard, resp.Body)
					resp.Body.Close()
					break
				}
			}
		}
	}
}

/**
Write out additional tags
*/
func openTSDB(c *OpenTSDBConfig) error {

	shortHostnameBase := getShortHostname()
	now := time.Now().Unix()
	sendBody := make([]TSDBMetric, 0)
	du := float64(c.DurationUnit)

	c.Registry.Each(func(baseName string, i interface{}) {

		pts := strings.Split(baseName, ".")
		name := pts[0]
		tags := make(map[string]string)

		// Copy these over as a base.
		for k, v := range c.Tags {
			tags[k] = v
		}

		tags["host"] = shortHostnameBase

		// This is all kind of a hack, but currently, chfserver registers its per-device metrics
		// with the names "server_<metric-name>.chfserver.<fqdn>.1.<cid>.<device-name>.<did>.<sid>",
		// and so hits the first block, below.  chfclient registers all its metrics as
		// "client_<metric>.<cid>.<device-name>.<did>, and needs the second. ("ft"/"dt"/"level"
		// are all set globally for chfclient, so don't need to be packed in here.)
		// Per-device proxy metrics are sent as "proxy_metric.<cid>.<flow-type>.<did>",
		// and we need to extract the <flow_type> into "ft", since it can't be set globally.
		if len(pts) > MIN_FOR_HOST {
			pLen := len(pts)
			tags["cid"] = pts[pLen-4]
			tags["did"] = pts[pLen-2]
			tags["sid"] = pts[pLen-1]
		} else {
			if len(pts) >= 4 {
				tags["cid"] = pts[1]
				tags["did"] = pts[3]
				if strings.HasPrefix(name, "proxy_") {
					tags["ft"] = pts[2]
				}
			}
		}

		if c.Extra != nil {
			for k, v := range c.Extra {
				tags[k] = v
			}
		}

		if c.Prefix != "" {
			name = c.Prefix + "." + name
		}

		switch metric := i.(type) {
		case metrics.Counter:
			sendBody = append(sendBody, TSDBMetric{Metric: name, Timestamp: now, Value: metric.Count(), Tags: addTypeTag(tags, "count")})
		case metrics.Gauge:
			sendBody = append(sendBody, TSDBMetric{Metric: name, Timestamp: now, Value: metric.Value(), Tags: addTypeTag(tags, "value")})
		case metrics.Histogram:
			h := metric.Snapshot()
			ps := h.Percentiles([]float64{0.5, 0.75, 0.95, 0.99, 0.999})
			sendBody = append(sendBody, TSDBMetric{Metric: name, Timestamp: now, Value: h.Count(), Tags: addTypeTag(tags, "count")})
			sendBody = append(sendBody, TSDBMetric{Metric: name, Timestamp: now, Value: h.Min(), Tags: addTypeTag(tags, "min")})
			sendBody = append(sendBody, TSDBMetric{Metric: name, Timestamp: now, Value: h.Max(), Tags: addTypeTag(tags, "max")})
			sendBody = append(sendBody, TSDBMetric{Metric: name, Timestamp: now, Value: int64(h.Mean()), Tags: addTypeTag(tags, "mean")})
			sendBody = append(sendBody, TSDBMetric{Metric: name, Timestamp: now, Value: int64(ps[2]), Tags: addTypeTag(tags, "95-percentile")})
			sendBody = append(sendBody, TSDBMetric{Metric: name, Timestamp: now, Value: int64(ps[3]), Tags: addTypeTag(tags, "99-percentile")})
			metric.Clear()
		case metrics.Meter:
			m := metric.Snapshot()
			sendBody = append(sendBody, TSDBMetric{Metric: name, Timestamp: now, Value: m.Count(), Tags: addTypeTag(tags, "count")})
			sendBody = append(sendBody, TSDBMetric{Metric: name, Timestamp: now, Value: int64(m.Rate1()), Tags: addTypeTag(tags, "one-minute")})
		case metrics.Timer:
			t := metric.Snapshot()
			ps := t.Percentiles([]float64{0.5, 0.75, 0.95, 0.99, 0.999})
			sendBody = append(sendBody, TSDBMetric{Metric: name, Timestamp: now, Value: t.Count(), Tags: addTypeTag(tags, "count")})
			sendBody = append(sendBody, TSDBMetric{Metric: name, Timestamp: now, Value: t.Min() / int64(du), Tags: addTypeTag(tags, "min")})
			sendBody = append(sendBody, TSDBMetric{Metric: name, Timestamp: now, Value: t.Max() / int64(du), Tags: addTypeTag(tags, "max")})
			sendBody = append(sendBody, TSDBMetric{Metric: name, Timestamp: now, Value: int64(t.Mean() / du), Tags: addTypeTag(tags, "mean")})
			sendBody = append(sendBody, TSDBMetric{Metric: name, Timestamp: now, Value: int64(ps[2] / du), Tags: addTypeTag(tags, "95-percentile")})
			sendBody = append(sendBody, TSDBMetric{Metric: name, Timestamp: now, Value: int64(ps[3] / du), Tags: addTypeTag(tags, "99-percentile")})
			sendBody = append(sendBody, TSDBMetric{Metric: name, Timestamp: now, Value: int64(t.Rate1()), Tags: addTypeTag(tags, "one-minute")})
			sendBody = append(sendBody, TSDBMetric{Metric: name, Timestamp: now, Value: int64(t.Rate5()), Tags: addTypeTag(tags, "five-minute")})
			sendBody = append(sendBody, TSDBMetric{Metric: name, Timestamp: now, Value: int64(t.Rate15()), Tags: addTypeTag(tags, "fifteen-minute")})
			metric.Clear()
		}
	})

	sendBodyPruned := make([]TSDBMetric, 0)
	for _, m := range sendBody {
		if m.Value > 0 {
			sendBodyPruned = append(sendBodyPruned, m)
		}
	}

	if len(sendBodyPruned) > 0 {
		if ebytes, err := json.Marshal(sendBodyPruned); err != nil {
			fmt.Printf("Error encoding json: %v\n", err)
		} else {
			if len(c.Send) < c.MaxHttpOutstanding {
				c.Send <- ebytes
			} else {
				fmt.Printf("Dropping flow: Q at %d\n", len(c.Send))
			}
		}
	}

	return nil
}
