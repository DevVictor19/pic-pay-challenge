package wallet

import (
	"fmt"
	"time"

	"github.com/DevVictor19/pic-pay-challenge/internal/infra/http_errors"
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

	err := wallRepo.Save(wall)
	if err != nil {
		return fmt.Errorf("error saving wallet: %w", http_errors.ErrInternal)
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
