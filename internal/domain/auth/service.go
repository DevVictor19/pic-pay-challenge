package auth

import (
	"context"
	"net/http"

	"github.com/DevVictor19/pic-pay-challenge/internal/domain/user"
	"github.com/DevVictor19/pic-pay-challenge/internal/domain/wallet"
	"github.com/DevVictor19/pic-pay-challenge/internal/infra/apperr"
)

type AuthService interface {
	Signup(ctx context.Context, dto SignupDTO) error
	Login(ctx context.Context, dto LoginDTO) (string, error)
}

type authSvc struct {
	userService   *user.UserService
	wallService   *wallet.WalletService
	bcryptService *BcryptService
}

func (s *authSvc) Signup(ctx context.Context, dto SignupDTO) error {
	userService := *s.userService
	wallService := *s.wallService
	bcryptService := *s.bcryptService

	if dto.CNPJ == nil && dto.CPF == nil {
		return apperr.NewHttpError(http.StatusBadRequest, "CPF or CNPJ must be passed")
	}

	if dto.CNPJ != nil && dto.CPF != nil {
		return apperr.NewHttpError(http.StatusBadRequest, "chose CPF or CNPJ for create a new user")
	}

	withSameEmail, err := userService.FindByEmail(ctx, dto.Email)
	if !apperr.IsHttpError(err) {
		return err
	}

	if withSameEmail != nil {
		return apperr.NewHttpError(http.StatusConflict, "email already in use")
	}

	if dto.CNPJ != nil {
		withSameCnpj, err := userService.FindByCNPJ(ctx, *dto.CNPJ)
		if !apperr.IsHttpError(err) {
			return err
		}

		if withSameCnpj != nil {
			return apperr.NewHttpError(http.StatusConflict, "cnpj already in use")
		}
	}

	if dto.CPF != nil {
		withSameCpf, err := userService.FindByCPF(ctx, *dto.CPF)
		if !apperr.IsHttpError(err) {
			return err
		}

		if withSameCpf != nil {
			return apperr.NewHttpError(http.StatusConflict, "cpf already in use")
		}
	}

	hashed, err := bcryptService.Hash(dto.Password)
	if err != nil {
		return err
	}

	var userId int

	if dto.CNPJ != nil {
		id, err := userService.CreateShopkeeper(ctx, user.ShopkeeperUserDTO{
			Fullname: dto.Fullname,
			CNPJ:     *dto.CNPJ,
			Email:    dto.Email,
			Password: hashed,
		})
		if err != nil {
			return err
		}
		userId = id
	}

	if dto.CPF != nil {
		id, err := userService.CreateCommon(ctx, user.CommonUserDTO{
			Fullname: dto.Fullname,
			CPF:      *dto.CPF,
			Email:    dto.Email,
			Password: hashed,
		})
		if err != nil {
			return err
		}
		userId = id
	}

	err = wallService.Create(ctx, userId, 0)
	if err != nil {
		return err
	}

	return nil
}

func (s *authSvc) Login(ctx context.Context, dto LoginDTO) (string, error) {
	return "", nil
}

var authSvcRef *authSvc

func NewAuthService(usrSvc *user.UserService, wSvc *wallet.WalletService, bcrSvc *BcryptService) AuthService {
	if authSvcRef == nil {
		authSvcRef = &authSvc{
			userService:   usrSvc,
			wallService:   wSvc,
			bcryptService: bcrSvc,
		}
	}
	return authSvcRef
}
