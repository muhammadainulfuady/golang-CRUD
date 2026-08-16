package exception

type NotFoundErr struct {
	ErrMessage string
}

func (e *NotFoundErr) Error() string {
	return e.ErrMessage
}

func NewNotFounfError(message string) NotFoundErr {
	return NotFoundErr{
		ErrMessage: message,
	}
}
