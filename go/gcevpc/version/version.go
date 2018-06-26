package version

const (
	VERSION_STRING = "ebf06f4fe270f034e880f5178069b110e1983686"
    DATE_STRING = "Tue, 26 Jun 2018 13:44:06 -0700"
    PLATFORM_STRING = "Darwin 17.6.0 x86_64 [go version go1.10.3 darwin/amd64]"
    DISTRO_STRING = "macOS Darwin 17.6.0 x86_64"
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
