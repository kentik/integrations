package version

const (
	VERSION_STRING = "dirty-a78b097538150480b36023d10f704808007d6dda"
    DATE_STRING = "Thu, 3 May 2018 15:59:32 -0700"
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
