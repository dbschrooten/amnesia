package main

import (
	"amnesia/src/lib/launch"
	"log"
)

func main() {
	if err := launch.Server(); err != nil {
		log.Fatal(err)
	}

	log.Print("Done")
}
