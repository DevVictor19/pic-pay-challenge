package accounting

type WalletRepository interface {
	Save(w Wallet) error
}
