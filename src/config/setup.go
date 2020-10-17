package config

import (
	"amnesia/src/helpers"
	"amnesia/src/service"
	"fmt"
	"log"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/mitchellh/mapstructure"
)

type ConfigDefaults struct {
	Interval time.Duration
	Timeout  time.Duration
}

var (
	Defaults ConfigDefaults
	Services []service.Service
)

func mapDefaults(config map[string]interface{}) error {
	var (
		defaultInterval time.Duration
		defaultTimeout  time.Duration
		err             error
	)

	if _, ok := config["default"]; !ok {
		defaultInterval, err = time.ParseDuration(DefaultInterval)
		defaultTimeout, err = time.ParseDuration(DefaultTimeout)
	} else {
		if _, ok := config["default"].(map[string]interface{})["interval"]; !ok {
			defaultInterval, err = time.ParseDuration(DefaultInterval)
		} else {
			defaultInterval, err = time.ParseDuration(
				config["default"].(map[string]interface{})["interval"].(string),
			)
		}

		if _, ok := config["default"].(map[string]interface{})["timeout"]; !ok {
			defaultTimeout, err = time.ParseDuration(DefaultTimeout)
		} else {
			defaultTimeout, err = time.ParseDuration(
				config["default"].(map[string]interface{})["timeout"].(string),
			)
		}
	}

	if err != nil {
		return err
	}

	Defaults.Interval = defaultInterval
	Defaults.Timeout = defaultTimeout

	return nil
}

func mapServices(config map[string]interface{}) error {
	// Parse Services
	for Type, Service := range config["service"].(map[string]interface{}) {
		if !helpers.StringInSlice(Type, service.ServiceTypes()) {
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

	return nil
}

func Setup() error {
	var config map[string]interface{}

	if _, err := toml.DecodeFile(ConfigPath, &config); err != nil {
		return err
	}

	log.Print("Load config.toml")

	if err := mapDefaults(config); err != nil {
		return err
	}

	if err := mapServices(config); err != nil {
		return err
	}

	return nil
}
