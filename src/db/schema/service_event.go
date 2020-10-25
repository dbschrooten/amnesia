package schema

import (
	"time"
)

type ServiceEvent struct {
	ID        string
	ServiceID string
	Time      time.Time
}
