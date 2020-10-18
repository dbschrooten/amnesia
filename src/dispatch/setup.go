package dispatch

import (
	"amnesia/src/config"
	"log"
)

func Setup() {
	log.Print("Dispatch setup")

	for _, service := range config.Services {
		service.Run()
	}
}
