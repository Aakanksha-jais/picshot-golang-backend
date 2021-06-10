package errors

type DBError struct {
	Err error
}

func (e DBError) Error() string {
	if e.Err == nil {
		return "database error"
	}

	return e.Err.Error()
}
