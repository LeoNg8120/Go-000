package sentinel

import (
	"errors"
	"fmt"
)

type StatusError struct {
	Code int
	Message string
}

func (e *StatusError) Error() string {
	return fmt.Sprintf("error: code = %d desc = %s", e.Code, e.Message)
}


func NotFound() error {
	return &StatusError{
		Code:    5,
		Message: "NotFound",
	}
}

func IsNotFound(err error) bool {
	if se := new(StatusError); errors.As(err, &se) {
		return se.Code == 5
	}
	return false
}