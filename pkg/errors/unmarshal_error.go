package errors

import "fmt"

type Unmarshal struct {
	Err error
}

func (e Unmarshal) Error() string {
	return fmt.Sprintf("invalid request body: %s", e.Err.Error())
}
