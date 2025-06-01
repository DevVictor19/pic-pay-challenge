package server

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/DevVictor19/pic-pay-challenge/internal/domain/auth"
	"github.com/DevVictor19/pic-pay-challenge/internal/domain/user"
	"github.com/DevVictor19/pic-pay-challenge/internal/infra/apperror"
	"github.com/DevVictor19/pic-pay-challenge/internal/infra/utils"
	"github.com/golang-jwt/jwt/v5"
)

type Middleware func(http.Handler) http.Handler

func MakeJWTAuthMiddleware(jwtService auth.JWTService, userService user.UserService) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				code := http.StatusUnauthorized
				utils.WriteJSON(w, code, apperror.NewHttpError(
					code,
					"Authorization header is missing",
				))
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				code := http.StatusUnauthorized
				utils.WriteJSON(w, code, apperror.NewHttpError(
					code,
					"Authorization header is malformed",
				))
				return
			}

			token := parts[1]
			jwtToken, err := jwtService.ValidateToken(token)
			if err != nil {
				code := http.StatusUnauthorized
				utils.WriteJSON(w, code, apperror.NewHttpError(
					code,
					err.Error(),
				))
				return
			}

			claims, _ := jwtToken.Claims.(jwt.MapClaims)

			userId, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["sub"]), 10, 32)
			if err != nil {
				code := http.StatusUnauthorized
				utils.WriteJSON(w, code, apperror.NewHttpError(
					code,
					err.Error(),
				))
				return
			}

			ctx := r.Context()

			user, err := userService.FindByID(ctx, int(userId))
			if err != nil {
				code := http.StatusUnauthorized
				utils.WriteJSON(w, code, apperror.NewHttpError(
					code,
					err.Error(),
				))
				return
			}

			ctx = context.WithValue(ctx, utils.UserKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
