package errors

import "fmt"

type BodyRead struct {
	Err error
}

func (e BodyRead) Error() string {
	return fmt.Sprintf("error in reading request body: %s", e.Err.Error())
}
