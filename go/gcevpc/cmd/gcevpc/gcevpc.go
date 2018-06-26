package main

import (
	"flag"
	"fmt"

	"github.com/kentik/integrations/go/gcevpc/pkg"

	"github.com/kentik/eggs/pkg/baseserver"
	"github.com/kentik/eggs/pkg/logger"
	"github.com/kentik/eggs/pkg/properties"
)

var (
	FLAG_sourceSub     = flag.String("sub", "", "Google Sub to listen for flows on")
	FLAG_projectID     = flag.String("project", "", "Google ProjectID to listen for flows on")
	FLAG_dstAddr       = flag.String("dest", "", "Address to send flow to. If not set, defaults to https://flow.kentik.com")
	FLAG_email         = flag.String("api_email", "", "Kentik Email Address")
	FLAG_token         = flag.String("api_token", "", "Kentik Email Token")
	FLAG_plan          = flag.Int("plan_id", 0, "Kentik Plan ID to use for devices")
	FLAG_site          = flag.Int("site_id", 0, "Kentik Site ID to use for devices")
	FLAG_device        = flag.String("device_map_type", "subnet", "Define mapping to Kentik device. Options: subnet, vmname, project")
	FLAG_dropIntraDest = flag.Bool("drop_intra_dest", false, "Drop all intra-VPC Dest logs.")
	FLAG_dropIntraSrc  = flag.Bool("drop_intra_src", false, "Drop all intra-VPC Src logs")
	FLAG_writeStdout   = flag.Bool("json", false, "Write flows to stdout in json form.")

	ValidDeviceMappings = map[string]bool{
		"subnet":  true,
		"vmname":  true,
		"project": true,
	}
)

func main() {
	bs := baseserver.Boilerplate("gcevpc", gcevpc.VERSION, properties.NewEnvPropertyBacking())
	lc := logger.NewContextLFromUnderlying(logger.SContext{S: "GCEVPC"}, bs.Logger)

	if !ValidDeviceMappings[*FLAG_device] {
		bs.Fail(fmt.Sprintf("Invalid device mapping: %s. Options: %v", *FLAG_device, ValidDeviceMappings))
	}

	cpr, err := gcevpc.NewCp(lc, *FLAG_sourceSub, *FLAG_projectID, *FLAG_dstAddr, *FLAG_email, *FLAG_token, *FLAG_plan, *FLAG_site, *FLAG_device,
		*FLAG_dropIntraDest, *FLAG_dropIntraSrc, *FLAG_writeStdout)
	if err != nil {
		bs.Fail(fmt.Sprintf("Cannot start gcevpc: %v", err))
	}

	bs.Run(cpr)
}
