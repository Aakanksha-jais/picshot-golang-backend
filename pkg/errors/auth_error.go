package errors

import (
	"fmt"
)

type AuthError struct {
	Message string
	Err     error
}

func (e AuthError) Error() string {
	if e.Err == nil {
		return e.Message
	}

	return fmt.Sprintf("%v: %s", e.Err, e.Message)
}
