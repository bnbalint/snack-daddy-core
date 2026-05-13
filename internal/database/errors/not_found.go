package database_errors

import "fmt"

type NotFoundError struct {
	Entity string
	ID     string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Unable to find %s with %s", e.Entity, e.ID)
}
