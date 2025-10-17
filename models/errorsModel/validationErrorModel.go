package errorsModel

type ValidationErrorModel struct {
	Err error
}

func (e *ValidationErrorModel) Error() string {
	return e.Err.Error()
}
