package config

import (
	"os"
)

var (
	ConfigPath = "config.toml"
	SrvAddr    = "localhost:2600"
)

func Load() {
	if len(os.Getenv("AM_CONFIG_PATH")) > 0 {
		ConfigPath = os.Getenv("AM_CONFIG_PATH")
	}

	if len(os.Getenv("AM_SRV_ADDR")) > 0 {
		SrvAddr = os.Getenv("AM_SRV_ADDR")
	}
}
