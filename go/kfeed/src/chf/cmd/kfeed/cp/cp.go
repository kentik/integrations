package cp

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"

	"chf/kt"
	"version"

	"github.com/davecgh/go-spew/spew"
	capn "github.com/glycerine/go-capnproto"
	"github.com/kentik/eggs/pkg/baseserver"
	"github.com/kentik/eggs/pkg/logger"
	model "github.com/kentik/proto/kflow"
)

// These are enumerations from kentik
const (
	KFLOW_DNS_QUERY      = 179
	KFLOW_DNS_QUERY_TYPE = 180
	KFLOW_DNS_RET_CODE   = 181
	KFLOW_DNS_RESPONSE   = 204
)

type Alpha struct {
	CompanyId kt.Cid
	Flow      model.CHF
}

type Cp struct {
	log       logger.ContextL
	listen    string
	alphaChan chan *Alpha
}

type Beta struct {
	Protocol uint32
	Bytes    uint64
	Packets  uint64
	Dstip    net.IP
	Srcip    net.IP
}

func NewCp(log logger.ContextL, listen string) (*Cp, error) {
	cp := Cp{
		log:       log,
		listen:    listen,
		alphaChan: make(chan *Alpha, kt.CHAN_SLACK_LARGE),
	}
	return &cp, nil
}

// nolint: errcheck
func (cp *Cp) cleanup() {

}

// Main loop for chfcp
// Run the reducer steps, fan out on a per company basis
func (cp *Cp) generateBeta(ctx context.Context) error {

	cp.log.Infof("Generate Beta Online")
	for {
		select {
		case alpha := <-cp.alphaChan:
			cp.log.Debugf("Got %d", alpha.CompanyId)
			beta := Beta{
				Protocol: alpha.Flow.Protocol(),
				Bytes:    alpha.Flow.InBytes(),
				Packets:  alpha.Flow.InPkts(),
			}

			if alpha.Flow.Ipv4SrcAddr() > 0 {
				beta.Srcip = int2ip(alpha.Flow.Ipv4SrcAddr())
			} else {
				beta.Srcip = net.IP(alpha.Flow.Ipv6SrcAddr())
			}

			if alpha.Flow.Ipv4DstAddr() > 0 {
				beta.Dstip = int2ip(alpha.Flow.Ipv4DstAddr())
			} else {
				beta.Dstip = net.IP(alpha.Flow.Ipv6DstAddr())
			}

			// do something.
			cp.log.Infof("Flow found: %v", spew.Sdump(beta))
		case <-ctx.Done():
			cp.log.Infof("Generate Beta Done")
			return nil
		}
	}
}

// Take flow from http requests, deserialize and pass it on to alphaChan
// Gets called from a goroutine-per-request
func (cp *Cp) handleFlow(w http.ResponseWriter, r *http.Request) {
	var err error

	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			cp.log.Errorf("Error handling request: %v", err)
			fmt.Fprint(w, "BAD")
		} else {
			fmt.Fprint(w, "GOOD")
		}
	}()

	// check company id
	vals := r.URL.Query()
	cidBase := vals.Get(kt.HTTP_COMPANY_ID)
	cid, err := kt.AtoiCid(cidBase)
	if err != nil {
		return
	}

	// read all data
	evt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	// decompress
	dser, err := ioutil.ReadAll(capn.NewDecompressor(bytes.NewBuffer(evt)))
	if err != nil {
		return
	}

	// consume message stream
	src, err := capn.ReadFromStream(bytes.NewBuffer(dser), nil)
	if err != nil {
		return
	}

	// unpack flow messages and pass them down
	messages := model.ReadRootPackedCHF(src).Msgs()
	var sent, dropped int64
	for i := 0; i < messages.Len(); i++ {
		msg := messages.At(i)
		if !msg.Big() { // Don't work on low res data
			if !msg.SampleAdj() {
				msg.SetSampleRate(msg.SampleRate() * 100) // Apply re-sample trick here.
			}

			// send without blocking, dropping the message if the channel buffer is full
			alpha := &Alpha{CompanyId: cid, Flow: msg}
			select {
			case cp.alphaChan <- alpha:
				sent++
			default:
				dropped++
			}

		}
	}
	return

}

func (cp *Cp) handleIntrospectPolicy(w http.ResponseWriter, r *http.Request) {

}

// Get messages into the system.
func (cp *Cp) listenHTTP() {
	cp.log.Infof("Setting up HTTP system on %s%s", cp.listen, kt.ALERT_INBOUND_PATH)
	http.HandleFunc(kt.ALERT_INBOUND_PATH, cp.handleFlow)
	http.HandleFunc(kt.INTROSPECT_POLICY_PATH, cp.handleIntrospectPolicy)
	http.HandleFunc(kt.HEALTH_CHECK_PATH, func(w http.ResponseWriter, r *http.Request) {
		// FIXME(stefan): backport new healthcheck logic here?
		fmt.Fprintf(w, "OK\n")
	})
	err := http.ListenAndServe(cp.listen, nil)
	if err != nil {
		cp.log.Errorf("Error up HTTP system on %s%s -- %v", cp.listen, kt.ALERT_INBOUND_PATH, err)
		panic(err)
	}
}

func (cp *Cp) GetStatus() []byte {
	b := new(bytes.Buffer)
	b.WriteString(fmt.Sprintf("\nCHF Kfeed: %s Built on %s %s (%s)\n", version.VERSION_STRING, version.PLATFORM_STRING, version.DISTRO_STRING, version.DATE_STRING))

	return b.Bytes()
}

// RunHealthCheck implements the baseserver.Service interface.
func (cp *Cp) RunHealthCheck(ctx context.Context, result *baseserver.HealthCheckResult) {
}

// HttpInfo implements the baseserver.Service interface.
func (cp *Cp) HttpInfo(w http.ResponseWriter, r *http.Request) {}

func (cp *Cp) Run(ctx context.Context) error {
	defer cp.cleanup()
	cp.log.Infof("Cp Tx System running")
	go cp.listenHTTP()
	return cp.generateBeta(ctx)
}

func (cp *Cp) Close() {
	// this service uses the ctx object passed in Run, do nothing here
}

func int2ip(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}
