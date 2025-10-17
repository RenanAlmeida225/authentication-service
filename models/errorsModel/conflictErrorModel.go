package errorsModel

type ConflictErrorModel struct {
	Err error
}

func (e *ConflictErrorModel) Error() string {
	return e.Err.Error()
}
