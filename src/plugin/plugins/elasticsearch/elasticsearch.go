package main

import (
	"amnesia/src/service"
	"log"
)

type ImplEsRequired struct {
	MinStatus string
}

type ImplEs struct {
	Service  service.Service
	required ImplEsRequired
	service.ServiceImpl
}

func (i *ImplEs) Info() map[string]string {
	return map[string]string{
		"name":    "elasticsearch",
		"version": "1.0.0",
		"author":  "David Schrooten",
	}
}

func (i *ImplEs) Check() error {
	log.Print("Check Elasticsearch")
	return nil
}

func (i *ImplEs) Connect() error {
	log.Print("Connect Elasticsearch")
	return nil
}

var (
	Implementation ImplEs
)
