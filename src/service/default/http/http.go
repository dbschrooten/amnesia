package http

import (
	"amnesia/src/lib"
	"log"
)

type Service struct {
	Service map[string]interface{}
	lib.Implementation
}

func (s *Service) Error(err error) error {
	return &lib.ServiceError{
		ServiceID: s.Service["id"].(string),
		Err:       err,
	}
}

func (s *Service) Run() error {
	return nil
}

func (s *Service) connect() error {
	log.Print("Connect HTTP")
	return nil
}
