package wallet

import "testing"

func TestWalletValidate_WithValidData(t *testing.T) {
	wallet := Wallet{
		UserID:  1,
		Balance: 1000,
	}

	err := wallet.Validate()
	if err != nil {
		t.Errorf("expected no error for valid wallet, got: %v", err)
	}
}

func TestWalletValidate_WithInvalidUserID(t *testing.T) {
	wallet := Wallet{
		UserID:  0,
		Balance: 1000,
	}

	err := wallet.Validate()
	if err == nil {
		t.Error("expected error for invalid user ID, got nil")
	}
}

func TestWalletValidate_WithNegativeBalance(t *testing.T) {
	wallet := Wallet{
		UserID:  1,
		Balance: -500,
	}

	err := wallet.Validate()
	if err == nil {
		t.Error("expected error for negative balance, got nil")
	}
}

func TestIsValidUserID_WithValidID(t *testing.T) {
	userID := 1
	err := isValidUserID(userID)
	if err != nil {
		t.Errorf("expected no error for valid user ID, got: %v", err)
	}
}

func TestIsValidUserID_WithZeroID(t *testing.T) {
	userID := 0
	err := isValidUserID(userID)
	if err == nil {
		t.Error("expected error for user ID equal to 0, got nil")
	}
}

func TestIsValidUserID_WithNegativeID(t *testing.T) {
	userID := -10
	err := isValidUserID(userID)
	if err == nil {
		t.Error("expected error for negative user ID, got nil")
	}
}

func TestIsValidBalance_WithPositiveBalance(t *testing.T) {
	balance := int64(100)
	err := isValidBalance(balance)
	if err != nil {
		t.Errorf("expected no error for positive balance, got: %v", err)
	}
}

func TestIsValidBalance_WithZeroBalance(t *testing.T) {
	balance := int64(0)
	err := isValidBalance(balance)
	if err != nil {
		t.Errorf("expected no error for zero balance, got: %v", err)
	}
}

func TestIsValidBalance_WithNegativeBalance(t *testing.T) {
	balance := int64(-50)
	err := isValidBalance(balance)
	if err == nil {
		t.Error("expected error for negative balance, got nil")
	}
}
