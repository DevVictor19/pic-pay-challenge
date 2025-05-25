package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockWalletService struct {
	mock.Mock
}

func (m *MockWalletService) Create(ctx context.Context, userID int, balance int64) error {
	args := m.Called(ctx, userID, balance)
	return args.Error(0)
}
