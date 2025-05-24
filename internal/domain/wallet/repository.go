package wallet

import (
	"context"
	"database/sql"
	"time"
)

type WalletRepository interface {
	Save(ctx context.Context, w Wallet) error
}

type walletRepo struct {
	database     *sql.DB
	queryTimeout time.Duration
}

func (r *walletRepo) Save(ctx context.Context, w Wallet) error {
	query := `
		INSERT INTO wallets (user_id, active, balance, updated_at, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	ctx, cancel := context.WithTimeout(ctx, r.queryTimeout)
	defer cancel()

	var walletID int
	err := r.database.QueryRowContext(
		ctx,
		query,
		w.UserID,
		w.Active,
		w.Balance,
		w.UpdatedAt,
		w.CreatedAt,
	).Scan(&walletID)

	if err != nil {
		return err
	}

	w.ID = walletID

	return nil
}

var wallRepoRef *walletRepo

func NewWalletRepository(database *sql.DB, qt time.Duration) WalletRepository {
	if wallRepoRef == nil {
		wallRepoRef = &walletRepo{
			database:     database,
			queryTimeout: qt,
		}
	}
	return wallRepoRef
}
