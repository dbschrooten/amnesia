package config

import (
	"amnesia/src/service"
	"fmt"

	"github.com/spf13/viper"
)

var (
	Services = make(map[string]service.Service)
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./")

	viper.SetDefault("default.interval", "5m")
	viper.SetDefault("default.timeout", "30s")
	viper.SetDefault("default.language", "en")
	viper.SetDefault("default.retention", "7d")

	viper.SetDefault("server.address", "localhost:2600")
	viper.SetDefault("server.write_timeout", "15s")
	viper.SetDefault("server.read_timeout", "15s")

	viper.SetDefault("system.plugin_folder", "./plugins")
}

func mapServices() error {
	var svc []service.Service

	if err := viper.UnmarshalKey("service", &svc); err != nil {
		return err
	}

	for _, s := range svc {
		Services[s.ID] = s
	}

	return nil
}

func Setup() error {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return fmt.Errorf("Config file not found")
		}

		return err
	}

	if err := mapServices(); err != nil {
		return err
	}

	return nil
}
