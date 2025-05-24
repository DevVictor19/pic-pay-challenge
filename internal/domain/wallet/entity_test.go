package wallet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWalletValidate_WithValidData(t *testing.T) {
	wallet := Wallet{
		UserID:  1,
		Balance: 1000,
	}

	err := wallet.Validate()

	assert.NoError(t, err)
}

func TestWalletValidate_WithInvalidUserID(t *testing.T) {
	wallet := Wallet{
		UserID:  0,
		Balance: 1000,
	}

	err := wallet.Validate()

	assert.Error(t, err)
}

func TestWalletValidate_WithNegativeBalance(t *testing.T) {
	wallet := Wallet{
		UserID:  1,
		Balance: -500,
	}

	err := wallet.Validate()

	assert.Error(t, err)
}

func TestIsValidUserID_WithValidID(t *testing.T) {
	userID := 1
	err := isValidUserID(userID)
	assert.NoError(t, err)
}

func TestIsValidUserID_WithZeroID(t *testing.T) {
	userID := 0
	err := isValidUserID(userID)
	assert.Error(t, err)
}

func TestIsValidUserID_WithNegativeID(t *testing.T) {
	userID := -10
	err := isValidUserID(userID)
	assert.Error(t, err)
}

func TestIsValidBalance_WithPositiveBalance(t *testing.T) {
	balance := int64(100)
	err := isValidBalance(balance)
	assert.NoError(t, err)
}

func TestIsValidBalance_WithZeroBalance(t *testing.T) {
	balance := int64(0)
	err := isValidBalance(balance)
	assert.NoError(t, err)
}

func TestIsValidBalance_WithNegativeBalance(t *testing.T) {
	balance := int64(-50)
	err := isValidBalance(balance)
	assert.Error(t, err)
}
