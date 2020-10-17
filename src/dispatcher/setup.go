package dispatcher

import (
	"amnesia/src/config"
	"log"
)

func Setup() {
	log.Print("Dispatcher setup")

	for _, service := range config.Services {
		log.Print(service.Type)
	}
}
