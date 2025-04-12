package service

import "fmt"

type ValidationError struct {
	Message string
}

func (ve *ValidationError) Error() string {
	return fmt.Sprintf("ValidationError: %s", ve.Message)
}
