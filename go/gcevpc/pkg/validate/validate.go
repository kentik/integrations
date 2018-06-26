package validate

import (
	"context"
	"github.com/kentik/eggs/pkg/baseserver"
	"github.com/kentik/eggs/pkg/logger"
	"net/http"
)

type ValidatorService struct {
	log         logger.ContextL
	writeStdOut bool
}

func NewValidatorService(lc logger.ContextL, writeStdOut bool) (*ValidatorService, error) {
	return &ValidatorService{
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

func (svc *ValidatorService) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			break
		}
	}
	return nil
}
