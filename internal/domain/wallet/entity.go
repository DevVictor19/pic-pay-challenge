package wallet

import (
	"errors"
	"time"
)

type Wallet struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Active    bool      `json:"active"`
	Balance   int64     `json:"balance"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func (w *Wallet) Validate() error {
	if err := isValidUserID(w.UserID); err != nil {
		return err
	}
	if err := isValidBalance(w.Balance); err != nil {
		return err
	}
	return nil
}

func isValidUserID(userID int) error {
	if userID <= 0 {
		return errors.New("user id must be greater than 0")
	}
	return nil
}

func isValidBalance(balance int64) error {
	if balance < 0 {
		return errors.New("balance must be equal or greater than 0")
	}
	return nil
}
