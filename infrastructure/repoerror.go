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

type EntityDBMappingError struct {
	Message string
}

func (eme *EntityDBMappingError) Error() string {
	return fmt.Sprintf("EntityDBMappingError: %s", eme.Message)
}

type ForeignKeyViolationError struct {
	Message string
}

func (fke *ForeignKeyViolationError) Error() string {
	return fmt.Sprintf("ForeignKeyViolationError: %s", fke.Message)
}
