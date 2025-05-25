package user

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/DevVictor19/pic-pay-challenge/internal/infra/apperr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserService_CreateCommon(t *testing.T) {
	t.Run("should return unprocessable entity if validation fails", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		service := NewUserService(mockRepo)

		cpf := "12345678900"
		dto := CommonUserDTO{
			Fullname: "",
			CPF:      &cpf,
			Email:    "test@test.com",
			Password: "pass123",
		}

		id, err := service.CreateCommon(context.Background(), dto)

		assert.Error(t, err)
		var httpError *apperr.HttpError
		assert.ErrorAs(t, err, &httpError)
		assert.Equal(t, httpError.Code, http.StatusUnprocessableEntity)
		assert.Equal(t, 0, id)

		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error if database fails", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		mockRepo.
			On("Save", mock.Anything, mock.Anything).
			Return(0, errors.New("db fail"))

		service := NewUserService(mockRepo)

		cpf := "12345678900"
		dto := CommonUserDTO{
			Fullname: "John Doe",
			CPF:      &cpf,
			Email:    "test@test.com",
			Password: "pass123",
		}

		id, err := service.CreateCommon(context.Background(), dto)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "db fail")
		assert.Equal(t, 0, id)

		mockRepo.AssertExpectations(t)
	})

	t.Run("should create a common user and return id", func(t *testing.T) {
		mockRepo := new(MockUserRepository)

		var entity User

		mockRepo.On("Save", mock.Anything, mock.Anything).
			Run(func(args mock.Arguments) {
				entity = args.Get(1).(User)
			}).Return(0, nil)

		service := NewUserService(mockRepo)

		cpf := "12345678900"
		dto := CommonUserDTO{
			Fullname: "John Doe",
			CPF:      &cpf,
			Email:    "test@test.com",
			Password: "pass123",
		}

		id, err := service.CreateCommon(context.Background(), dto)

		assert.NoError(t, err)
		assert.Equal(t, dto.Fullname, entity.Fullname)
		assert.Equal(t, dto.CPF, entity.CPF)
		assert.Nil(t, entity.CNPJ)
		assert.Equal(t, dto.Email, entity.Email)
		assert.Equal(t, dto.Password, entity.Password)
		assert.Equal(t, entity.Role, Common)
		assert.Equal(t, 0, id)

		mockRepo.AssertExpectations(t)
	})
}

func TestUserService_CreateShopkeeper(t *testing.T) {
	t.Run("should return unprocessable entity if validation fails", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		service := NewUserService(mockRepo)

		cnpj := "12345678901234"
		dto := ShopkeeperUserDTO{
			Fullname: "",
			CNPJ:     &cnpj,
			Email:    "test@test.com",
			Password: "pass123",
		}

		id, err := service.CreateShopkeeper(context.Background(), dto)

		assert.Error(t, err)
		var httpError *apperr.HttpError
		assert.ErrorAs(t, err, &httpError)
		assert.Equal(t, httpError.Code, http.StatusUnprocessableEntity)
		assert.Equal(t, 0, id)

		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error if database fails", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		mockRepo.
			On("Save", mock.Anything, mock.Anything).
			Return(0, errors.New("db fail"))

		service := NewUserService(mockRepo)

		cnpj := "12345678901234"
		dto := ShopkeeperUserDTO{
			Fullname: "John Doe",
			CNPJ:     &cnpj,
			Email:    "test@test.com",
			Password: "pass123",
		}

		id, err := service.CreateShopkeeper(context.Background(), dto)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "db fail")
		assert.Equal(t, 0, id)

		mockRepo.AssertExpectations(t)
	})

	t.Run("should create a shopkeeper user and return id", func(t *testing.T) {
		mockRepo := new(MockUserRepository)

		var entity User

		mockRepo.On("Save", mock.Anything, mock.Anything).
			Run(func(args mock.Arguments) {
				entity = args.Get(1).(User)
			}).Return(0, nil)

		service := NewUserService(mockRepo)

		cnpj := "12345678901234"
		dto := ShopkeeperUserDTO{
			Fullname: "John Doe",
			CNPJ:     &cnpj,
			Email:    "test@test.com",
			Password: "pass123",
		}

		id, err := service.CreateShopkeeper(context.Background(), dto)

		assert.NoError(t, err)
		assert.Equal(t, dto.Fullname, entity.Fullname)
		assert.Equal(t, dto.CNPJ, entity.CNPJ)
		assert.Nil(t, entity.CPF)
		assert.Equal(t, dto.Email, entity.Email)
		assert.Equal(t, dto.Password, entity.Password)
		assert.Equal(t, entity.Role, Shopkeeper)
		assert.Equal(t, 0, id)

		mockRepo.AssertExpectations(t)
	})
}

func TestUserService_FindByCPF(t *testing.T) {
	t.Run("should return error if db fails", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		mockRepo.
			On("FindByCPF", mock.Anything, mock.Anything).
			Return(nil, errors.New("db fail"))

		service := NewUserService(mockRepo)
		user, err := service.FindByCPF(context.Background(), "123")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "db fail")
		assert.Nil(t, user)

		mockRepo.AssertExpectations(t)
	})

	t.Run("should return not found if user is nil", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		mockRepo.
			On("FindByCPF", mock.Anything, mock.Anything).
			Return(nil, nil)

		service := NewUserService(mockRepo)
		user, err := service.FindByCPF(context.Background(), "123")

		assert.Error(t, err)
		var httpError *apperr.HttpError
		assert.ErrorAs(t, err, &httpError)
		assert.Equal(t, httpError.Code, http.StatusNotFound)
		assert.Nil(t, user)

		mockRepo.AssertExpectations(t)
	})

	t.Run("should return the user", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		mockUser := &User{}
		mockRepo.
			On("FindByCPF", mock.Anything, mock.Anything).
			Return(mockUser, nil)

		service := NewUserService(mockRepo)
		user, err := service.FindByCPF(context.Background(), "123")

		assert.NoError(t, err)
		assert.Equal(t, mockUser, user)

		mockRepo.AssertExpectations(t)
	})
}

func TestUserService_FindByCNPJ(t *testing.T) {
	t.Run("should return error if db fails", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		mockRepo.
			On("FindByCNPJ", mock.Anything, mock.Anything).
			Return(nil, errors.New("db fail"))

		service := NewUserService(mockRepo)
		user, err := service.FindByCNPJ(context.Background(), "123")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "db fail")
		assert.Nil(t, user)

		mockRepo.AssertExpectations(t)
	})

	t.Run("should return not found if user is nil", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		mockRepo.
			On("FindByCNPJ", mock.Anything, mock.Anything).
			Return(nil, nil)

		service := NewUserService(mockRepo)
		user, err := service.FindByCNPJ(context.Background(), "123")

		assert.Error(t, err)
		var httpError *apperr.HttpError
		assert.ErrorAs(t, err, &httpError)
		assert.Equal(t, httpError.Code, http.StatusNotFound)
		assert.Nil(t, user)

		mockRepo.AssertExpectations(t)
	})

	t.Run("should return the user", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		mockUser := &User{}
		mockRepo.
			On("FindByCNPJ", mock.Anything, mock.Anything).
			Return(mockUser, nil)

		service := NewUserService(mockRepo)
		user, err := service.FindByCNPJ(context.Background(), "123")

		assert.NoError(t, err)
		assert.Equal(t, mockUser, user)

		mockRepo.AssertExpectations(t)
	})
}

func TestUserService_FindByEmail(t *testing.T) {
	t.Run("should return error if db fails", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		mockRepo.
			On("FindByEmail", mock.Anything, mock.Anything).
			Return(nil, errors.New("db fail"))

		service := NewUserService(mockRepo)
		user, err := service.FindByEmail(context.Background(), "john@example.com")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "db fail")
		assert.Nil(t, user)

		mockRepo.AssertExpectations(t)
	})

	t.Run("should return not found if user is nil", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		mockRepo.
			On("FindByEmail", mock.Anything, mock.Anything).
			Return(nil, nil)

		service := NewUserService(mockRepo)
		user, err := service.FindByEmail(context.Background(), "john@example.com")

		assert.Error(t, err)
		var httpError *apperr.HttpError
		assert.ErrorAs(t, err, &httpError)
		assert.Equal(t, httpError.Code, http.StatusNotFound)
		assert.Nil(t, user)

		mockRepo.AssertExpectations(t)
	})

	t.Run("should return the user", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		mockUser := &User{}
		mockRepo.
			On("FindByEmail", mock.Anything, mock.Anything).
			Return(mockUser, nil)

		service := NewUserService(mockRepo)
		user, err := service.FindByEmail(context.Background(), "john@example.com")

		assert.NoError(t, err)
		assert.Equal(t, mockUser, user)

		mockRepo.AssertExpectations(t)
	})
}

func TestUserService_FindByID(t *testing.T) {
	t.Run("should return error if db fails", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		mockRepo.
			On("FindByID", mock.Anything, mock.Anything).
			Return(nil, errors.New("db fail"))

		service := NewUserService(mockRepo)
		user, err := service.FindByID(context.Background(), 1)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "db fail")
		assert.Nil(t, user)

		mockRepo.AssertExpectations(t)
	})

	t.Run("should return not found if user is nil", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		mockRepo.
			On("FindByID", mock.Anything, mock.Anything).
			Return(nil, nil)

		service := NewUserService(mockRepo)
		user, err := service.FindByID(context.Background(), 1)

		assert.Error(t, err)
		var httpError *apperr.HttpError
		assert.ErrorAs(t, err, &httpError)
		assert.Equal(t, httpError.Code, http.StatusNotFound)
		assert.Nil(t, user)

		mockRepo.AssertExpectations(t)
	})

	t.Run("should return the user", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		mockUser := &User{}
		mockRepo.
			On("FindByID", mock.Anything, mock.Anything).
			Return(mockUser, nil)

		service := NewUserService(mockRepo)
		user, err := service.FindByID(context.Background(), 1)

		assert.NoError(t, err)
		assert.Equal(t, mockUser, user)

		mockRepo.AssertExpectations(t)
	})
}
