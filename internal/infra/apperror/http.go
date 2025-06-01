package apperror

type HttpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *HttpError) Error() string {
	return e.Message
}

func NewHttpError(code int, msg string) error {
	return &HttpError{
		Code:    code,
		Message: msg,
	}
}
