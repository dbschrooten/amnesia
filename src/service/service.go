package service

import (
	"amnesia/src/helpers"
	"amnesia/src/lib"
	"amnesia/src/service/default/http"
	"errors"
	"time"
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
	PluginImplementations []string
)

type Service struct {
	Type     string
	Id       string
	Label    string        `toml:"label"`
	Address  string        `toml:"address"`
	Host     string        `toml:"host"`
	Port     int           `toml:"port"`
	Interval time.Duration `toml:"interval"`
	Timeout  time.Duration `toml:"timeout"`
	Alerts   []string      `toml:"alerts"`
	Event    []lib.ServiceEvent
	Alert    []interface{}
}

func ServiceTypes() []string {
	var result []string

	result = append(result, ServiceImplementations...)
	result = append(result, PluginImplementations...)

	return result
}

func (s *Service) ExecDefault() error {
	var impl lib.Implementation

	switch s.Type {
	case "http":
		impl = &http.Service{
			Service: s,
		}
	case "https":
	case "tcp":
	case "udp":
	case "telnet":
	case "graphql":
	}

	if err := impl.Run(); err != nil {
		return err
	}

	return nil
}

func (s *Service) ExecExt() error {
	return nil
}

func (s *Service) Run() error {
	// check if exists in all
	if !helpers.StringInSlice(
		s.Type,
		ServiceTypes(),
	) {
		return errors.New("Unknown service type")
	}

	if helpers.StringInSlice(s.Type, ServiceImplementations) {
		if err := s.ExecDefault(); err != nil {
			return err
		}
	}

	if helpers.StringInSlice(s.Type, PluginImplementations) {
		if err := s.ExecExt(); err != nil {
			return err
		}
	}

	return nil
}
