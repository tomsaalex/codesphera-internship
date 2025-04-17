package service

import (
	"fmt"
	"strconv"
)

type InvalidReason int

const (
	EMPTY   int = 1
	INVALID int = 2
)

type ValidationError struct {
	fieldErrors map[string]int
}

func (v *ValidationError) GetField(name string) (int, bool) {
	fieldError, errorExists := v.fieldErrors[name]
	return fieldError, errorExists
}

func NewValidationError() *ValidationError {
	return &ValidationError{
		fieldErrors: make(map[string]int),
	}
}

func (ve *ValidationError) Error() string {
	invalidFields := ""
	for fieldName, fieldError := range ve.fieldErrors {
		invalidFields += "(" + fieldName + "," + strconv.Itoa(fieldError) + ") "
	}
	return fmt.Sprintf("ValidationError: %s", invalidFields)
}
