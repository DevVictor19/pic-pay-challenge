package mocks

import (
	"context"

	"github.com/DevVictor19/pic-pay-challenge/internal/domain/user"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateCommon(ctx context.Context, dto user.CommonUserDTO) (int, error) {
	args := m.Called(ctx, dto)
	return args.Int(0), args.Error(1)
}

func (m *MockUserService) CreateShopkeeper(ctx context.Context, dto user.ShopkeeperUserDTO) (int, error) {
	args := m.Called(ctx, dto)
	return args.Int(0), args.Error(1)
}

func (m *MockUserService) FindByCPF(ctx context.Context, cpf string) (*user.User, error) {
	args := m.Called(ctx, cpf)
	u, ok := args.Get(0).(*user.User)
	if !ok && args.Get(0) != nil {
		panic("expected *User or nil")
	}
	return u, args.Error(1)
}

func (m *MockUserService) FindByCNPJ(ctx context.Context, cnpj string) (*user.User, error) {
	args := m.Called(ctx, cnpj)
	u, ok := args.Get(0).(*user.User)
	if !ok && args.Get(0) != nil {
		panic("expected *User or nil")
	}
	return u, args.Error(1)
}

func (m *MockUserService) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	args := m.Called(ctx, email)
	u, ok := args.Get(0).(*user.User)
	if !ok && args.Get(0) != nil {
		panic("expected *User or nil")
	}
	return u, args.Error(1)
}

func (m *MockUserService) FindByID(ctx context.Context, id int) (*user.User, error) {
	args := m.Called(ctx, id)
	u, ok := args.Get(0).(*user.User)
	if !ok && args.Get(0) != nil {
		panic("expected *User or nil")
	}
	return u, args.Error(1)
}
