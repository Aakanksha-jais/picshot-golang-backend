package errors

type DBError struct {
	Err error
}

func (e DBError) Error() string {
	return e.Err.Error()
}
