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
	for _, Service := range config["service"].([]map[string]interface{}) {
		if !helpers.StringInSlice(Service["type"].(string), service.ServiceTypes()) {
			return fmt.Errorf("Config contains non-existing service type %s", Service["type"].(string))
		}

		var serviceConfig service.Service
		mapstructure.Decode(Service, &serviceConfig)

		// do more validation in here
		if Service["interval"] != nil {
			intVal, err := time.ParseDuration(Service["interval"].(string))

			if err != nil {
				return err
			}

			serviceConfig.Interval = intVal
		} else {
			serviceConfig.Interval = Defaults.Interval
		}

		if Service["timeout"] != nil {
			timeOut, err := time.ParseDuration(Service["timeout"].(string))

			if err != nil {
				return err
			}

			serviceConfig.Timeout = timeOut
		} else {
			serviceConfig.Timeout = Defaults.Timeout
		}

		Services = append(Services, serviceConfig)
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
