package status

import (
	"amnesia/src/config"
	"log"
	"time"
)

// Create atomic counter per event
// Reset atomic counter when status change
// Create a unique id per service event
// Connect consecutive service events when last one falls in interval range, or perhaps
// make this user customizable
var (
	Current Status
)

func Setup() error {
	Current.Build()
	log.Print("Setup status")

	return nil
}

type Status struct {
	Services map[string]ServiceStatus `yaml:"services" json:"services"`
}

func (s *Status) Build() {
	s.Services = make(map[string]ServiceStatus)

	for _, svc := range config.Services {
		s.Services[svc.ID] = ServiceStatus{
			Type:   svc.Type,
			ID:     svc.ID,
			Label:  svc.Label,
			Status: true,
			ServiceEvents: func() map[string]EventStatus {
				var res = make(map[string]EventStatus)

				for _, svcEvent := range svc.Event {
					res[svcEvent.ID] = EventStatus{
						Failed: 0,
					}
				}

				return res
			}(),
		}
	}
}

func (s *Status) AddFailure(ServiceID string, EventID string) {
	if event, ok := s.Services[ServiceID].ServiceEvents[EventID]; ok {
		event.Failed++
		s.Services[ServiceID].ServiceEvents[EventID] = event
	}
}

func (s *Status) GetFailure(ServiceID string, EventID string) int {
	return s.Services[ServiceID].ServiceEvents[EventID].Failed
}

func (s *Status) Status() map[string]ServiceStatus {
	return s.Services
}

type EventStatus struct {
	Failed int       `json:"failed"`
	Time   time.Time `json:"last_event"`
}

type ServiceStatus struct {
	Type          string                 `json:"type"`
	ID            string                 `json:"id"`
	Label         string                 `json:"label"`
	Status        bool                   `json:"status"`
	ServiceEvents map[string]EventStatus `json:"service_events"`
	LastCheck     time.Time              `json:"last_check"`
}
