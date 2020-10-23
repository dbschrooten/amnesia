package lib

import "fmt"

type ServiceError struct {
	ServiceID string
	Err       error
}

func (s *ServiceError) Error() string {
	return fmt.Sprintf("%s", s.Err)
}
