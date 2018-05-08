package main

import (
	"flag"
	"fmt"

	"chf/cmd/gcevpc/cp"
	"chf/kt"
	"version"

	"github.com/kentik/eggs/pkg/baseserver"
	"github.com/kentik/eggs/pkg/logger"
	ev "github.com/kentik/eggs/pkg/version"
)

var (
	FLAG_sourceSub     = flag.String("sub", "", "Google Sub to listen for flows on")
	FLAG_projectID     = flag.String("project", "", "Google ProjectID to listen for flows on")
	FLAG_dstAddr       = flag.String("dest", "", "Address to send flow to. If not set, defaults to https://flow.kentik.com")
	FLAG_email         = flag.String("api_email", "", "Kentik Email Address")
	FLAG_token         = flag.String("api_token", "", "Kentik Email Token")
	FLAG_plan          = flag.Int("plan_id", 0, "Kentik Plan ID to use for devices")
	FLAG_site          = flag.Int("site_id", 0, "Kentik Site ID to use for devices")
	FLAG_isDevice      = flag.Bool("is_device_primary", false, "Create one device in kentik per vm, vs on device per VPC.")
	FLAG_dropIntraDest = flag.Bool("drop_intra_dest", false, "Drop all intra-VPC Dest logs.")
	FLAG_dropIntraSrc  = flag.Bool("drop_intra_src", false, "Drop all intra-VPC Src logs")
	FLAG_writeStdout   = flag.Bool("json", false, "Write flows to stdout in json form.")
)

func main() {
	eVeg := ev.VersionInfo{
		Version:  version.VERSION.Version,
		Date:     version.VERSION.Date,
		Platform: version.VERSION.Platform,
		Distro:   version.VERSION.Distro,
	}

	bs := baseserver.Boilerplate("gcevpc", eVeg, kt.DefaultGCEVPCProperties)
	lc := logger.NewContextLFromUnderlying(logger.SContext{S: "GCEVPC"}, bs.Logger)

	cpr, err := cp.NewCp(lc, *FLAG_sourceSub, *FLAG_projectID, *FLAG_dstAddr, *FLAG_email, *FLAG_token, *FLAG_plan, *FLAG_site, *FLAG_isDevice,
		*FLAG_dropIntraDest, *FLAG_dropIntraSrc, *FLAG_writeStdout)
	if err != nil {
		bs.Fail(fmt.Sprintf("Cannot start gcevpc: %v", err))
	}

	bs.Run(cpr)
}
