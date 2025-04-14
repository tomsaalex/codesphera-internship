package sharederrors

import "fmt"

type DuplicateEntityError struct {
	Message string
}

func (dee *DuplicateEntityError) Error() string {
	return fmt.Sprintf("DuplicateEntityError: %s", dee.Message)
}
