package utils

import "fmt"

type Error struct {
	Status int32
	Err    error
}

func (e *Error) Error() string {
	return fmt.Sprintf("status: %d, error: %s", e.Status, e.Err)
}

// create new error
func NewError(status int32, err error) *Error {
	return &Error{
		Status: status,
		Err:    err,
	}
}
