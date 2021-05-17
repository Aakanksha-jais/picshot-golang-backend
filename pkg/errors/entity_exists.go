package errors

import "fmt"

type EntityAlreadyExists struct {
	Entity    string `json:"entity"`
	ValueType string `json:"value_type"`
	Value     string `json:"value"`
}

func (e EntityAlreadyExists) Error() string {
	return fmt.Sprintf("%s already exists with %s %s", e.Entity, e.ValueType, e.Value)
}
