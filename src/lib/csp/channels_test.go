package channels

import (
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

func TestSetup(t *testing.T) {
	Setup()
}

func TestListener(t *testing.T) {
	go Listener()

	AlertEvents <- Event{Id: "test1", Label: "something1", Time: time.Now()}
	WarningEvents <- Event{Id: "test2", Label: "something2", Time: time.Now()}
	DebugEvents <- Event{Id: "test3", Label: "something3", Time: time.Now()}
}
