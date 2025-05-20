package wallet

type WalletService interface {
	Create(userID int, balance int64) error
}

type walletSvc struct {
}

func (s *walletSvc) Create(userID int, balance int64) error {
	return nil
}

var walletSvcRef *walletSvc

func NewWalletService(wallRepo *WalletRepository) WalletService {
	if walletSvcRef == nil {
		walletSvcRef = &walletSvc{}
	}
	return walletSvcRef
}
