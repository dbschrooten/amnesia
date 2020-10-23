package channels

import (
	"amnesia/src/lib"
	"testing"
	"time"
)

func TestSetup(t *testing.T) {
	Setup()
}

func TestListener(t *testing.T) {
	go Listener()

	ServiceEvents <- lib.ChannelEvent{ID: "test1", Time: time.Now()}
	CheckEvents <- lib.ChannelEvent{ID: "test2", Time: time.Now()}
	DebugEvents <- lib.ChannelEvent{ID: "test4", Time: time.Now()}
}
