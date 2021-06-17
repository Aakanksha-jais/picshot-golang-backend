package errors

import "fmt"

type Unmarshal struct {
	Err error
	Msg string
}

func (e Unmarshal) Error() string {
	if e.Msg == "" {
		return fmt.Sprintf("invalid request body: %s", e.Err.Error())
	}

	return fmt.Sprintf("%s: %s", e.Msg, e.Err.Error())
}
