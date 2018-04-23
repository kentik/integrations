package version

const (
	VERSION_STRING = "dirty-c1433ac3c2510c4f0116e82f131d283ae4c07105"
    DATE_STRING = "Sat, 21 Apr 2018 02:56:54 +0000"
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
