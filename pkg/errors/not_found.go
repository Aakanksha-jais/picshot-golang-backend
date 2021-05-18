package errors

import "fmt"

type EntityNotFound struct {
	Entity string `json:"entity"`
	ID     string `json:"id"`
}

func (e EntityNotFound) Error() string {
	if e.ID != "" {
		return fmt.Sprintf("%v %v not found", e.Entity, e.ID)
	}

	return fmt.Sprintf("%v not found", e.Entity)
}
