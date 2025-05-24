package utils

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/DevVictor19/pic-pay-challenge/internal/infra/apperr"
)

type apiHandler func(w http.ResponseWriter, r *http.Request) error

func MakeHandler(hdl apiHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := hdl(w, r); err != nil {
			var httpError *apperr.HttpError

			if ok := errors.As(err, &httpError); ok {
				WriteJSON(w, httpError.Code, httpError)
			} else {
				errResp := apperr.HttpError{
					Code:    http.StatusInternalServerError,
					Message: "internal server error",
				}
				WriteJSON(w, errResp.Code, errResp)
			}

			slog.Error("HTTP API error", "err", err.Error(), "path", r.URL.Path)
		}
	}
}
