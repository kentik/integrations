package metrics

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/kentik/common/cmetrics/httptsdb"
	"github.com/kentik/go-metrics"
)

const (
	MaxHttpRequests    = 3
	MetricsSampleSize  = 1028
	MetricsSampleAlpha = 0.015
)

type Metrics struct {
	TotalFlowsIn   metrics.Meter
	TotalFlowsOut  metrics.Meter
	OrigSampleRate metrics.Histogram
	NewSampleRate  metrics.Histogram
	RateLimitDrops metrics.Meter
	Extra          map[string]string
}

func New(clientid, program, version string) *Metrics {
	clientid = strings.Replace(clientid, ":", ".", -1)

	name := func(key string) string {
		return fmt.Sprintf("client_%s.%s", key, clientid)
	}

	sample := func() metrics.Sample {
		return metrics.NewExpDecaySample(MetricsSampleSize, MetricsSampleAlpha)
	}

	extra := map[string]string{
		"ver":   program + "-" + version,
		"ft":    program,
		"dt":    "libkflow",
		"level": "primary",
	}

	return &Metrics{
		TotalFlowsIn:   metrics.GetOrRegisterMeter(name("Total"), nil),
		TotalFlowsOut:  metrics.GetOrRegisterMeter(name("DownsampleFPS"), nil),
		OrigSampleRate: metrics.GetOrRegisterHistogram(name("OrigSampleRate"), nil, sample()),
		NewSampleRate:  metrics.GetOrRegisterHistogram(name("NewSampleRate"), nil, sample()),
		RateLimitDrops: metrics.GetOrRegisterMeter(name("RateLimitDrops"), nil),
		Extra:          extra,
	}
}

func (m *Metrics) Start(url, email, token string, interval time.Duration, proxy *url.URL) {
	proxyURL := ""
	if proxy != nil {
		proxyURL = proxy.String()
	}

	go httptsdb.OpenTSDBWithConfig(httptsdb.OpenTSDBConfig{
		Addr:               url,
		Registry:           metrics.DefaultRegistry,
		FlushInterval:      interval,
		DurationUnit:       time.Millisecond,
		Prefix:             "chf",
		Debug:              false,
		Send:               make(chan []byte, MaxHttpRequests),
		ProxyUrl:           proxyURL,
		MaxHttpOutstanding: MaxHttpRequests,
		Extra:              m.Extra,
		ApiEmail:           &email,
		ApiPassword:        &token,
	})
}

func NewMeter() metrics.Meter {
	return metrics.NewMeter()
}

func NewHistogram(s metrics.Sample) metrics.Histogram {
	return metrics.NewHistogram(s)
}

func NewUniformSample(n int) metrics.Sample {
	return metrics.NewUniformSample(n)
}
