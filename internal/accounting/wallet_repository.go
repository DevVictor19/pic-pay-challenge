package accounting

type walletRepository interface {
	save(w Wallet) error
}
