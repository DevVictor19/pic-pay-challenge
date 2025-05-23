package wallet

import (
	"net/http"
	"time"

	"github.com/DevVictor19/pic-pay-challenge/internal/infra/apperr"
)

type WalletService interface {
	Create(userID int, balance int64) error
}

type walletSvc struct {
	wallRepo *WalletRepository
}

func (s *walletSvc) Create(userID int, balance int64) error {
	wallRepo := *s.wallRepo

	now := time.Now()
	wall := Wallet{
		UserID:    userID,
		Active:    true,
		Balance:   balance,
		UpdatedAt: now,
		CreatedAt: now,
	}

	if err := wall.Validate(); err != nil {
		return apperr.NewHttpError(http.StatusUnprocessableEntity, err.Error())
	}

	if err := wallRepo.Save(wall); err != nil {
		return err
	}

	return nil
}

var walletSvcRef *walletSvc

func NewWalletService(wallRepo *WalletRepository) WalletService {
	if walletSvcRef == nil {
		walletSvcRef = &walletSvc{
			wallRepo,
		}
	}
	return walletSvcRef
}
