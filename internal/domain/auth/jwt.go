package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	GenerateToken(userId int, exp time.Duration) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtSvc struct {
	secret string
	aud    string
	iss    string
}

func (a *jwtSvc) GenerateToken(userId int, exp time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(exp).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"iss": a.iss,
		"aud": a.aud,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(a.secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (a *jwtSvc) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}

		return []byte(a.secret), nil
	},
		jwt.WithExpirationRequired(),
		jwt.WithAudience(a.aud),
		jwt.WithIssuer(a.aud),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
	)
}

var jwtServiceRef *jwtSvc

func NewJWTService(secret, aud, iss string) JWTService {
	if jwtServiceRef == nil {
		jwtServiceRef = &jwtSvc{secret, iss, aud}
	}
	return jwtServiceRef
}
