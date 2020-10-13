package implementations

import (
	"amnesia/src/lib/service"
	"log"
)

type ImplHTTPRequired struct {
	InResponse string
	StatusCode int
}

type ImplHTTP struct {
	Service  service.Service
	required ImplHTTPRequired
	service.ServiceImpl
}

func (i *ImplHTTP) Check() error {
	log.Print("Check HTTP")
	return nil
}

func (i *ImplHTTP) Connect() error {
	log.Print("Connect HTTP")
	return nil
}
