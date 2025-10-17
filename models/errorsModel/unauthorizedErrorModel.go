package errorsModel

type UnauthorizedErrorModel struct {
	Err error
}

func (e *UnauthorizedErrorModel) Error() string {
	return e.Err.Error()
}
