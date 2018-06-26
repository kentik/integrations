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
		writeStdout = flag.Bool("json", false, "Write flows to stdout in json form.")
	)

	bs := baseserver.Boilerplate("gcevpc-validate", gcevpc.Version, properties.NewEnvPropertyBacking())
	lc := logger.NewContextLFromUnderlying(logger.SContext{S: "GCEVPC-VALIDATE"}, bs.Logger)

	svc, err := validate.NewValidatorService(lc, *writeStdout)
	if err != nil {
		bs.Fail(fmt.Sprintf("Cannot start: %v", err))
	}

	bs.Run(svc)
}
