package libkflow

import (
	"bytes"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/kentik/libkflow/agg"
	"github.com/kentik/libkflow/api"
	"github.com/kentik/libkflow/flow"
	"github.com/kentik/libkflow/log"
	"github.com/kentik/libkflow/metrics"
	"zombiezen.com/go/capnproto2"
)

// A Sender aggregates and transmits flow information to Kentik.
type Sender struct {
	agg     *agg.Agg
	exit    chan struct{}
	url     *url.URL
	timeout time.Duration
	client  *api.Client
	sample  int
	ticker  *time.Ticker
	workers sync.WaitGroup
	Device  *api.Device
	Errors  chan<- error
	Metrics *metrics.Metrics
}

func newSender(url *url.URL, timeout time.Duration) *Sender {
	return &Sender{
		exit:    make(chan struct{}),
		url:     url,
		timeout: timeout,
		ticker:  time.NewTicker(20 * time.Minute),
	}
}

// Send adds a flow record to the outgoing queue.
func (s *Sender) Send(flow *flow.Flow) {
	log.Debugf("sending flow to aggregator")
	flow.DeviceId = uint32(s.Device.ID)
	s.agg.Add(flow)
}

// Stop requests a graceful shutdown of the Sender.
func (s *Sender) Stop(wait time.Duration) bool {
	s.agg.Stop()
	select {
	case <-s.exit:
		return true
	case <-time.After(wait):
		return false
	}
}

func (s *Sender) GetClient() *api.Client {
	return s.client
}

func (s *Sender) start(agg *agg.Agg, client *api.Client, device *api.Device, n int) error {
	q := s.url.Query()
	q.Set("sid", "0")
	q.Set("sender_id", device.ClientID())

	s.agg = agg
	s.url.RawQuery = q.Encode()
	s.Device = device
	s.client = client
	s.workers.Add(n)

	for i := 0; i < n; i++ {
		go s.dispatch()
	}
	go s.monitor()
	go s.update()

	log.Debugf("sender started with %d workers", n)

	return nil
}

func (s *Sender) dispatch() {
	buf := &bytes.Buffer{}
	cid := [80]byte{}
	url := s.url.String()

	for msg := range s.agg.Output() {
		log.Debugf("dispatching aggregated flow")

		buf.Reset()
		buf.Write(cid[:])

		err := capnp.NewPackedEncoder(buf).Encode(msg)
		if err != nil {
			s.error(err)
			continue
		}

		err = s.client.SendFlow(url, buf)
		if err != nil {
			s.error(err)
			continue
		}
	}
	s.workers.Done()
}

func (s *Sender) monitor() {
	for {
		select {
		case err := <-s.agg.Errors():
			s.error(err)
		case <-s.agg.Done():
			s.workers.Wait()
			s.ticker.Stop()
			s.exit <- struct{}{}
			log.Debugf("sender stopped")
			return
		}
	}
}

func (s *Sender) update() {
	for range s.ticker.C {
		updated, err := s.client.GetDeviceByID(s.Device.ID)
		if err != nil {
			if api.IsErrorWithStatusCode(err, 404) {
				updated = &api.Device{}
			} else {
				log.Debugf("device API request failed: %s", err)
				continue
			}
		}

		if s.Device.MaxFlowRate != updated.MaxFlowRate {
			log.Debugf("updating max FPS to %d", updated.MaxFlowRate)
			s.Device.MaxFlowRate = updated.MaxFlowRate
			s.agg.Configure(updated.MaxFlowRate)
		}

		// if the configured sample rate is 0 then the sender
		// may be using the device sample rate which has just
		// changed, so abort the program
		if s.sample == 0 && s.Device.SampleRate != updated.SampleRate {
			log.Debugf("device sample rate changed, aborting")
			os.Exit(1)
		}
	}
}

func (s *Sender) error(err error) {
	select {
	case s.Errors <- err:
	default:
	}
}
