package errors

import "fmt"

type InvalidParam struct {
	Param string `json:"param"`
}

func (e InvalidParam) Error() string {
	return fmt.Sprintf("invalid value for param: %s", e.Param)
}
