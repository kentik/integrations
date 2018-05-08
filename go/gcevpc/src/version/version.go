package version

const (
	VERSION_STRING = "dirty-41e64fc457ab0230fdaa3eacfbf432988f05e048"
    DATE_STRING = "Mon, 7 May 2018 22:07:38 +0000"
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
