package entities

import "time"

type TransactionType string

const (
	TRAN_PAYMENT_RECEIVED TransactionType = "payment_received"
	TRAN_PAYMENT_SENT     TransactionType = "payment_sent"
)

type Transaction struct {
	ID          int             `json:"id"`
	PayerID     int             `json:"payer_id"`
	PayeeID     int             `json:"payee_id"`
	Type        TransactionType `json:"type"`
	Amount      int64           `json:"amount"`
	Description string          `json:"description"`
	UpdatedAt   time.Time       `json:"updated_at"`
	CreatedAt   time.Time       `json:"created_at"`
}
