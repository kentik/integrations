package version

const (
	VERSION_STRING = "dirty-a5fb2d72b360468baa3f27a73f2f029e786fe206"
    DATE_STRING = "Tue, 8 May 2018 00:14:54 +0000"
    PLATFORM_STRING = "Linux 3.16.0-4-amd64 x86_64 [go version go1.9.1 linux/amd64]"
    DISTRO_STRING = "Debian GNU/Linux 8.7 (jessie)"
)

type VersionInfo struct {
	Version  string
	Date     string
	Platform string
	Distro   string
}

var VERSION = VersionInfo{
	Version:  VERSION_STRING,
	Date:     DATE_STRING,
	Platform: PLATFORM_STRING,
	Distro:   DISTRO_STRING,
}
