package errors

import (
	"fmt"
)

type AuthError struct {
	Msg string
	Err error
}

func (e AuthError) Error() string {
	if e.Err == nil {
		return e.Msg
	}

	return fmt.Sprintf("%v: %s", e.Err, e.Msg)
}
