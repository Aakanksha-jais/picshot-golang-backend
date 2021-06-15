package errors

import (
	"fmt"
)

type Error struct {
	Type string
	Msg  string
	Err  error
}

func (e Error) Error() string {
	if e.Err == nil {
		return e.Msg
	}

	return fmt.Sprintf("%v; %s", e.Err, e.Msg)
}
