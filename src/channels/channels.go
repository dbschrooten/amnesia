package channels

import (
	"amnesia/src/channels/processors"
	"amnesia/src/lib"
	"log"
)

var (
	ServiceEvents = make(chan lib.ChannelEvent)
	CheckEvents   = make(chan lib.ChannelEvent)
	DebugEvents   = make(chan lib.ChannelEvent)
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

			if err := processors.ServiceEvent(e); err != nil {
				log.Print(e)
			}
		case e := <-CheckEvents:
			log.Printf("Status event received %v", e)

			if err := processors.CheckEvent(e); err != nil {
				log.Print(e)
			}
		case e := <-DebugEvents:
			log.Printf("Debug event received %v", e)
		}
	}
}
