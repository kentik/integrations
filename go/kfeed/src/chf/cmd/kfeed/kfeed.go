package main

import (
	"flag"
	"fmt"

	"chf/cmd/kfeed/cp"
	"chf/kt"
	"version"

	"github.com/kentik/eggs/pkg/baseserver"
	"github.com/kentik/eggs/pkg/logger"
	ev "github.com/kentik/eggs/pkg/version"
)

var (
	FLAG_listen = flag.String("listen", "127.0.0.1:3456", "HTTP port to bind on")
)

func main() {
	eVeg := ev.VersionInfo{
		Version:  version.VERSION.Version,
		Date:     version.VERSION.Date,
		Platform: version.VERSION.Platform,
		Distro:   version.VERSION.Distro,
	}

	bs := baseserver.Boilerplate("kfeed", eVeg, kt.DefaultKFeedProperties)
	lc := logger.NewContextLFromUnderlying(logger.SContext{S: "RunBGP "}, bs.Logger)

	cpr, err := cp.NewCp(lc, *FLAG_listen)
	if err != nil {
		bs.Fail(fmt.Sprintf("Cannot start kfeed: %v", err))
	}

	bs.Run(cpr)
}
