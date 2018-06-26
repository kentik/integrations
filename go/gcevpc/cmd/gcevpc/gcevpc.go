package main

import (
	"flag"
	"fmt"

	"github.com/kentik/eggs/pkg/baseserver"
	"github.com/kentik/eggs/pkg/logger"
	"github.com/kentik/eggs/pkg/properties"
	"github.com/kentik/integrations/go/gcevpc/pkg"
	"github.com/kentik/integrations/go/gcevpc/pkg/cp"
)

func main() {
	var (
		sourceSub     = flag.String("sub", "", "Google Sub to listen for flows on")
		projectID     = flag.String("project", "", "Google ProjectID to listen for flows on")
		dstAddr       = flag.String("dest", "", "Address to send flow to. If not set, defaults to https://flow.kentik.com")
		email         = flag.String("api_email", "", "Kentik Email Address")
		token         = flag.String("api_token", "", "Kentik Email Token")
		plan          = flag.Int("plan_id", 0, "Kentik Plan ID to use for devices")
		site          = flag.Int("site_id", 0, "Kentik Site ID to use for devices")
		device        = flag.String("device_map_type", "subnet", "Define mapping to Kentik device. Options: subnet, vmname, project")
		dropIntraDest = flag.Bool("drop_intra_dest", false, "Drop all intra-VPC Dest logs.")
		dropIntraSrc  = flag.Bool("drop_intra_src", false, "Drop all intra-VPC Src logs")
		writeStdout   = flag.Bool("json", false, "Write flows to stdout in json form.")

		ValidDeviceMappings = map[string]bool{
			"subnet":  true,
			"vmname":  true,
			"project": true,
		}
	)

	bs := baseserver.Boilerplate("gcevpc", gcevpc.Version, properties.NewEnvPropertyBacking())
	lc := logger.NewContextLFromUnderlying(logger.SContext{S: "GCEVPC"}, bs.Logger)

	if !ValidDeviceMappings[*device] {
		bs.Fail(fmt.Sprintf("Invalid device mapping: %s. Options: %v", *device, ValidDeviceMappings))
	}

	cpr, err := cp.NewCp(lc, *sourceSub, *projectID, *dstAddr, *email, *token, *plan, *site, *device,
		*dropIntraDest, *dropIntraSrc, *writeStdout)
	if err != nil {
		bs.Fail(fmt.Sprintf("Cannot start gcevpc: %v", err))
	}

	bs.Run(cpr)
}
