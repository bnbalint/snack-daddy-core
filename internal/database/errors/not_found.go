package database_errors

import "fmt"

type NotFoundError struct {
	Entity string
	ID     int
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Unable to find %s with %d", e.Entity, e.ID)
}
