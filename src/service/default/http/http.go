package http

import (
	"amnesia/src/lib"
	"log"
)

type Service struct {
	Service interface{}
	lib.Implementation
}

func (s *Service) Run() error {
	log.Print("Check HTTP")
	return nil
}

func (s *Service) connect() error {
	log.Print("Connect HTTP")
	return nil
}
