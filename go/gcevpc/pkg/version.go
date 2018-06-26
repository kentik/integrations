package gcevpc

type VersionInfo struct {
	Version  string
	Date     string
	Platform string
	Distro   string
}

var VERSION = VersionInfo{
	Version:  "0.2",
	Date:     "2018",
	Platform: "N/A",
	Distro:   "N/A",
}
