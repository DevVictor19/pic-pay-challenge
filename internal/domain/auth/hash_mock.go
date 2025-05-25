package auth

import (
	"github.com/stretchr/testify/mock"
)

// MockBcryptService é o mock para a interface BcryptService
type MockBcryptService struct {
	mock.Mock
}

func (m *MockBcryptService) Hash(pwd string) (string, error) {
	args := m.Called(pwd)
	return args.String(0), args.Error(1)
}

func (m *MockBcryptService) Compare(pwd, hash string) bool {
	args := m.Called(pwd, hash)
	return args.Bool(0)
}
