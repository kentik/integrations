package influx

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/kentik/go-metrics"
)

const (
	MIN_FOR_HOST            = 6
	MAX_SEND_TRIES          = 2
	SEND_SLEEP              = 1 * time.Second
	CLIENT_RESPONSE_TIMEOUT = 5 * time.Second
	CLIENT_KEEP_ALIVE       = 60 * time.Second
	CLIENT_TLS_TIMEOUT      = 5 * time.Second
	ContentType             = "application/x-www-form-urlencoded"
)

var shortHostName string = ""

type INFLUXMetricSet struct {
	metrics []*INFLUXMetric
}

type INFLUXMetric struct {
	Metric    string
	Timestamp int64
	Value     int64
	Type      string
	Tags      map[string]string
}

// OpenINFLUXConfig provides a container with configuration parameters for
// the OpenINFLUX exporter
type OpenINFLUXConfig struct {
	Addr               string                // Network address to connect to
	Registry           metrics.Registry      // Registry to be exported
	FlushInterval      time.Duration         // Flush interval
	DurationUnit       time.Duration         // Time conversion unit for durations
	Prefix             string                // Prefix to be prepended to metric names
	Debug              bool                  // write to stdout for debug
	Quiet              bool                  // silence all comms
	Tags               map[string]string     // add these tags to each metric writen
	Send               chan *INFLUXMetricSet // manage # of outstanding http requests here.
	MaxHttpOutstanding int
	ProxyUrl           string
	Extra              map[string]string
}

// OpenINFLUX is a blocking exporter function which reports metrics in r
// to a INFLUX server located at addr, flushing them every d duration
// and prepending metric names with prefix.
func OpenINFLUX(r metrics.Registry, d time.Duration, prefix string, addr string, maxOutstanding int) {
	OpenINFLUXWithConfig(OpenINFLUXConfig{
		Addr:               addr,
		Registry:           r,
		FlushInterval:      d,
		DurationUnit:       time.Nanosecond,
		Prefix:             prefix,
		Debug:              false,
		MaxHttpOutstanding: maxOutstanding,
		Send:               make(chan *INFLUXMetricSet, maxOutstanding),
		Tags:               map[string]string{},
		Extra:              nil,
	})
}

// OpenINFLUXWithConfig is a blocking exporter function just like OpenINFLUX,
// but it takes a OpenINFLUXConfig instead.
func OpenINFLUXWithConfig(c OpenINFLUXConfig) {
	go c.runSend()

	for _ = range time.Tick(c.FlushInterval) {
		if err := openINFLUX(&c); nil != err {
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

func addTypeTag(tags map[string]string, mtype string) map[string]string {
	tags["type"] = mtype
	return tags
}

func (c *OpenINFLUXConfig) runSend() {
	if strings.HasPrefix(c.Addr, "http") {
		c.runSendViaHTTP()
	} else if strings.HasPrefix(c.Addr, "tcp") || strings.HasPrefix(c.Addr, "udp") {
		c.runSendViaSocket()
	}
}

func (c *OpenINFLUXConfig) runSendViaSocket() {
	for r := range c.Send {
		var w *bufio.Writer
		var conn net.Conn = nil

		if c.Debug {
			w = bufio.NewWriter(os.Stdout)
		} else {
			var err error
			pts := strings.Split(c.Addr, "://")
			if len(pts) == 2 {
				conn, err = net.Dial(pts[0], pts[1])
				if nil != err {
					if !c.Quiet {
						fmt.Printf("Invalid metrics address: %s, %v\n", c.Addr, err)
					}
					continue
				}
				w = bufio.NewWriter(conn)
			} else {
				if !c.Quiet {
					fmt.Printf("Invalid metrics address: %s\n", c.Addr)
				}
				continue
			}
		}

		if ebytes, err := r.ToWire(); err != nil {
			if !c.Quiet {
				fmt.Printf("Error encoding to wire format: %v\n", err)
			}
			continue
		} else {
			w.Write(ebytes)
			w.Flush()
		}

		if !c.Debug {
			conn.Close()
		}
	}
}

func (c *OpenINFLUXConfig) runSendViaHTTP() {
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
			if !c.Quiet {
				fmt.Printf("Error setting proxy: %v\n", err)
			}
		} else {
			tr.Proxy = http.ProxyURL(proxyUrl)
			if !c.Quiet {
				fmt.Printf("Set outbound proxy: %s\n", c.ProxyUrl)
			}
		}
	}

	client := &http.Client{Transport: tr, Timeout: CLIENT_RESPONSE_TIMEOUT}

	for r := range c.Send {
		if ebytes, err := r.ToWire(); err != nil {
			if !c.Quiet {
				fmt.Printf("Error encoding to wire format: %v\n", err)
			}
			continue
		} else {
			if c.Debug {
				fmt.Printf("Metrics: %v", string(ebytes))
			} else {
				req, err := http.NewRequest("POST", c.Addr, bytes.NewBuffer(ebytes))
				if err != nil {
					if !c.Quiet {
						fmt.Printf("Error Creating Request: %v\n", err)
					}
					continue
				}
				req.Header.Add("Content-Type", ContentType)

				success := false
				for i := 0; i < MAX_SEND_TRIES; i++ {
					resp, err := client.Do(req)
					if err != nil {
						time.Sleep(SEND_SLEEP)
						client = &http.Client{Transport: tr, Timeout: CLIENT_RESPONSE_TIMEOUT}
					} else {
						// FIXME: check response code here and retry on non-200?
						success = true
						io.Copy(ioutil.Discard, resp.Body)
						resp.Body.Close()
						break
					}
				}

				if !success && !c.Quiet {
					fmt.Printf("Error Posting to %s: %v\n", c.Addr, err)
				}
			}
		}
	}
}

/**
Write out additional tags
*/
func openINFLUX(c *OpenINFLUXConfig) error {

	shortHostnameBase := getShortHostname()
	now := time.Now().UnixNano()
	sendBody := NewINFLUXMetricSet()
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

		nPts := strings.Split(name, "^")
		if len(nPts) > 1 {
			name = nPts[0]
			for _, np := range nPts[1:] {
				npr := strings.Split(np, "=")
				// e.g. the name was "mymetric^$CID=1234" and
				// the tags array should have "cid=$CID".
				// Don't use this. Prefer the branch below.
				if npr[0][0] == '$' {
					for k, v := range tags {
						if v == npr[0] {
							tags[k] = npr[1]
						}
					}
				} else {
					// e.g. the name was "mymetric^sometag=somevalue".
					// We will send a tag sometag=somevalue.
					tags[npr[0]] = npr[1]
				}
			}
		}

		if c.Prefix != "" {
			name = c.Prefix + "." + name
		}

		switch metric := i.(type) {
		case metrics.Counter:
			sendBody.Add(&INFLUXMetric{Metric: name, Timestamp: now, Value: metric.Count(), Tags: tags, Type: "count"})
		case metrics.Gauge:
			sendBody.Add(&INFLUXMetric{Metric: name, Timestamp: now, Value: metric.Value(), Tags: tags, Type: "value"})
		case metrics.Histogram:
			h := metric.Snapshot()
			ps := h.Percentiles([]float64{0.5, 0.75, 0.95, 0.99, 0.999})
			sendBody.Add(&INFLUXMetric{Metric: name, Timestamp: now, Value: h.Count(), Tags: tags, Type: "count"})
			sendBody.Add(&INFLUXMetric{Metric: name, Timestamp: now, Value: h.Min(), Tags: tags, Type: "min"})
			sendBody.Add(&INFLUXMetric{Metric: name, Timestamp: now, Value: h.Max(), Tags: tags, Type: "max"})
			sendBody.Add(&INFLUXMetric{Metric: name, Timestamp: now, Value: int64(h.Mean()), Tags: tags, Type: "mean"})
			sendBody.Add(&INFLUXMetric{Metric: name, Timestamp: now, Value: int64(ps[2]), Tags: tags, Type: "95-percentile"})
			sendBody.Add(&INFLUXMetric{Metric: name, Timestamp: now, Value: int64(ps[3]), Tags: tags, Type: "99-percentile"})
			metric.Clear()
		case metrics.Meter:
			m := metric.Snapshot()
			sendBody.Add(&INFLUXMetric{Metric: name, Timestamp: now, Value: m.Count(), Tags: tags, Type: "count"})
			sendBody.Add(&INFLUXMetric{Metric: name, Timestamp: now, Value: int64(m.Rate1()), Tags: tags, Type: "one-minute"})
		case metrics.Timer:
			t := metric.Snapshot()
			ps := t.Percentiles([]float64{0.5, 0.75, 0.95, 0.99, 0.999})
			sendBody.Add(&INFLUXMetric{Metric: name, Timestamp: now, Value: t.Count(), Tags: tags, Type: "count"})
			sendBody.Add(&INFLUXMetric{Metric: name, Timestamp: now, Value: t.Min() / int64(du), Tags: tags, Type: "min"})
			sendBody.Add(&INFLUXMetric{Metric: name, Timestamp: now, Value: t.Max() / int64(du), Tags: tags, Type: "max"})
			sendBody.Add(&INFLUXMetric{Metric: name, Timestamp: now, Value: int64(t.Mean() / du), Tags: tags, Type: "mean"})
			sendBody.Add(&INFLUXMetric{Metric: name, Timestamp: now, Value: int64(ps[2] / du), Tags: tags, Type: "95-percentile"})
			sendBody.Add(&INFLUXMetric{Metric: name, Timestamp: now, Value: int64(ps[3] / du), Tags: tags, Type: "99-percentile"})
			sendBody.Add(&INFLUXMetric{Metric: name, Timestamp: now, Value: int64(t.Rate1()), Tags: tags, Type: "one-minute"})
			sendBody.Add(&INFLUXMetric{Metric: name, Timestamp: now, Value: int64(t.Rate5()), Tags: tags, Type: "five-minute"})
			sendBody.Add(&INFLUXMetric{Metric: name, Timestamp: now, Value: int64(t.Rate15()), Tags: tags, Type: "fifteen-minute"})
			metric.Clear()
		}
	})

	if sendBody.Len() > 0 {
		if len(c.Send) < c.MaxHttpOutstanding {
			c.Send <- sendBody
		} else {
			if !c.Quiet {
				fmt.Printf("Dropping flow: Q at %d\n", len(c.Send))
			}
		}
	}

	return nil
}

func (m *INFLUXMetric) ToWire() []byte {
	tags := make([]string, (len(m.Tags))+2)
	tags[0] = m.Metric
	tags[1] = "type=" + m.Type
	i := 2
	for k, v := range m.Tags {
		tags[i] = k + "=" + v
		i++
	}

	return []byte(fmt.Sprintf("%s value=%d %d\n", strings.Join(tags, ","), m.Value, m.Timestamp))
}

func (m *INFLUXMetricSet) Len() int {
	return len(m.metrics)
}

func (m *INFLUXMetricSet) ToWire() ([]byte, error) {
	var buf bytes.Buffer
	for _, met := range m.metrics {
		buf.Write(met.ToWire())
	}
	return buf.Bytes(), nil
}

func (m *INFLUXMetricSet) Add(met *INFLUXMetric) {
	if met.Value > 0 {
		m.metrics = append(m.metrics, met)
	}
}

func NewINFLUXMetricSet() *INFLUXMetricSet {
	return &INFLUXMetricSet{
		metrics: make([]*INFLUXMetric, 0),
	}
}
