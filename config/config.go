package config

import "os"

var (
	Addr = ":2547"
	Unix = false
)

func init() {
	addr := os.Getenv("NETCLIP_ADDR")
	if addr != "" {
		Addr = addr
	}
	unix := os.Getenv("NETCLIP_UNIX")
	Unix = unix == "1"
}
