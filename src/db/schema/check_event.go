package schema

import (
	"time"
)

type CheckEvent struct {
	ID        string
	ServiceID string
	Time      time.Time
}
