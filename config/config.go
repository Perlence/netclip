package config

import "os"

var Addr = ":2547"

func init() {
	addr := os.Getenv("NETCLIP_ADDR")
	if addr != "" {
		Addr = addr
	}
}
