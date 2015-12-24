package config

import (
	"os"
	"time"
)

var (
	Addr    = ":2547"
	Unix    = false
	Timeout = 1 * time.Second
)

func init() {
	addr := os.Getenv("NETCLIP_ADDR")
	if addr != "" {
		Addr = addr
	}

	unix := os.Getenv("NETCLIP_UNIX")
	Unix = unix == "1"

	timeoutStr := os.Getenv("NETCLIP_TIMEOUT")
	timeout, err := time.ParseDuration(timeoutStr)
	if err == nil {
		Timeout = timeout
	}
}
