package errs

type Impl struct {
	Code    int
	Message string
}

type Error interface {
	GetCode() int
	error
}

func NewError(code int, message string) error {
	return &Impl{
		Code:    code,
		Message: message,
	}
}

func (e *Impl) Error() string {
	return e.Message
}

func (e *Impl) GetCode() int {
	return e.Code
}
