package apperr

import "errors"

type HttpError struct {
	Code    int    `json:"status_code"`
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

func IsHttpError(err error) bool {
	var httpError *HttpError
	return err != nil && errors.As(err, &httpError)
}
