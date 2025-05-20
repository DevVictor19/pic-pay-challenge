package wallet

import "time"

type Wallet struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Active    bool      `json:"active"`
	Balance   int64     `json:"balance"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
