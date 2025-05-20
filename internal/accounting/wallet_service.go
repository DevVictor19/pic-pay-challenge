package accounting

type WalletService interface {
	Create(userID int, balance int64) error
}

type WalletSvc struct {
}

func (s *WalletSvc) Create(userID int, balance int64) error {
	return nil
}

var walletSvc *WalletSvc

func NewWalletService(wallRepo *WalletRepository) WalletService {
	if walletSvc == nil {
		walletSvc = &WalletSvc{}
	}
	return walletSvc
}
