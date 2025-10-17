package errorsModel

type UnprocessableErrorModel struct {
	Err error
}

func (e *UnprocessableErrorModel) Error() string {
	return e.Err.Error()
}
