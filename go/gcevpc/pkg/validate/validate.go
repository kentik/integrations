package validate

import (
	"context"
	"net/http"

	"encoding/json"

	"regexp"

	"fmt"

	"net"

	"time"

	"strings"

	"sync/atomic"

	"cloud.google.com/go/pubsub"
	"github.com/gorilla/mux"
	"github.com/kentik/eggs/pkg/baseserver"
	"github.com/kentik/eggs/pkg/logger"
)

const (
	validationTimeout = 2 * time.Minute // total timeout for validation handler
	existTimeout      = 30 * time.Second
	receiveTimeout    = 1 * time.Minute
	maxMessageSize    = 16 * 1024
)

type ValidatorService struct {
	listenAddr  string
	log         logger.ContextL
	writeStdOut bool
}

func NewValidatorService(listenAddr string, lc logger.ContextL, writeStdOut bool) (*ValidatorService, error) {
	return &ValidatorService{
		listenAddr:  listenAddr,
		log:         lc,
		writeStdOut: writeStdOut,
	}, nil
}

func (svc *ValidatorService) GetStatus() []byte {
	return []byte("OK")
}

func (svc *ValidatorService) RunHealthCheck(ctx context.Context, result *baseserver.HealthCheckResult) {
	// noop for now
}

func (svc *ValidatorService) Close() {
	// this service uses the ctx object passed in Run, do nothing here
}

func (svc *ValidatorService) HttpInfo(w http.ResponseWriter, req *http.Request) {
	// noop for now
}

type response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func resp(w http.ResponseWriter, success bool, message string, code int) {
	body, err := json.Marshal(response{
		Success: success,
		Message: message,
		Code:    code,
	})
	if err != nil {
		body = []byte("{}")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(body)
	w.Write([]byte{'\n'}) // heh
}

func (svc *ValidatorService) success(w http.ResponseWriter, message string) {
	resp(w, true, message, http.StatusOK)
	svc.log.Infof("returning success response: %s", message)
}

func (svc *ValidatorService) fail(w http.ResponseWriter, message string, code int) {
	resp(w, false, message, code)
	svc.log.Infof("returning fail response: %s", message)
}

func (svc *ValidatorService) errfail(w http.ResponseWriter, err error, code int) {
	resp(w, false, fmt.Sprintf("error: %+v", err), code)
	svc.log.Infof("returning error response: %+v", err)
}

func partial(data []byte) string {
	var ret string
	if len(data) > maxMessageSize {
		ret = string(data[0:maxMessageSize]) + "..."
	} else {
		ret = string(data)
	}
	return strings.Replace(ret, "\n", " ", -1)
}

// projects/kentik-vpc-flow/subscriptions/self-vpc-flows-sub
var validProject = regexp.MustCompile(`^[a-zA-Z0-9-_]+$`)

// Must be 3-255 characters, start with a letter, and contain only the following characters: letters, numbers, dashes (-), periods (.), underscores (_), tildes (~), percents (%) or plus signs (+). Cannot start with goog.
var validSub = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9\-\._\~\%\+]{3,255}$`)

func (svc *ValidatorService) validateHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), validationTimeout)
	defer cancel()

	vars := mux.Vars(r)

	projectName := vars["project"]
	subName := vars["subscription"]

	svc.log.Infof("validate: project=%s, sub=%s", projectName, subName)

	if !validProject.MatchString(projectName) {
		svc.fail(w, "invalid project name", http.StatusBadRequest)
		return
	}
	if !validSub.MatchString(subName) {
		svc.fail(w, "invalid subscription name", http.StatusBadRequest)
		return
	}

	svc.log.Infof("validate: %s/%s creating client", projectName, subName)
	client, err := pubsub.NewClient(ctx, projectName)
	if err != nil {
		svc.errfail(w, err, http.StatusBadRequest) // could also be internal error?
		return
	}

	// create sub (this cannot fail)
	sub := client.Subscription(subName)
	sub.ReceiveSettings.NumGoroutines = 1

	// check for existence
	svc.log.Infof("validate: %s/%s: checking existence", projectName, subName)
	existCtx, existCancel := context.WithTimeout(ctx, existTimeout)
	defer existCancel()
	exists, err := sub.Exists(existCtx)
	if err != nil {
		svc.errfail(w, err, http.StatusInternalServerError)
		return
	}
	if !exists {
		svc.fail(w, "subscription does not exist", http.StatusBadRequest)
		return
	}

	// check that we can receive messages
	svc.log.Infof("validate: %s/%s: checking receive", projectName, subName)
	var received uint32
	receiveCtx, receiveCancel := context.WithTimeout(ctx, receiveTimeout)
	if err := sub.Receive(receiveCtx, func(_ context.Context, message *pubsub.Message) {
		message.Nack() // don't actually consume this message
		if atomic.CompareAndSwapUint32(&received, 0, 1) {
			receiveCancel() // stop receiving
			svc.success(w, "subscription OK - sample message: "+partial(message.Data))
		}
	}); err != nil {
		svc.errfail(w, err, http.StatusInternalServerError)
	}

	if received == 0 {
		svc.success(w, "subscription exists, but we did not receive a message")
	}
}

func (svc *ValidatorService) Run(parentCtx context.Context) (err error) {

	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	r := mux.NewRouter()
	apiv1 := r.PathPrefix("/api/v1").Subrouter()
	apiv1.HandleFunc("/validate/{project}/{subscription}", svc.validateHandler)

	ln, err := net.Listen("tcp", svc.listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()

	go func() {
		svc.log.Infof("ready on http://%+v", ln.Addr())
		httpErr := http.Serve(ln, r)
		if ctx.Err() == nil {
			err = httpErr // if we're not shutting down, bubble this error up
		}
		svc.log.Infof("http server stopped")
		cancel()
	}()

	select {
	case <-ctx.Done():
		svc.log.Infof("shutting down")
		break
	}

	return err
}
