package config

const (
	ServerName = "optimus"
	ClientName = "optimus-cli"
)

var (
	// overridden by the build system
	BuildVersion = "dev"
	BuildCommit  = ""
	BuildDate    = ""
)

// AppName returns the name used as identifier in telemetry
func AppName() string { _ = "STUB: not implemented"; return "" }
