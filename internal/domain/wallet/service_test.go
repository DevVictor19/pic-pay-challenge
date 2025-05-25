package wallet

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/DevVictor19/pic-pay-challenge/internal/infra/apperr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestWalletService_Create(t *testing.T) {
	t.Run("should create wallet successfully", func(t *testing.T) {
		mockRepo := new(MockWalletRepository)

		var entity Wallet
		userId := 123
		balance := int64(1000)

		mockRepo.On("Save", mock.Anything, mock.Anything).
			Run(func(args mock.Arguments) {
				entity = args.Get(1).(Wallet)
			}).
			Return(nil).Once()

		service := NewWalletService(mockRepo)

		err := service.Create(context.Background(), userId, balance)
		assert.NoError(t, err)
		assert.Equal(t, userId, entity.UserID)
		assert.Equal(t, balance, entity.Balance)

		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when db fails", func(t *testing.T) {
		mockRepo := new(MockWalletRepository)
		mockRepo.On("Save", mock.Anything, mock.Anything).
			Return(errors.New("db fail")).Once()

		service := NewWalletService(mockRepo)

		err := service.Create(context.Background(), 1, 1000)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "db fail")

		mockRepo.AssertExpectations(t)
	})

	t.Run("should return validation error for invalid userID", func(t *testing.T) {
		mockRepo := new(MockWalletRepository)

		service := NewWalletService(mockRepo)

		err := service.Create(context.Background(), 0, 1000)
		var httpError *apperr.HttpError
		assert.ErrorAs(t, err, &httpError)
		assert.Equal(t, http.StatusUnprocessableEntity, httpError.Code)

		mockRepo.AssertExpectations(t)
	})

	t.Run("should return validation error for invalid balance", func(t *testing.T) {
		mockRepo := new(MockWalletRepository)

		service := NewWalletService(mockRepo)

		err := service.Create(context.Background(), 1, -1000)
		var httpError *apperr.HttpError
		assert.ErrorAs(t, err, &httpError)
		assert.Equal(t, http.StatusUnprocessableEntity, httpError.Code)

		mockRepo.AssertExpectations(t)
	})
}
