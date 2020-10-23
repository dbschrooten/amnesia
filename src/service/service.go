package service

import (
	"amnesia/src/channels"
	"amnesia/src/helpers"
	"amnesia/src/lib"
	"amnesia/src/service/default/http"
	"encoding/json"
	"errors"
	"log"
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
	Type     string             `toml:"type" json:"type"`
	ID       string             `toml:"id" json:"id"`
	Label    string             `toml:"label" json:"label"`
	Address  string             `toml:"address" json:"address"`
	Host     string             `toml:"host" json:"host"`
	Port     int                `toml:"port" json:"port"`
	Interval time.Duration      `toml:"interval" json:"interval"`
	Timeout  time.Duration      `toml:"timeout" json:"timeout"`
	Alerts   []string           `toml:"alerts" json:"alerts"`
	Event    []lib.ServiceEvent `toml:"event" json:"event"`
	Alert    []interface{}      `toml:"alert" json:"alert"`
	Export   []interface{}      `toml:"export" json:"export"`
}

func ServiceTypes() []string {
	var result []string

	result = append(result, ServiceImplementations...)
	result = append(result, PluginImplementations...)

	return result
}

func (s *Service) ServiceToInterface() (map[string]interface{}, error) {
	var res map[string]interface{}

	b, err := json.Marshal(s)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Service) Executor(impl lib.Implementation) {
	t := time.NewTimer(s.Interval)
	<-t.C

	channels.CheckEvents <- lib.ChannelEvent{ID: s.ID, Time: time.Now()}

	if err := impl.Run(); err != nil {
		log.Print(err)
		_, ok := err.(*lib.ServiceError)

		if ok {
			channels.ServiceEvents <- lib.ChannelEvent{ID: s.ID, Err: err, Time: time.Now()}
		}
	}

	s.Executor(impl)
}

func (s *Service) ExecDefault() error {
	var impl lib.Implementation

	svc, err := s.ServiceToInterface()

	if err != nil {
		return err
	}

	switch s.Type {
	case "http":
		impl = &http.Service{
			Service: svc,
		}
	case "https":
	case "tcp":
	case "udp":
	case "telnet":
	case "graphql":
	}

	go s.Executor(impl)

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

type Event struct {
	ID     string    `json:"id"`
	Status bool      `json:"status,omitempty"`
	Time   time.Time `json:"time"`
}

type ServiceStatus struct {
	Type          string  `json:"type"`
	ID            string  `json:"id"`
	Label         string  `json:"label"`
	ServiceEvents []Event `json:"service_events"`
	StatusEvents  []Event `json:"status_events"`
}

func (s *Service) Status() (ServiceStatus, error) {
	return ServiceStatus{}, nil
}
