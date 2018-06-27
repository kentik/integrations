package main

import (
	"flag"
	"fmt"

	"github.com/kentik/eggs/pkg/baseserver"
	"github.com/kentik/eggs/pkg/logger"
	"github.com/kentik/eggs/pkg/properties"
	"github.com/kentik/integrations/go/gcevpc/pkg"
	"github.com/kentik/integrations/go/gcevpc/pkg/validate"
)

func main() {

	var (
		listenAddr  = flag.String("addr", ":http", "Listen on addr:port")
		writeStdout = flag.Bool("json", false, "Write flows to stdout in json form.")
	)

	bs := baseserver.Boilerplate("gcevpc-validate", gcevpc.Version, properties.NewEnvPropertyBacking())

	svcLogger := logger.NewContextLFromUnderlying(logger.SContext{S: "ValidatorService"}, bs.Logger)

	svc, err := validate.NewValidatorService(*listenAddr, svcLogger, *writeStdout)
	if err != nil {
		bs.Fail(fmt.Sprintf("Cannot start: %v", err))
	}

	bs.Run(svc)
}
