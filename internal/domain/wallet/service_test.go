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

type MockWalletRepository struct {
	mock.Mock
}

func (m *MockWalletRepository) Save(ctx context.Context, wall Wallet) error {
	args := m.Called(ctx, wall)
	return args.Error(0)
}

func TestWalletService_Create(t *testing.T) {
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
}

func TestWalletService_CreateWithDBFail(t *testing.T) {
	mockRepo := new(MockWalletRepository)
	mockRepo.On("Save", mock.Anything, mock.Anything).Return(errors.New("db fail"))

	service := NewWalletService(mockRepo)

	err := service.Create(context.Background(), 1, 1000)
	assert.Error(t, err)

	mockRepo.AssertExpectations(t)
}

func TestWalletService_CreateWithInvalidUserID(t *testing.T) {
	mockRepo := new(MockWalletRepository)

	service := NewWalletService(mockRepo)

	err := service.Create(context.Background(), 0, 1000)
	var httpError *apperr.HttpError
	assert.ErrorAs(t, err, &httpError)
	assert.EqualValues(t, httpError.Code, http.StatusUnprocessableEntity)

	mockRepo.AssertExpectations(t)
}

func TestWalletService_CreateWithInvalidBalance(t *testing.T) {
	mockRepo := new(MockWalletRepository)

	service := NewWalletService(mockRepo)

	err := service.Create(context.Background(), 1, -1000)
	var httpError *apperr.HttpError
	assert.ErrorAs(t, err, &httpError)
	assert.EqualValues(t, httpError.Code, http.StatusUnprocessableEntity)

	mockRepo.AssertExpectations(t)
}
