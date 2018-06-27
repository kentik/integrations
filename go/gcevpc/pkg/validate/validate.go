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
	Key     string `json:"key"`
	Code    int    `json:"code"`
}

func resp(w http.ResponseWriter, success bool, key string, message string, code int) {
	body, err := json.Marshal(response{
		Success: success,
		Message: message,
		Code:    code,
		Key:     key,
	})
	if err != nil {
		body = []byte("{}")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(body)
	w.Write([]byte{'\n'}) // heh
}

func (svc *ValidatorService) success(w http.ResponseWriter, key string, message string) {
	resp(w, true, key, message, http.StatusOK)
	svc.log.Infof("%s: returning success response: %s", key, message)
}

func (svc *ValidatorService) fail(w http.ResponseWriter, key string, message string, code int) {
	resp(w, false, key, message, code)
	svc.log.Infof("%s: returning fail response: %s", key, message)
}

func (svc *ValidatorService) errfail(w http.ResponseWriter, key string, err error, code int) {
	resp(w, false, key, fmt.Sprintf("error: %+v", err), code)
	svc.log.Infof("%s: returning error response: %+v", key, err)
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
var validSub = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9\-._~%+]{3,255}$`)

func (svc *ValidatorService) validateHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), validationTimeout)
	defer cancel()

	vars := mux.Vars(r)

	projectName := vars["project"]
	subName := vars["subscription"]

	key := projectName + "/" + subName

	svc.log.Infof("validate: project=%s, sub=%s", projectName, subName)

	if !validProject.MatchString(projectName) {
		svc.fail(w, key, "invalid project name", http.StatusBadRequest)
		return
	}
	if !validSub.MatchString(subName) {
		svc.fail(w, key, "invalid subscription name", http.StatusBadRequest)
		return
	}

	svc.log.Infof("validate: %s/%s creating client", projectName, subName)
	client, err := pubsub.NewClient(ctx, projectName)
	if err != nil {
		svc.errfail(w, key, err, http.StatusBadRequest) // could also be internal error?
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
		svc.errfail(w, key, err, http.StatusInternalServerError)
		return
	}
	if !exists {
		svc.fail(w, key, "subscription does not exist", http.StatusBadRequest)
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
			svc.success(w, key, "subscription OK - sample message: "+partial(message.Data))
		}
	}); err != nil {
		svc.errfail(w, key, err, http.StatusInternalServerError)
	}

	if received == 0 {
		svc.success(w, key, "subscription exists, but we did not receive a message")
	}
}

func (svc *ValidatorService) uiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<html>
<head><title>Kentik GCE VPC - Subscription checker</title>
<style>
legend {
    background-color: #000;
    color: #fff;
    padding: 3px 6px;
}

input,
label {
    width: 43%;
}

input {
    margin: .5rem 0;
    padding: .5rem;
    border-radius: 4px;
    border: 1px solid #ddd;
}

label {
    display: inline-block;
}

input:invalid + span:after {
    content: '✖';
    color: red;
    padding-left: 5px;
}

input:valid + span:after {
    content: '✓';
    color: green;
    padding-left: 5px;
}

#submit {
	margin-left: 20px;
    border: 2px solid #daa;
}
</style>
</head>

<body>

<script>

function reset() {
	document.getElementById("submit").disabled = false;
}

function log(key, msg) {
	var d = new Date().toLocaleTimeString();
	document.getElementById("result").insertAdjacentHTML('beforeend', '<p>' + d + ' <b>' + key + "</b>:&nbsp;" + msg + '</p>');
}

function validate() {
	document.getElementById("submit").disabled = true;

	project = document.getElementById("project").value;
	sub = document.getElementById("subscription").value;

	response = fetch("/api/v1/validate/" + project + "/" + sub , {
	    method: 'GET',
		cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
	})
	.then(res => res.json())
	.then(function(res) {
		console.log(res);
		log(res.key, res.message);
		reset();
	}).catch(function(err) {
		console.log(err);
		log("Error", err);
		reset();
	})

}
</script>


<div>
<fieldset>
	<legend>Kentik GCE VPC - Subscription checker</legend>
	<form>
	<div>
		<label for="project">Project name:&nbsp;</label>
		<input id="project" type="text" value="kentik-vpc-flow" pattern="[a-zA-Z0-9-_]+" minlength="3" maxlength="255" /><span></span><br/>

		<label for=""subscription">Subscription name:&nbsp;</label>
		<input id="subscription" type="text" value="self-vpc-flows-sub" pattern="[a-zA-Z][a-zA-Z0-9\-\+\._~]+" minlength="3" maxlength="255"/><span></span><br/>

		<input id="submit" type="button" value="Run server check" onClick="validate()" />
	</div>
	</form>
</fieldset>
</div>

<p>
	<div id="result"/>
</p>

</body>
</html>
`)
}

func (svc *ValidatorService) Run(parentCtx context.Context) (err error) {

	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		http.Redirect(w, req, "/ui/index.html", http.StatusTemporaryRedirect)
	})

	ui := r.PathPrefix("/ui").Subrouter()
	ui.HandleFunc("/index.html", svc.uiHandler)

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
