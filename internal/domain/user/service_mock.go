package user

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateCommon(ctx context.Context, dto CommonUserDTO) (int, error) {
	args := m.Called(ctx, dto)
	return args.Int(0), args.Error(1)
}

func (m *MockUserService) CreateShopkeeper(ctx context.Context, dto ShopkeeperUserDTO) (int, error) {
	args := m.Called(ctx, dto)
	return args.Int(0), args.Error(1)
}

func (m *MockUserService) FindByCPF(ctx context.Context, cpf string) (*User, error) {
	args := m.Called(ctx, cpf)
	u, ok := args.Get(0).(*User)
	if !ok && args.Get(0) != nil {
		panic("expected *User or nil")
	}
	return u, args.Error(1)
}

func (m *MockUserService) FindByCNPJ(ctx context.Context, cnpj string) (*User, error) {
	args := m.Called(ctx, cnpj)
	u, ok := args.Get(0).(*User)
	if !ok && args.Get(0) != nil {
		panic("expected *User or nil")
	}
	return u, args.Error(1)
}

func (m *MockUserService) FindByEmail(ctx context.Context, email string) (*User, error) {
	args := m.Called(ctx, email)
	u, ok := args.Get(0).(*User)
	if !ok && args.Get(0) != nil {
		panic("expected *User or nil")
	}
	return u, args.Error(1)
}

func (m *MockUserService) FindByID(ctx context.Context, id int) (*User, error) {
	args := m.Called(ctx, id)
	u, ok := args.Get(0).(*User)
	if !ok && args.Get(0) != nil {
		panic("expected *User or nil")
	}
	return u, args.Error(1)
}
