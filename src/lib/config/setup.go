package config

import (
	"amnesia/src/lib/helpers"
	"amnesia/src/lib/service"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/mitchellh/mapstructure"
)

var (
	Services []service.Service
)

func Setup() error {
	var config map[string]interface{}

	if _, err := toml.DecodeFile("src/config.toml", &config); err != nil {
		return err
	}

	log.Print("Load config.toml")

	// Parse Services
	for Type, Service := range config["service"].(map[string]interface{}) {
		if !helpers.StringInSlice(Type, service.ServiceTypes) {
			return fmt.Errorf("Config contains non-existing service type %s", Type)
		}

		// Need more validation
		for Id, Config := range Service.(map[string]interface{}) {
			var serviceConfig service.Service
			serviceConfig.Type = Type
			serviceConfig.Id = Id
			mapstructure.Decode(Config, &serviceConfig)
			Services = append(Services, serviceConfig)
		}
	}

	// Parse Alerts
	log.Print(Services)

	return nil
}
