package mocks

import (
	"context"

	"github.com/DevVictor19/pic-pay-challenge/internal/domain/wallet"
	"github.com/stretchr/testify/mock"
)

type MockWalletRepository struct {
	mock.Mock
}

func (m *MockWalletRepository) Save(ctx context.Context, wall wallet.Wallet) error {
	args := m.Called(ctx, wall)
	return args.Error(0)
}
