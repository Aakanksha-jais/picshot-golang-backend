package errors

import "fmt"

type BodyRead struct {
	Err error
	Msg string
}

func (e BodyRead) Error() string {
	if e.Msg == "" {
		return fmt.Sprintf("error in reading request body: %s", e.Err.Error())
	}

	return fmt.Sprintf("%s: %s", e.Msg, e.Err.Error())
}
