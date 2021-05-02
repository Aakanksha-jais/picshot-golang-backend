package errors

import "fmt"

type Error struct {
	Type    string
	Message string
	Err     error
}

func (e Error) Error() string {
	if e.Err == nil {
		return e.Message
	}

	return fmt.Sprintf("%v; %s", e.Err, e.Message)
}
