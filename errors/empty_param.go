package errors

import "fmt"

type MissingParam struct {
	Param string `json:"param"`
}

func (e MissingParam) Error() string {
	return fmt.Sprintf("missing value for param %s", e.Param)
}
