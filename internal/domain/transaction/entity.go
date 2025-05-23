package transaction

import (
	"time"
)

type TransactionType string

const (
	PaymentReceived TransactionType = "payment_received"
	PaymentSent     TransactionType = "payment_sent"
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
