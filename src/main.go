package main

import (
	"amnesia/src/cmd"
	"amnesia/src/helpers"
	"log"
	"os"
)

const (
	ArgServe  = "serve"
	ArgDryRun = "migrate"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatal("Need an argument to proceed")
	}

	if !helpers.StringInSlice(args[0], []string{ArgServe, ArgDryRun}) {
		log.Fatal("Unknown argument")
	}

	switch args[0] {
	case ArgServe:
		log.Fatal(cmd.Server())
	case ArgDryRun:
		log.Print("DRY RUN")
	}
}
