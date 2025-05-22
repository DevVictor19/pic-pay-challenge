package apperr

import "errors"

type HttpErrCode error

var (
	ErrNotFound   HttpErrCode = errors.New("NOT_FOUND")
	ErrBadRequest HttpErrCode = errors.New("BAD_REQUEST")
	ErrConflict   HttpErrCode = errors.New("CONFLICT")
	ErrInternal   HttpErrCode = errors.New("INTERNAL")
)
