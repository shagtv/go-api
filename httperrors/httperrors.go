package httperrors

type IHttperror interface {
	Error() string
	Code() int
}

type HttpError struct {
	msg  string
	code int
}

func (e *HttpError) Error() string {
	return e.msg
}

func (e *HttpError) Code() int {
	return e.code
}

func New(msg string, code int) IHttperror {
	return &HttpError{msg, code}
}
