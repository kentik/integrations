package cp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"chf/cmd/gcevpc/cp/types"
	"version"

	"cloud.google.com/go/pubsub"
	"github.com/kentik/eggs/pkg/baseserver"
	"github.com/kentik/eggs/pkg/logger"
	go_metrics "github.com/kentik/go-metrics"
	"github.com/kentik/gohippo"
	"github.com/kentik/libkflow"
)

const (
	CHAN_SLACK     = 1000
	PROGRAM_NAME   = "gcevpc"
	TAG_CHECK_TIME = 60 * time.Second
	TAG_RESET_TIME = 1 * time.Hour
)

type Cp struct {
	log       logger.ContextL
	sub       string
	project   string
	dest      string
	email     string
	token     string
	plan      int
	site      int
	client    *pubsub.Client
	hippo     *hippo.Client
	rateCheck go_metrics.Meter
	rateError go_metrics.Meter
	msgs      chan *types.GCELogLine
}

type hc struct {
	Check float64 `json:"Check"`
	Error float64 `json:"Error"`
	Depth int     `json:"Depth"`
}

func NewCp(log logger.ContextL, sub string, project string, dest string, email string, token string, plan int, site int) (*Cp, error) {
	cp := Cp{
		log:       log,
		sub:       sub,
		project:   project,
		dest:      dest,
		email:     email,
		token:     token,
		plan:      plan,
		site:      site,
		msgs:      make(chan *types.GCELogLine, CHAN_SLACK),
		rateCheck: go_metrics.NewMeter(),
		rateError: go_metrics.NewMeter(),
	}

	hc := hippo.NewHippo("", email, token)
	if hc == nil {
		return nil, fmt.Errorf("Could not create Hippo Client")
	} else {
		cp.hippo = hc
	}

	return &cp, nil
}

// nolint: errcheck
func (cp *Cp) cleanup() {
	cp.client.Close()
}

type flowClient struct {
	sender      *libkflow.Sender
	setSrcTags  bool
	setDestTags bool
}

// Main loop. Take in messages, turn them into kflow, and send them out.
func (cp *Cp) generateKflow(ctx context.Context) error {
	config := libkflow.NewConfig(cp.email, cp.token, PROGRAM_NAME, version.VERSION_STRING)
	clients := map[string]*flowClient{}
	customs := map[string]map[string]uint32{}
	errors := make(chan error, CHAN_SLACK)
	fullUpserts := map[string][]hippo.Upsert{}
	newTag := false

	if cp.dest != "" {
		config.SetFlow(cp.dest)
	}

	tagTick := time.NewTicker(TAG_CHECK_TIME)
	defer tagTick.Stop()

	tagReset := time.NewTicker(TAG_RESET_TIME)
	defer tagReset.Stop()

	for {
		select {
		case msg := <-cp.msgs:
			host := msg.GetHost()
			if _, ok := clients[host]; !ok {
				var client *libkflow.Sender
				var err error

				client, err = libkflow.NewSenderWithDeviceName(host, errors, config)
				if err != nil {
					dconf := msg.GetDeviceConfig(cp.plan, cp.site)
					cp.log.Infof("Creating new device: %s -> %v", dconf.Name, dconf.IPs)
					client, err = libkflow.NewSenderWithNewDevice(dconf, errors, config)
					if err != nil {
						cp.log.Errorf("Cannot start client: %s %v", host, err)
					}
				} else {
					cp.log.Infof("Found existing device: %s", host)
				}

				clients[host] = &flowClient{sender: client}
				customs[host] = map[string]uint32{}

				if client != nil {
					for _, c := range client.Device.Customs {
						customs[host][c.Name] = uint32(c.ID)
					}
				}
			}

			req := msg.ToFlow(customs[host])

			if msg.IsIn() {
				if !clients[host].setSrcTags {
					if clients[host].sender != nil {
						if nu, cnt, err := msg.SetTags(fullUpserts); err != nil {
							cp.log.Errorf("Error setting src tags: %v", err)
						} else {
							cp.log.Infof("%d SRC Tags set for: %s", cnt, host)
							fullUpserts = nu
							newTag = true
						}
					}
					clients[host].setSrcTags = true
				}
				cp.log.Debugf("%s -> %s", msg.Payload.Connection.SrcIP, msg.Payload.Connection.DestIP)
			} else {
				if !clients[host].setDestTags {
					if clients[host].sender != nil {
						if nu, cnt, err := msg.SetTags(fullUpserts); err != nil {
							cp.log.Errorf("Error setting dst tags: %v", err)
						} else {
							cp.log.Infof("%d DST Tags set for: %s", cnt, host)
							fullUpserts = nu
							newTag = true
						}
					}
					clients[host].setDestTags = true
				}
				cp.log.Debugf("%s -> %s", msg.Payload.Connection.DestIP, msg.Payload.Connection.SrcIP)
			}

			if clients[host].sender != nil {
				clients[host].sender.Send(req)
			}
		case _ = <-tagReset.C:
			for h, _ := range clients {
				clients[h].setSrcTags = false
				clients[h].setDestTags = false
			}
		case _ = <-tagTick.C:
			if newTag {
				sent, err := cp.sendHippoTags(fullUpserts)
				if err != nil {
					cp.log.Errorf("Error setting tags: %v", err)
				} else {
					cp.log.Infof("%d tags set", sent)
				}
				newTag = false
			}

		case err := <-errors:
			cp.log.Errorf("Error in kflow: %v", err)
		case <-ctx.Done():
			cp.log.Infof("Generate kflow Done")
			return nil
		}
	}
}

func (cp *Cp) sendHippoTags(upserts map[string][]hippo.Upsert) (int, error) {
	done := 0
	for col, up := range upserts {
		req := &hippo.Req{
			Replace:  false,
			Complete: true,
			Upserts:  up,
		}

		b, err := cp.hippo.EncodeReq(req)
		if err != nil {
			return done, err
		}

		url := fmt.Sprintf("https://api.kentik.com/api/v5/batch/customdimensions/%s/populators", col)
		if req, err := cp.hippo.NewRequest("POST", url, b); err != nil {
			return done, err
		} else {
			if _, err := cp.hippo.Do(context.Background(), req); err != nil {
				return done, err
			}
		}
		done++
	}
	return done, nil
}

// Runs the subscription and reads messages.
func (cp *Cp) runSubscription(sub *pubsub.Subscription) {
	for {
		err := sub.Receive(context.Background(), func(ctx context.Context, m *pubsub.Message) {
			m.Ack()
			var data types.GCELogLine
			if err := json.Unmarshal(m.Data, &data); err != nil {
				cp.rateError.Mark(1)
				cp.log.Errorf("Error reading log line: %v", err)
			} else {
				cp.rateCheck.Mark(1)
				cp.msgs <- &data
			}
		})
		if err != nil {
			cp.log.Errorf("Error on sub system receive -- %v", err)
		}
	}
}

func (cp *Cp) handleIntrospectPolicy(w http.ResponseWriter, r *http.Request) {

}

func (cp *Cp) GetStatus() []byte {
	b := new(bytes.Buffer)
	b.WriteString(fmt.Sprintf("\nCHF GCEVPC: %s Built on %s %s (%s)\n", version.VERSION_STRING, version.PLATFORM_STRING, version.DISTRO_STRING, version.DATE_STRING))

	return b.Bytes()
}

// RunHealthCheck implements the baseserver.Service interface.
func (cp *Cp) RunHealthCheck(ctx context.Context, result *baseserver.HealthCheckResult) {
}

// HttpInfo implements the baseserver.Service interface.
func (cp *Cp) HttpInfo(w http.ResponseWriter, r *http.Request) {
	h := hc{
		Check: cp.rateCheck.Rate5(),
		Error: cp.rateError.Rate5(),
		Depth: len(cp.msgs),
	}

	b, err := json.Marshal(h)
	if err != nil {
		cp.log.Errorf("Error in HC: %v", err)
	} else {
		w.Write(b)
	}
}

func (cp *Cp) Run(ctx context.Context) error {
	defer cp.cleanup()
	cp.log.Infof("GCE VPC System running")

	// Creates a client.
	client, err := pubsub.NewClient(ctx, cp.project)
	if err != nil {
		return err
	} else {
		cp.client = client
	}

	sub := client.Subscription(cp.sub)
	if sub == nil {
		return fmt.Errorf("Subscription not found: %s", cp.sub)
	}

	go cp.runSubscription(sub)
	return cp.generateKflow(ctx)
}

func (cp *Cp) Close() {
	// this service uses the ctx object passed in Run, do nothing here
}
