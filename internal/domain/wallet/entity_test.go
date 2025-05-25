package wallet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWalletValidation(t *testing.T) {
	t.Run("Validate with valid data", func(t *testing.T) {
		wallet := Wallet{
			UserID:  1,
			Balance: 1000,
		}

		err := wallet.Validate()

		assert.NoError(t, err)
	})

	t.Run("Validate with invalid UserID", func(t *testing.T) {
		wallet := Wallet{
			UserID:  0,
			Balance: 1000,
		}

		err := wallet.Validate()

		assert.Error(t, err)
	})

	t.Run("Validate with negative balance", func(t *testing.T) {
		wallet := Wallet{
			UserID:  1,
			Balance: -500,
		}

		err := wallet.Validate()

		assert.Error(t, err)
	})
}

func TestIsValidUserID(t *testing.T) {
	t.Run("with valid ID", func(t *testing.T) {
		userID := 1
		err := isValidUserID(userID)
		assert.NoError(t, err)
	})

	t.Run("with zero ID", func(t *testing.T) {
		userID := 0
		err := isValidUserID(userID)
		assert.Error(t, err)
	})

	t.Run("with negative ID", func(t *testing.T) {
		userID := -10
		err := isValidUserID(userID)
		assert.Error(t, err)
	})
}

func TestIsValidBalance(t *testing.T) {
	t.Run("with positive balance", func(t *testing.T) {
		balance := int64(100)
		err := isValidBalance(balance)
		assert.NoError(t, err)
	})

	t.Run("with zero balance", func(t *testing.T) {
		balance := int64(0)
		err := isValidBalance(balance)
		assert.NoError(t, err)
	})

	t.Run("with negative balance", func(t *testing.T) {
		balance := int64(-50)
		err := isValidBalance(balance)
		assert.Error(t, err)
	})
}
