package infrastructure

import "fmt"

type RepositoryError struct {
	Message string
}

func (re *RepositoryError) Error() string {
	return fmt.Sprintf("RepositoryError: %s", re.Message)
}

type EntityNotFoundError struct {
	Message string
}

func (enfe *EntityNotFoundError) Error() string {
	return fmt.Sprintf("EntityNotFoundError: %s", enfe.Message)
}
