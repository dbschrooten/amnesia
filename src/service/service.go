package service

import (
	"amnesia/src/helpers"
	"errors"
)

var (
	ServiceTypes = []string{
		"elasticsearch",
		"rabbitmq",
		"mysql",
		"postgresql",

		"http",
		"https",
		"tcp",
		"udp",

		"kubernetes_pod",
		"kubernetes_statefulset",
		"kubernetes_deployment",
	}
)

type Service struct {
	Type     string
	Id       string
	Label    string   `toml:"label"`
	Address  string   `toml:"address"`
	Host     string   `toml:"host"`
	Port     int      `toml:"port"`
	Alerts   []string `toml:"alerts"`
	Required interface{}
	Alert    interface{}
}

func (s *Service) Setup() error {
	if !helpers.StringInSlice(
		s.Type,
		ServiceTypes,
	) {
		return errors.New("Unknown service type")
	}

	return nil
}
