package http

import (
	"amnesia/src/service"
	"log"
)

type ServiceRequired struct {
	InResponse string
	StatusCode int
}

type Service struct {
	Service  service.Service
	required ServiceRequired
	service.Implementation
}

func (s *Service) Run() error {
	log.Print("Check HTTP")
	return nil
}

func (s *Service) connect() error {
	log.Print("Connect HTTP")
	return nil
}
