package wallet

import "context"

type WalletRepository interface {
	Save(ctx context.Context, w Wallet) error
}
