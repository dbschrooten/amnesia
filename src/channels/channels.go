package channels

import (
	"log"
	"time"
)

type Event struct {
	Id   string
	Time time.Time
}

var (
	ServiceEvents = make(chan Event)
	CheckEvents   = make(chan Event)
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
		case e := <-ServiceEvents:
			log.Printf("Service event received %v", e)
		case e := <-CheckEvents:
			log.Printf("Check event received %v", e)
		case e := <-DebugEvents:
			log.Printf("Debug event received %v", e)
		}
	}
}
