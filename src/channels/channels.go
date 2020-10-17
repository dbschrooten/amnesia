package channels

import (
	"log"
	"time"
)

type Event struct {
	Id    string
	Label string
	Time  time.Time
}

var (
	AlertEvents   = make(chan Event)
	WarningEvents = make(chan Event)
	DebugEvents   = make(chan Event)
)

func Setup() {
	log.Print("Setup channels")
	go Listener()
}

func Listener() {
	log.Print("Initiating Listener")

	for {
		select {
		case e := <-AlertEvents:
			log.Printf("Alert event received %v", e)
		case e := <-WarningEvents:
			log.Printf("Warning event received %v", e)
		case e := <-DebugEvents:
			log.Printf("Debug event received %v", e)
		}
	}
}
