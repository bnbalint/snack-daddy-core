package database_errors

type ConflictError struct{}

func (e *ConflictError) Error() string {
	return "Create attempt violated a key constraint on the table"
}
