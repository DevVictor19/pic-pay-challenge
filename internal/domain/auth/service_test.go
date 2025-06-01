package auth

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/DevVictor19/pic-pay-challenge/internal/domain/user"
	"github.com/DevVictor19/pic-pay-challenge/internal/domain/wallet"
	"github.com/DevVictor19/pic-pay-challenge/internal/infra/apperror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthService_Signup(t *testing.T) {
	t.Run("should return bad request if cpnj and cpf is nil", func(t *testing.T) {
		userServiceMock := new(user.MockUserService)
		userServiceMock.AssertNotCalled(t, "FindByEmail")
		userServiceMock.AssertNotCalled(t, "FindByCPF")
		userServiceMock.AssertNotCalled(t, "FindByCNPJ")
		userServiceMock.AssertNotCalled(t, "CreateCommon")
		userServiceMock.AssertNotCalled(t, "CreateShopkeeper")

		wallServiceMock := new(wallet.MockWalletService)
		wallServiceMock.AssertNotCalled(t, "Create")

		bcryptServiceMock := new(MockBcryptService)
		bcryptServiceMock.AssertNotCalled(t, "Hash")

		jwtServiceMock := new(MockJWTService)

		dto := SignupDTO{
			Fullname: "John Doe",
			Email:    "email@email.com",
			Password: "password123",
		}

		service := NewAuthService(
			userServiceMock,
			wallServiceMock,
			bcryptServiceMock,
			jwtServiceMock,
		)

		err := service.Signup(context.Background(), dto)

		assert.Error(t, err)
		var httpError *apperror.HttpError
		assert.ErrorAs(t, err, &httpError)
		assert.Equal(t, httpError.Code, http.StatusBadRequest)

		userServiceMock.AssertExpectations(t)
		wallServiceMock.AssertExpectations(t)
		bcryptServiceMock.AssertExpectations(t)
		jwtServiceMock.AssertExpectations(t)
	})

	t.Run("should return bad request if cpnj and cpf is not nil", func(t *testing.T) {
		userServiceMock := new(user.MockUserService)
		userServiceMock.AssertNotCalled(t, "FindByEmail")
		userServiceMock.AssertNotCalled(t, "FindByCPF")
		userServiceMock.AssertNotCalled(t, "FindByCNPJ")
		userServiceMock.AssertNotCalled(t, "CreateCommon")
		userServiceMock.AssertNotCalled(t, "CreateShopkeeper")

		wallServiceMock := new(wallet.MockWalletService)
		wallServiceMock.AssertNotCalled(t, "Create")

		bcryptServiceMock := new(MockBcryptService)
		bcryptServiceMock.AssertNotCalled(t, "Hash")

		jwtServiceMock := new(MockJWTService)

		cpf := "12345678990"
		cnpj := "12345678912345"
		dto := SignupDTO{
			Fullname: "John Doe",
			CPF:      &cpf,
			CNPJ:     &cnpj,
			Email:    "email@email.com",
			Password: "password123",
		}

		service := NewAuthService(
			userServiceMock,
			wallServiceMock,
			bcryptServiceMock,
			jwtServiceMock,
		)

		err := service.Signup(context.Background(), dto)

		assert.Error(t, err)
		var httpError *apperror.HttpError
		assert.ErrorAs(t, err, &httpError)
		assert.Equal(t, http.StatusBadRequest, httpError.Code)

		userServiceMock.AssertExpectations(t)
		wallServiceMock.AssertExpectations(t)
		bcryptServiceMock.AssertExpectations(t)
		jwtServiceMock.AssertExpectations(t)
	})

	t.Run("should return error if FindByEmail returns a generic error", func(t *testing.T) {
		userServiceMock := new(user.MockUserService)
		userServiceMock.On("FindByEmail", mock.Anything, mock.Anything).
			Return(nil, errors.New("generic error"))
		userServiceMock.AssertNotCalled(t, "FindByCPF")
		userServiceMock.AssertNotCalled(t, "FindByCNPJ")
		userServiceMock.AssertNotCalled(t, "CreateCommon")
		userServiceMock.AssertNotCalled(t, "CreateShopkeeper")

		wallServiceMock := new(wallet.MockWalletService)
		wallServiceMock.AssertNotCalled(t, "Create")

		bcryptServiceMock := new(MockBcryptService)
		bcryptServiceMock.AssertNotCalled(t, "Hash")

		jwtServiceMock := new(MockJWTService)

		cpf := "12345678990"
		dto := SignupDTO{
			Fullname: "John Doe",
			CPF:      &cpf,
			Email:    "email@email.com",
			Password: "password123",
		}

		service := NewAuthService(
			userServiceMock,
			wallServiceMock,
			bcryptServiceMock,
			jwtServiceMock,
		)

		err := service.Signup(context.Background(), dto)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "generic error")

		userServiceMock.AssertExpectations(t)
		wallServiceMock.AssertExpectations(t)
		bcryptServiceMock.AssertExpectations(t)
		jwtServiceMock.AssertExpectations(t)
	})

	t.Run("should return conflict if finds a user with same email", func(t *testing.T) {
		userServiceMock := new(user.MockUserService)
		userServiceMock.On("FindByEmail", mock.Anything, mock.Anything).
			Return(&user.User{}, nil)
		userServiceMock.AssertNotCalled(t, "FindByCPF")
		userServiceMock.AssertNotCalled(t, "FindByCNPJ")
		userServiceMock.AssertNotCalled(t, "CreateCommon")
		userServiceMock.AssertNotCalled(t, "CreateShopkeeper")

		wallServiceMock := new(wallet.MockWalletService)
		wallServiceMock.AssertNotCalled(t, "Create")

		bcryptServiceMock := new(MockBcryptService)
		bcryptServiceMock.AssertNotCalled(t, "Hash")

		jwtServiceMock := new(MockJWTService)

		cpf := "12345678990"
		dto := SignupDTO{
			Fullname: "John Doe",
			CPF:      &cpf,
			Email:    "email@email.com",
			Password: "password123",
		}

		service := NewAuthService(
			userServiceMock,
			wallServiceMock,
			bcryptServiceMock,
			jwtServiceMock,
		)

		err := service.Signup(context.Background(), dto)

		assert.Error(t, err)
		var httpError *apperror.HttpError
		assert.ErrorAs(t, err, &httpError)
		assert.Equal(t, http.StatusConflict, httpError.Code)

		userServiceMock.AssertExpectations(t)
		wallServiceMock.AssertExpectations(t)
		bcryptServiceMock.AssertExpectations(t)
		jwtServiceMock.AssertExpectations(t)
	})

	t.Run("should return error if FindByCNPJ returns a generic error", func(t *testing.T) {
		userServiceMock := new(user.MockUserService)
		userServiceMock.On("FindByEmail", mock.Anything, mock.Anything).
			Return(nil, nil)
		userServiceMock.On("FindByCNPJ", mock.Anything, mock.Anything).
			Return(nil, errors.New("generic error"))
		userServiceMock.AssertNotCalled(t, "FindByCPF")
		userServiceMock.AssertNotCalled(t, "CreateCommon")
		userServiceMock.AssertNotCalled(t, "CreateShopkeeper")

		wallServiceMock := new(wallet.MockWalletService)
		wallServiceMock.AssertNotCalled(t, "Create")

		bcryptServiceMock := new(MockBcryptService)
		bcryptServiceMock.AssertNotCalled(t, "Hash")

		jwtServiceMock := new(MockJWTService)

		cnpj := "12345678901234"
		dto := SignupDTO{
			Fullname: "John Doe",
			CNPJ:     &cnpj,
			Email:    "email@email.com",
			Password: "password123",
		}

		service := NewAuthService(
			userServiceMock,
			wallServiceMock,
			bcryptServiceMock,
			jwtServiceMock,
		)

		err := service.Signup(context.Background(), dto)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "generic error")

		userServiceMock.AssertExpectations(t)
		wallServiceMock.AssertExpectations(t)
		bcryptServiceMock.AssertExpectations(t)
		jwtServiceMock.AssertExpectations(t)
	})

	t.Run("should return conflict if finds a user with same cnpj", func(t *testing.T) {
		userServiceMock := new(user.MockUserService)
		userServiceMock.On("FindByEmail", mock.Anything, mock.Anything).
			Return(nil, nil)
		userServiceMock.On("FindByCNPJ", mock.Anything, mock.Anything).
			Return(&user.User{}, nil)
		userServiceMock.AssertNotCalled(t, "FindByCPF")
		userServiceMock.AssertNotCalled(t, "CreateCommon")
		userServiceMock.AssertNotCalled(t, "CreateShopkeeper")

		wallServiceMock := new(wallet.MockWalletService)
		wallServiceMock.AssertNotCalled(t, "Create")

		bcryptServiceMock := new(MockBcryptService)
		bcryptServiceMock.AssertNotCalled(t, "Hash")

		jwtServiceMock := new(MockJWTService)

		cnpj := "12345678901234"
		dto := SignupDTO{
			Fullname: "John Doe",
			CNPJ:     &cnpj,
			Email:    "email@email.com",
			Password: "password123",
		}

		service := NewAuthService(
			userServiceMock,
			wallServiceMock,
			bcryptServiceMock,
			jwtServiceMock,
		)

		err := service.Signup(context.Background(), dto)

		assert.Error(t, err)
		var httpError *apperror.HttpError
		assert.ErrorAs(t, err, &httpError)
		assert.Equal(t, http.StatusConflict, httpError.Code)

		userServiceMock.AssertExpectations(t)
		wallServiceMock.AssertExpectations(t)
		bcryptServiceMock.AssertExpectations(t)
		jwtServiceMock.AssertExpectations(t)
	})

	t.Run("should return generic error if FindByCPF returns a generic error", func(t *testing.T) {
		userServiceMock := new(user.MockUserService)
		userServiceMock.On("FindByEmail", mock.Anything, mock.Anything).
			Return(nil, nil)
		userServiceMock.AssertNotCalled(t, "FindByCNPJ")
		userServiceMock.On("FindByCPF", mock.Anything, mock.Anything).
			Return(nil, errors.New("generic error"))
		userServiceMock.AssertNotCalled(t, "CreateCommon")
		userServiceMock.AssertNotCalled(t, "CreateShopkeeper")

		wallServiceMock := new(wallet.MockWalletService)
		wallServiceMock.AssertNotCalled(t, "Create")

		bcryptServiceMock := new(MockBcryptService)
		bcryptServiceMock.AssertNotCalled(t, "Hash")

		jwtServiceMock := new(MockJWTService)

		cpf := "12345678901"
		dto := SignupDTO{
			Fullname: "John Doe",
			CPF:      &cpf,
			Email:    "email@email.com",
			Password: "password123",
		}

		service := NewAuthService(
			userServiceMock,
			wallServiceMock,
			bcryptServiceMock,
			jwtServiceMock,
		)

		err := service.Signup(context.Background(), dto)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "generic error")

		userServiceMock.AssertExpectations(t)
		wallServiceMock.AssertExpectations(t)
		bcryptServiceMock.AssertExpectations(t)
		jwtServiceMock.AssertExpectations(t)
	})

	t.Run("should return conflict if finds a user with same cpf", func(t *testing.T) {
		userServiceMock := new(user.MockUserService)
		userServiceMock.On("FindByEmail", mock.Anything, mock.Anything).
			Return(nil, nil)
		userServiceMock.AssertNotCalled(t, "FindByCNPJ")
		userServiceMock.On("FindByCPF", mock.Anything, mock.Anything).
			Return(&user.User{}, nil)
		userServiceMock.AssertNotCalled(t, "CreateCommon")
		userServiceMock.AssertNotCalled(t, "CreateShopkeeper")

		wallServiceMock := new(wallet.MockWalletService)
		wallServiceMock.AssertNotCalled(t, "Create")

		bcryptServiceMock := new(MockBcryptService)
		bcryptServiceMock.AssertNotCalled(t, "Hash")

		jwtServiceMock := new(MockJWTService)

		cpf := "12345678901"
		dto := SignupDTO{
			Fullname: "John Doe",
			CPF:      &cpf,
			Email:    "email@email.com",
			Password: "password123",
		}

		service := NewAuthService(
			userServiceMock,
			wallServiceMock,
			bcryptServiceMock,
			jwtServiceMock,
		)

		err := service.Signup(context.Background(), dto)

		assert.Error(t, err)
		var httpError *apperror.HttpError
		assert.ErrorAs(t, err, &httpError)
		assert.Equal(t, http.StatusConflict, httpError.Code)

		userServiceMock.AssertExpectations(t)
		wallServiceMock.AssertExpectations(t)
		bcryptServiceMock.AssertExpectations(t)
		jwtServiceMock.AssertExpectations(t)
	})

	t.Run("should create a Shopkeeper user if cnpj is present", func(t *testing.T) {
		ctx := context.Background()
		cnpj := "12345678000199"
		signupDto := SignupDTO{
			Fullname: "Jane Doe",
			CNPJ:     &cnpj,
			Email:    "jane@email.com",
			Password: "password456",
		}
		userId := 2

		userServiceMock := new(user.MockUserService)
		userServiceMock.On("FindByEmail", ctx, signupDto.Email).Return(nil, nil)
		userServiceMock.AssertNotCalled(t, "FindByCPF")
		userServiceMock.On("FindByCNPJ", ctx, *signupDto.CNPJ).Return(nil, nil)

		bcryptServiceMock := new(MockBcryptService)
		bcryptServiceMock.On("Hash", signupDto.Password).Return("hashed456", nil)

		shopkeeperDto := user.ShopkeeperUserDTO{
			Fullname: signupDto.Fullname,
			CNPJ:     signupDto.CNPJ,
			Email:    signupDto.Email,
			Password: "hashed456",
		}

		userServiceMock.On("CreateShopkeeper", ctx, shopkeeperDto).Return(userId, nil)
		userServiceMock.AssertNotCalled(t, "CreateCommon")

		wallServiceMock := new(wallet.MockWalletService)
		wallServiceMock.On("Create", ctx, userId, int64(0)).Return(nil)

		jwtServiceMock := new(MockJWTService)

		service := NewAuthService(
			userServiceMock,
			wallServiceMock,
			bcryptServiceMock,
			jwtServiceMock,
		)

		err := service.Signup(ctx, signupDto)

		assert.NoError(t, err)

		userServiceMock.AssertExpectations(t)
		wallServiceMock.AssertExpectations(t)
		bcryptServiceMock.AssertExpectations(t)
		jwtServiceMock.AssertExpectations(t)
	})

	t.Run("should create a Common user if cpf is present", func(t *testing.T) {
		ctx := context.Background()
		cpf := "12345678901"
		signupDto := SignupDTO{
			Fullname: "John Doe",
			CPF:      &cpf,
			Email:    "email@email.com",
			Password: "password123",
		}
		userId := 1

		userServiceMock := new(user.MockUserService)
		userServiceMock.On("FindByEmail", ctx, signupDto.Email).Return(nil, nil)
		userServiceMock.AssertNotCalled(t, "FindByCNPJ")
		userServiceMock.On("FindByCPF", ctx, *signupDto.CPF).Return(nil, nil)

		wallServiceMock := new(wallet.MockWalletService)
		wallServiceMock.On("Create", ctx, userId, int64(0)).Return(nil)

		bcryptServiceMock := new(MockBcryptService)
		bcryptServiceMock.On("Hash", signupDto.Password).Return("hashed", nil)

		commonDto := user.CommonUserDTO{
			Fullname: signupDto.Fullname,
			CPF:      signupDto.CPF,
			Email:    signupDto.Email,
			Password: "hashed",
		}

		userServiceMock.On("CreateCommon", ctx, commonDto).Return(userId, nil)
		userServiceMock.AssertNotCalled(t, "CreateShopkeeper")

		jwtServiceMock := new(MockJWTService)

		service := NewAuthService(
			userServiceMock,
			wallServiceMock,
			bcryptServiceMock,
			jwtServiceMock,
		)

		err := service.Signup(ctx, signupDto)

		assert.NoError(t, err)

		userServiceMock.AssertExpectations(t)
		wallServiceMock.AssertExpectations(t)
		bcryptServiceMock.AssertExpectations(t)
		jwtServiceMock.AssertExpectations(t)
	})
}

func TestAuthService_Login(t *testing.T) {
	t.Run("should return unauthorized if user with email is not found", func(t *testing.T) {
		userServiceMock := new(user.MockUserService)
		bcryptServiceMock := new(MockBcryptService)
		jwtServiceMock := new(MockJWTService)

		dto := LoginDTO{
			Email:    "notfound@example.com",
			Password: "password123",
		}

		userServiceMock.On("FindByEmail", mock.Anything, dto.Email).
			Return(nil, apperror.NewHttpError(http.StatusNotFound, "user not found"))

		service := NewAuthService(userServiceMock, nil, bcryptServiceMock, jwtServiceMock)

		token, err := service.Login(context.Background(), dto)

		assert.Error(t, err)
		assert.Empty(t, token)

		userServiceMock.AssertExpectations(t)
	})

	t.Run("should return unauthorized if password compare returns false", func(t *testing.T) {
		userServiceMock := new(user.MockUserService)
		bcryptServiceMock := new(MockBcryptService)
		jwtServiceMock := new(MockJWTService)

		dto := LoginDTO{
			Email:    "user@example.com",
			Password: "wrongpassword",
		}

		user := &user.User{
			ID:       1,
			Password: "hashedpassword",
		}

		userServiceMock.On("FindByEmail", mock.Anything, dto.Email).
			Return(user, nil)
		bcryptServiceMock.On("Compare", dto.Password, user.Password).
			Return(false)

		service := NewAuthService(userServiceMock, nil, bcryptServiceMock, jwtServiceMock)

		token, err := service.Login(context.Background(), dto)

		assert.Error(t, err)
		assert.Empty(t, token)

		userServiceMock.AssertExpectations(t)
		bcryptServiceMock.AssertExpectations(t)
	})

	t.Run("should return a token if pass all validations", func(t *testing.T) {
		userServiceMock := new(user.MockUserService)
		bcryptServiceMock := new(MockBcryptService)
		jwtServiceMock := new(MockJWTService)

		dto := LoginDTO{
			Email:    "user@example.com",
			Password: "correctpassword",
		}

		user := &user.User{
			ID:       1,
			Password: "hashedpassword",
		}

		userServiceMock.On("FindByEmail", mock.Anything, dto.Email).
			Return(user, nil)
		bcryptServiceMock.On("Compare", dto.Password, user.Password).
			Return(true)
		jwtServiceMock.On("GenerateToken", user.ID, mock.Anything).
			Return("generated-token", nil)

		service := NewAuthService(userServiceMock, nil, bcryptServiceMock, jwtServiceMock)

		token, err := service.Login(context.Background(), dto)

		assert.NoError(t, err)
		assert.Equal(t, "generated-token", token)

		userServiceMock.AssertExpectations(t)
		bcryptServiceMock.AssertExpectations(t)
		jwtServiceMock.AssertExpectations(t)
	})
}
