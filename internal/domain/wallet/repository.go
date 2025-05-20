package wallet

type WalletRepository interface {
	Save(w Wallet) error
}
