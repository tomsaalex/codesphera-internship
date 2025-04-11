package service

import "fmt"

type ServiceError struct {
	Message string
}

func (se *ServiceError) Error() string {
	return fmt.Sprintf("ServiceError: %s", se.Message)
}
