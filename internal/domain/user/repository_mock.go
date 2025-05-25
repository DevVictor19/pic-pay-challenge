package user

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Save(ctx context.Context, u User) (int, error) {
	args := m.Called(ctx, u)
	return args.Int(0), args.Error(1)
}

func (m *MockUserRepository) FindByCPF(ctx context.Context, cpf string) (*User, error) {
	args := m.Called(ctx, cpf)
	if u, ok := args.Get(0).(*User); ok {
		return u, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) FindByCNPJ(ctx context.Context, cnpj string) (*User, error) {
	args := m.Called(ctx, cnpj)
	if u, ok := args.Get(0).(*User); ok {
		return u, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
	args := m.Called(ctx, email)
	if u, ok := args.Get(0).(*User); ok {
		return u, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) FindByID(ctx context.Context, id int) (*User, error) {
	args := m.Called(ctx, id)
	if u, ok := args.Get(0).(*User); ok {
		return u, args.Error(1)
	}
	return nil, args.Error(1)
}
