package validate

import (
	"context"
	"net/http"

	"fmt"
	"html"
	"net"

	"github.com/kentik/eggs/pkg/baseserver"
	"github.com/kentik/eggs/pkg/logger"
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

func (svc *ValidatorService) validateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func (svc *ValidatorService) Run(ctx context.Context) error {

	http.HandleFunc("/api/v1/validate/", svc.validateHandler)

	ln, err := net.Listen("tcp", svc.listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()

	go func() {
		http.Serve(ln, nil)
	}()

	for {
		select {
		case <-ctx.Done():
			ln.Close()
			break
		}
	}
	return nil
}
