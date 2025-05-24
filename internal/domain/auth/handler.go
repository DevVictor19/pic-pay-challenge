package auth

import (
	"net/http"

	"github.com/DevVictor19/pic-pay-challenge/internal/infra/apperr"
	"github.com/DevVictor19/pic-pay-challenge/internal/infra/utils"
)

type AuthHandler struct {
	authService *AuthService
}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) error {
	authService := *h.authService

	var body SignupDTO
	if err := utils.ReadJSON(w, r, &body); err != nil {
		return apperr.NewHttpError(http.StatusBadRequest, "error parsing body request")
	}

	if err := utils.Validate.Struct(body); err != nil {
		return apperr.NewHttpError(http.StatusBadRequest, err.Error())
	}

	if err := authService.Signup(r.Context(), body); err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)

	return nil
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) error {
	authService := *h.authService

	var body LoginDTO
	if err := utils.ReadJSON(w, r, &body); err != nil {
		return apperr.NewHttpError(http.StatusBadRequest, "error parsing body request")
	}

	if err := utils.Validate.Struct(body); err != nil {
		return apperr.NewHttpError(http.StatusBadRequest, err.Error())
	}

	token, err := authService.Login(r.Context(), body)
	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, map[string]string{
		"token": token,
	})
}

var authHdlRef *AuthHandler

func NewAuthHandler(authService *AuthService) *AuthHandler {
	if authHdlRef == nil {
		authHdlRef = &AuthHandler{
			authService,
		}
	}
	return authHdlRef
}
