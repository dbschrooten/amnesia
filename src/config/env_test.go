package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	os.Setenv("AM_CONFIG_PATH", "src/config.toml")
	os.Setenv("AM_SRV_ADDR", "localhost:2600")

	Load()

	if ConfigPath != "src/config.toml" {
		t.Errorf("Configpath variable not set")
	}

	if SrvAddr != "localhost:2600" {
		t.Errorf("Srv address not set")
	}
}
