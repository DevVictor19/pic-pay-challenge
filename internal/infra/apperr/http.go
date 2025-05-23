package apperr

type httpError struct {
	Code    int    `json:"status_code"`
	Message string `json:"message"`
}

func (e *httpError) Error() string {
	return e.Message
}

func NewHttpError(code int, msg string) error {
	return &httpError{
		Code:    code,
		Message: msg,
	}
}
