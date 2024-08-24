package dev_mode

type DevModeConfig struct {
	Cmd          string
	Lang         string
	Title        string
	Coverage     int16
	ServeAddress string
	OpenBrowser  bool
}

const (
	golang = "go"
	dotnet = ".net"
)
