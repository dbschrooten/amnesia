package service

import (
	"amnesia/src/helpers"
	"errors"
)

var (
	ServiceImplementations = []string{
		"http",
		"https",
		"tcp",
		"udp",
		"telnet",
		"graphql",
	}
	KubernetesImplementations = []string{
		"kubernetes_pod",
		"kubernetes_statefulset",
		"kubernetes_deployment",
	}
	PluginImplementations []string
)

func ServiceTypes() []string {
	var result []string

	result = append(result, ServiceImplementations...)
	result = append(result, KubernetesImplementations...)
	result = append(result, PluginImplementations...)

	return result
}

type Service struct {
	Type     string
	Id       string
	Label    string   `toml:"label"`
	Address  string   `toml:"address"`
	Host     string   `toml:"host"`
	Port     int      `toml:"port"`
	Interval string   `toml:"interval"`
	Timeout  string   `toml:"timeout"`
	Alerts   []string `toml:"alerts"`
	Required interface{}
	Alert    interface{}
}

func (s *Service) Run() error {
	if !helpers.StringInSlice(
		s.Type,
		ServiceTypes(),
	) {
		return errors.New("Unknown service type")
	}

	return nil
}
