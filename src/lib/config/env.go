package config

import (
	"os"
)

var (
	ConfigPath string
	SrvAddr    string
)

func Load() {
	ConfigPath = os.Getenv("AM_CONFIG_PATH")
	SrvAddr = os.Getenv("AM_SRV_ADDR")
}
