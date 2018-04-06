package cmetrics

import (
	"log"
	"log/syslog"
	"net"
	"os"
	"strings"
	"time"

	"github.com/kentik/common/cmetrics/httptsdb"
	"github.com/kentik/common/cmetrics/influx"
	"github.com/kentik/common/cmetrics/tsdb"
	metrics "github.com/kentik/go-metrics"
)

var (
	SYSLOG_FILE_PATH    = "/dev/log"
	MAX_HTTP_REQ        = 3 // # in-flight metric calls
	CH_HTTP_LOCAL_PROXY = "CH_HTTP_LOCAL_PROXY"
)

// Logger abstracts away the logging implementation
type Logger interface {
	Debugf(prefix, format string, v ...interface{})
	Infof(prefix, format string, v ...interface{})
	Errorf(prefix, format string, v ...interface{})
	Warnf(prefix, format string, v ...interface{})
}

func SetConf(conf string, l Logger, log_prefix string, tsdb_prefix string, tags []string, extra []string, apiEmail *string, apiPassword *string) {

	l.Infof(log_prefix, "Setting metrics: %s", conf)

	if conf != "none" {
		switch conf {
		case "syslog":
			if w, err := syslog.New(syslog.LOG_INFO, "metrics"); err == nil && w != nil {
				go metrics.Syslog(metrics.DefaultRegistry, 60e9, w)
			} else {
				l.Errorf(log_prefix, "Could not start syslog metrics: %v", err)
			}
		case "stderr":
			go metrics.Log(metrics.DefaultRegistry, 60e9, log.New(os.Stderr, "metrics: ", log.Lmicroseconds))
		default:
			dest := strings.SplitN(conf, ":", 2)
			switch dest[0] {
			case "graphite":
				l.Infof(log_prefix, "Metrics: Connecting to graphite: %s", dest[1])
				addr, _ := net.ResolveTCPAddr("tcp", dest[1])
				go metrics.Graphite(metrics.DefaultRegistry, 10e9, "metrics", addr)
			case "tsdb", "tsdb_debug":
				flushTime := 60 * time.Second
				if dest[0] == "tsdb_debug" {
					flushTime = 30 * time.Second
				}

				if strings.HasPrefix(dest[1], "http") {
					l.Infof(log_prefix, "Metrics: Connecting to [%s]: %s. [HTTP]", dest[0], dest[1])
					tagsMap := make(map[string]string)
					for _, t := range tags {
						// TODO(stefan): error out on bad tag names, or fix them
						// http://opentsdb.net/docs/build/html/user_guide/writing/index.html
						// Only the following characters are allowed: a to z, A to Z, 0 to 9, -, _, ., / or Unicode letters (as per the specification)
						pts := strings.SplitN(t, "=", 2)
						if len(pts) > 1 {
							tagsMap[pts[0]] = pts[1]
						}
					}

					extraMap := make(map[string]string)
					for _, pt := range extra {
						pts := strings.SplitN(pt, "=", 2)
						extraMap[pts[0]] = pts[1]
					}

					go httptsdb.OpenTSDBWithConfig(httptsdb.OpenTSDBConfig{
						Addr:               dest[1],
						Registry:           metrics.DefaultRegistry,
						FlushInterval:      flushTime,
						DurationUnit:       time.Millisecond,
						Prefix:             tsdb_prefix,
						Debug:              (dest[0] == "tsdb_debug"),
						Tags:               tagsMap,
						Send:               make(chan []byte, MAX_HTTP_REQ),
						ProxyUrl:           os.Getenv(CH_HTTP_LOCAL_PROXY),
						MaxHttpOutstanding: MAX_HTTP_REQ,
						Extra:              extraMap,
						ApiEmail:           apiEmail,
						ApiPassword:        apiPassword,
					})
				} else {
					tagsSer := ""
					if tags != nil {
						tagsSer = strings.Join(tags, " ")
					}

					l.Infof(log_prefix, "Metrics: Connecting to [%s]: %s. [TCP]. Debug=%v", dest[0], dest[1], (dest[0] == "tsdb_debug"))
					addr, err := net.ResolveTCPAddr("tcp", dest[1])
					if err != nil {
						l.Errorf(log_prefix, "Could not resolve address: %s %v", dest[1], err)
					} else {
						go tsdb.OpenTSDBWithConfig(tsdb.OpenTSDBConfig{
							Addr:          addr,
							Registry:      metrics.DefaultRegistry,
							FlushInterval: flushTime,
							DurationUnit:  time.Millisecond,
							Prefix:        tsdb_prefix,
							Debug:         (dest[0] == "tsdb_debug"),
							Tags:          tagsSer,
							Extra:         extra,
						})
					}
				}
			case "influx", "influx_debug", "influx_quiet":
				flushTime := 60 * time.Second
				if dest[0] == "influx_debug" {
					flushTime = 30 * time.Second
				}

				if strings.HasPrefix(dest[1], "http") || strings.HasPrefix(dest[1], "tcp") || strings.HasPrefix(dest[1], "udp") {
					l.Infof(log_prefix, "Metrics: Connecting Influx to [%s]: %s. [HTTP|TCP|UDP]", dest[0], dest[1])
					tagsMap := make(map[string]string)
					for _, t := range tags {
						pts := strings.SplitN(t, "=", 2)
						if len(pts) > 1 {
							tagsMap[pts[0]] = pts[1]
						}
					}

					extraMap := make(map[string]string)
					for _, pt := range extra {
						pts := strings.SplitN(pt, "=", 2)
						extraMap[pts[0]] = pts[1]
					}

					go influx.OpenINFLUXWithConfig(influx.OpenINFLUXConfig{
						Addr:               dest[1],
						Registry:           metrics.DefaultRegistry,
						FlushInterval:      flushTime,
						DurationUnit:       time.Millisecond,
						Prefix:             tsdb_prefix,
						Debug:              (dest[0] == "influx_debug"),
						Quiet:              (dest[0] == "influx_quiet"),
						Tags:               tagsMap,
						Send:               make(chan *influx.INFLUXMetricSet, MAX_HTTP_REQ),
						ProxyUrl:           os.Getenv(CH_HTTP_LOCAL_PROXY),
						MaxHttpOutstanding: MAX_HTTP_REQ,
						Extra:              extraMap,
					})
				} else {
					l.Errorf(log_prefix, "Only HTTP|TCP|UDP endpoint for influx currently supported")
				}
			}
		}
	}
}
