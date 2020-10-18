package main

import (
	"amnesia/src/extension"
	"amnesia/src/service"
	"log"
)

type ServiceRequired struct {
	MinStatus string
}

type ServiceExt struct {
	Service  service.Service
	required ServiceRequired
	extension.Extension
}

func (s *ServiceExt) Info() map[string]string {
	return map[string]string{
		"type":    "service",
		"name":    "elasticsearch",
		"version": "1.0.0",
		"author":  "David Schrooten",
	}
}

func (s *ServiceExt) Run() error {
	log.Print("Check Elasticsearch")
	return nil
}

func (s *ServiceExt) connect() error {
	log.Print("Connect Elasticsearch")
	return nil
}

var (
	Extension ServiceExt
)
