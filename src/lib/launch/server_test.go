package launch

import (
	"os"
	"testing"
)

func TestServer(t *testing.T) {
	os.Setenv("AM_CONFIG_PATH", "../../config.toml")

	if err := Server(); err != nil {
		t.Error(err)
	}

	os.Setenv("AM_CONFIG_PATH", "src/config.toml")

	if err := Server(); err == nil {
		t.Errorf("Should return config not found error")
	}
}
