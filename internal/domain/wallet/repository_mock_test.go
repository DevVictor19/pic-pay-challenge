package wallet

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockWalletRepository struct {
	mock.Mock
}

func (m *MockWalletRepository) Save(ctx context.Context, wall Wallet) error {
	args := m.Called(ctx, wall)
	return args.Error(0)
}
