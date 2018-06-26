package cp

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	flowclient "github.com/kentik/integrations/go/gcevpc/pkg/client"
	"github.com/kentik/integrations/go/gcevpc/pkg/types"

	"cloud.google.com/go/pubsub"

	"github.com/kentik/eggs/pkg/baseserver"
	"github.com/kentik/eggs/pkg/logger"
	go_metrics "github.com/kentik/go-metrics"
	"github.com/kentik/gohippo"
	"github.com/kentik/integrations/go/gcevpc/pkg"
	"github.com/kentik/libkflow"
)

const (
	CHAN_SLACK           = 1000
	PROGRAM_NAME         = "gcevpc"
	TAG_CHECK_TIME       = 60 * time.Second
	INTERFACE_RESET_TIME = 24 * time.Hour
)

type Cp struct {
	log           logger.ContextL
	sub           string
	project       string
	dest          string
	email         string
	token         string
	plan          int
	deviceMap     string
	site          int
	client        *pubsub.Client
	hippo         *hippo.Client
	rateCheck     go_metrics.Meter
	rateError     go_metrics.Meter
	rateInvalid   go_metrics.Meter
	msgs          chan *types.GCELogLine
	dropIntraDest bool
	dropIntraSrc  bool
	writeStdOut   bool
}

type hc struct {
	Check   float64 `json:"Check"`
	Error   float64 `json:"Error"`
	Invalid float64 `json:"Invalid"`
	Depth   int     `json:"Depth"`
}

func NewCp(log logger.ContextL, sub string, project string, dest string, email string, token string, plan int, site int, deviceMap string, dropIntraDest, dropIntraSrc, writeStdOut bool) (*Cp, error) {
	cp := Cp{
		log:           log,
		sub:           sub,
		project:       project,
		dest:          dest,
		email:         email,
		token:         token,
		plan:          plan,
		site:          site,
		deviceMap:     deviceMap,
		msgs:          make(chan *types.GCELogLine, CHAN_SLACK),
		rateCheck:     go_metrics.NewMeter(),
		rateError:     go_metrics.NewMeter(),
		rateInvalid:   go_metrics.NewMeter(),
		dropIntraDest: dropIntraDest,
		dropIntraSrc:  dropIntraSrc,
		writeStdOut:   writeStdOut,
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
	if cp.client != nil {
		cp.client.Close()
	}
}

func (cp *Cp) initClient(msg *types.GCELogLine, host string, errors chan error, clients map[string]*flowclient.FlowClient,
	customs map[string]map[string]uint32) error {

	config := libkflow.NewConfig(cp.email, cp.token, PROGRAM_NAME, gcevpc.Version.Version)
	if cp.dest != "" {
		config.SetFlow(cp.dest)
	}

	var client *libkflow.Sender
	var err error

	client, err = libkflow.NewSenderWithDeviceName(host, errors, config)
	if err != nil {
		dconf := msg.GetDeviceConfig(cp.plan, cp.site, host)
		cp.log.Infof("Creating new device: %s -> %v", dconf.Name, dconf.IPs)
		client, err = libkflow.NewSenderWithNewDevice(dconf, errors, config)
		if err != nil {
			clients[host] = flowclient.NewFlowClient(nil, cp.deviceMap)
			customs[host] = map[string]uint32{}
			return fmt.Errorf("Cannot start client: %s %v", host, err)
		}
	} else {
		cp.log.Infof("Found existing device: %s", host)
	}

	clients[host] = flowclient.NewFlowClient(client, cp.deviceMap)
	customs[host] = map[string]uint32{}

	if client != nil {
		for _, c := range client.Device.Customs {
			customs[host][c.Name] = uint32(c.ID)
		}
	}

	return nil
}

// Main loop. Take in messages, turn them into kflow, and send them out.
func (cp *Cp) generateKflow(ctx context.Context) error {
	clients := map[string]*flowclient.FlowClient{}
	customs := map[string]map[string]uint32{}
	errors := make(chan error, CHAN_SLACK)
	fullUpserts := map[string][]hippo.Upsert{}
	newTag := false

	tagTick := time.NewTicker(TAG_CHECK_TIME)
	defer tagTick.Stop()

	updateInterfaces := time.NewTicker(INTERFACE_RESET_TIME)
	defer updateInterfaces.Stop()

	for {
		select {
		case msg := <-cp.msgs:
			host, err := msg.GetHost(cp.deviceMap)
			if err != nil {
				cp.log.Errorf("Invalid log line: %v", err)
				continue
			}

			vmname, err := msg.GetVMName()
			if err != nil {
				cp.log.Errorf("Invalid log line: %v", err)
				continue
			}

			if _, ok := clients[host]; !ok {
				err := cp.initClient(msg, host, errors, clients, customs)
				if err != nil {
					cp.log.Errorf("InitClient: %v", err)
					if clients[host] == nil {
						continue
					}
				}
			}

			if msg.IsIn() {
				if clients[host].Sender != nil {
					if _, ok := clients[host].SetSrcHostTags[vmname]; !ok {
						clients[host].SetSrcHostTags[vmname] = map[string]bool{}
					}
					if !clients[host].SetSrcHostTags[vmname][msg.Payload.Connection.SrcIP] {
						if nu, cnt, err := msg.SetTags(fullUpserts); err != nil {
							cp.log.Errorf("Error setting src tags: %v", err)
						} else {
							cp.log.Infof("%d SRC Tags set for: %s %s", cnt, host, vmname)
							fullUpserts = nu
							newTag = true
						}

						// And load in an interface for this guy here.
						if intf, err := msg.GetInterface(); err != nil {
							cp.log.Errorf("Error getting interface: %v", err)
						} else {
							clients[host].AddInterface(intf)
						}
						clients[host].SetSrcHostTags[vmname][msg.Payload.Connection.SrcIP] = true
						cp.log.Debugf("%s -> %s", msg.Payload.Connection.SrcIP, msg.Payload.Connection.DestIP)
					}
				}
			} else {
				if clients[host].Sender != nil {
					if _, ok := clients[host].SetDestHostTags[vmname]; !ok {
						clients[host].SetDestHostTags[vmname] = map[string]bool{}
					}
					if !clients[host].SetDestHostTags[vmname][msg.Payload.Connection.DestIP] {
						if nu, cnt, err := msg.SetTags(fullUpserts); err != nil {
							cp.log.Errorf("Error setting dst tags: %v", err)
						} else {
							cp.log.Infof("%d DST Tags set for: %s %s", cnt, host, vmname)
							fullUpserts = nu
							newTag = true
						}

						// And load in an interface for this guy here.
						if intf, err := msg.GetInterface(); err != nil {
							cp.log.Errorf("Error getting interface: %v", err)
						} else {
							clients[host].AddInterface(intf)
						}
						clients[host].SetDestHostTags[vmname][msg.Payload.Connection.DestIP] = true
						cp.log.Debugf("%s -> %s", msg.Payload.Connection.DestIP, msg.Payload.Connection.SrcIP)
					}
				}
			}

			// Turn into Kflow
			req, err := msg.ToFlow(customs[host], clients[host], cp.dropIntraDest, cp.dropIntraSrc)
			if err != nil {
				cp.log.Errorf("Invalid log line: %v", err)
				continue
			}

			// Send to kentik.
			if clients[host].Sender != nil {
				clients[host].Sender.Send(req)
			}

			// If we are logging these, log away.
			if cp.writeStdOut {
				cp.log.Infof("%s", string(msg.ToJson()))
			}
		case _ = <-updateInterfaces.C:
			for h, _ := range clients {
				err := clients[h].UpdateInterfaces(true)
				if err != nil {
					cp.log.Errorf("Error updating interfaces: %v", err)
				}
			}
		case _ = <-tagTick.C:
			if newTag {
				sent, err := cp.sendHippoTags(fullUpserts)
				if err != nil {
					cp.log.Errorf("Error setting tags: %v", err)
				} else {
					cp.log.Infof("%d tags set", sent)
				}

				// And send interfaces if this is the first time though.
				for h, _ := range clients {
					err = clients[h].UpdateInterfaces(false)
					if err != nil {
						cp.log.Errorf("Error updating interfaces: %v", err)
					}
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
		upComb := map[string]hippo.Upsert{}
		for _, ups := range up { // Combine here.
			if _, ok := upComb[ups.Val]; ok {
				for _, rule := range upComb[ups.Val].Rules {
					for _, ruleIn := range ups.Rules {
						if rule.Dir == ruleIn.Dir {
							rule.IPAddresses = append(rule.IPAddresses, ruleIn.IPAddresses...)
						}
					}
				}
			} else {
				upComb[ups.Val] = ups
			}
		}

		upFinal := []hippo.Upsert{}
		for _, ups := range upComb {
			upFinal = append(upFinal, ups)
		}

		for _, ups := range upFinal {
			for _, rule := range ups.Rules {
				// Dedup and verify IPs here.
				ips := map[string]bool{}
				for _, ip := range rule.IPAddresses {
					ips[ip] = true
				}
				ipsArr := make([]string, 0)
				for ip, _ := range ips {
					if ipp := net.ParseIP(ip); ipp != nil {
						ipsArr = append(ipsArr, ip)
					}
				}
				rule.IPAddresses = ipsArr

				cp.log.Debugf("%s %s -> %v", col, ups.Val, rule.IPAddresses)
			}
		}

		req := &hippo.Req{
			Replace:  true,
			Complete: true,
			Upserts:  upFinal,
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
				cp.log.Errorf("Uploading tags %v", err)
			} else {
				done++
			}
		}
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
				if data.IsValid() {
					cp.msgs <- &data
				} else {
					cp.rateInvalid.Mark(1)
				}
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
	return []byte(fmt.Sprintf("\nCHF GCEVPC: %s Built on %s\n", gcevpc.Version.Version, gcevpc.Version.Date))
}

// RunHealthCheck implements the baseserver.Service interface.
func (cp *Cp) RunHealthCheck(ctx context.Context, result *baseserver.HealthCheckResult) {
}

// HttpInfo implements the baseserver.Service interface.
func (cp *Cp) HttpInfo(w http.ResponseWriter, r *http.Request) {
	h := hc{
		Check:   cp.rateCheck.Rate5(),
		Error:   cp.rateError.Rate5(),
		Invalid: cp.rateInvalid.Rate5(),
		Depth:   len(cp.msgs),
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
