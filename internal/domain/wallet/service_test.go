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
	mockRepo.On("Save", mock.Anything, mock.Anything).Return(nil)

	service := NewWalletService(mockRepo)

	err := service.Create(context.Background(), 1, 1000)
	assert.NoError(t, err)

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

	err := service.Create(context.Background(), 0, -1000)
	var httpError *apperr.HttpError
	assert.ErrorAs(t, err, &httpError)
	assert.EqualValues(t, httpError.Code, http.StatusUnprocessableEntity)

	mockRepo.AssertExpectations(t)
}
