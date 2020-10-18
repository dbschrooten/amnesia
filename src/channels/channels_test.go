package channels

import (
	"testing"
	"time"
)

func TestSetup(t *testing.T) {
	Setup()
}

func TestListener(t *testing.T) {
	go Listener()

	ServiceEvents <- Event{Id: "test1", Time: time.Now()}
	CheckEvents <- Event{Id: "test2", Time: time.Now()}
	DebugEvents <- Event{Id: "test4", Time: time.Now()}
}
