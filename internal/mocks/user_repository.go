package mocks

import (
	"context"

	"github.com/DevVictor19/pic-pay-challenge/internal/domain/user"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Save(ctx context.Context, u user.User) (int, error) {
	args := m.Called(ctx, u)
	return args.Int(0), args.Error(1)
}

func (m *MockUserRepository) FindByCPF(ctx context.Context, cpf string) (*user.User, error) {
	args := m.Called(ctx, cpf)
	if u, ok := args.Get(0).(*user.User); ok {
		return u, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) FindByCNPJ(ctx context.Context, cnpj string) (*user.User, error) {
	args := m.Called(ctx, cnpj)
	if u, ok := args.Get(0).(*user.User); ok {
		return u, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	args := m.Called(ctx, email)
	if u, ok := args.Get(0).(*user.User); ok {
		return u, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) FindByID(ctx context.Context, id int) (*user.User, error) {
	args := m.Called(ctx, id)
	if u, ok := args.Get(0).(*user.User); ok {
		return u, args.Error(1)
	}
	return nil, args.Error(1)
}
