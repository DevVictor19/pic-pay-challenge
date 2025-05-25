package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/mock"
)

// MockJWTService é um mock gerado para a interface JWTService
type MockJWTService struct {
	mock.Mock
}

func (m *MockJWTService) GenerateToken(userId int, exp time.Duration) (string, error) {
	args := m.Called(userId, exp)
	return args.String(0), args.Error(1)
}

func (m *MockJWTService) ValidateToken(token string) (*jwt.Token, error) {
	args := m.Called(token)
	return args.Get(0).(*jwt.Token), args.Error(1)
}
