package status

import (
	"amnesia/src/config"
	"amnesia/src/db"
	"amnesia/src/db/schema"
	"log"
	"time"

	"github.com/genjidb/genji/document"
)

var (
	Current Status
)

func Setup() error {
	Current.Check()
	log.Print("Setup status")

	return nil
}

type Status struct {
	Services map[string]ServiceStatus `yaml:"services" json:"services"`
}

func (s *Status) Check() {
	s.Services = make(map[string]ServiceStatus)

	for _, svc := range config.Services {
		s.Services[svc.ID] = ServiceStatus{
			Type:   svc.Type,
			ID:     svc.ID,
			Label:  svc.Label,
			Status: true,
			ServiceEvents: func() map[string]EventStatus {
				// check last event
				var res = make(map[string]EventStatus)

				for _, svcEvent := range svc.Event {
					res[svcEvent.ID] = EventStatus{}
				}

				return res
			}(),
			LastCheck: func() time.Time {
				res, err := db.Conn.Query("SELECT * FROM check_events ORDER BY time DESC LIMIT 1")

				defer res.Close()

				if err != nil {
					log.Fatal(err)
				}

				var result schema.CheckEvent

				if err := res.Iterate(func(d document.Document) error {
					if err := document.StructScan(d, &result); err != nil {
						return err
					}

					return nil
				}); err != nil {
					log.Fatal(err)
				}

				return result.Time
			}(),
		}
	}
}

// func (s *Status) AddFailure(ServiceID string, EventID string) {
// 	if event, ok := s.Services[ServiceID].ServiceEvents[EventID]; ok {
// 		event.Failed++
// 		s.Services[ServiceID].ServiceEvents[EventID] = event
// 	}
// }

// func (s *Status) GetFailure(ServiceID string, EventID string) int {
// 	return s.Services[ServiceID].ServiceEvents[EventID].Failed
// }

func (s *Status) Status() map[string]ServiceStatus {
	return s.Services
}

type EventStatus struct {
	Time time.Time `json:"last_event_today"`
}

type AlertStatus struct {
	Event string    `json:"event"`
	Time  time.Time `json:"last_event_today"`
}

type ServiceStatus struct {
	Type          string                 `json:"type"`
	ID            string                 `json:"id"`
	Label         string                 `json:"label"`
	Status        bool                   `json:"status"`
	Alerts        map[string]AlertStatus `json:"triggered_alerts"`
	ServiceEvents map[string]EventStatus `json:"service_events"`
	LastCheck     time.Time              `json:"last_check"`
}
