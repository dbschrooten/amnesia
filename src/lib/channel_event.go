package lib

import "time"

type ChannelEvent struct {
	ID   string
	Err  error
	Time time.Time
}
