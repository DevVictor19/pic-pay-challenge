package auth

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/DevVictor19/pic-pay-challenge/internal/domain/user"
	"github.com/DevVictor19/pic-pay-challenge/internal/domain/wallet"
	"github.com/DevVictor19/pic-pay-challenge/internal/infra/apperr"
)

type AuthService interface {
	Signup(ctx context.Context, dto SignupDTO) error
	Login(ctx context.Context, dto LoginDTO) (string, error)
}

type authSvc struct {
	userService   user.UserService
	wallService   wallet.WalletService
	bcryptService BcryptService
	jwtService    JWTService
}

func (s *authSvc) Signup(ctx context.Context, dto SignupDTO) error {
	if dto.CNPJ == nil && dto.CPF == nil {
		return apperr.NewHttpError(http.StatusBadRequest, "CPF or CNPJ must be passed")
	}

	if dto.CNPJ != nil && dto.CPF != nil {
		return apperr.NewHttpError(http.StatusBadRequest, "chose CPF or CNPJ for create a new user")
	}

	withSameEmail, err := s.userService.FindByEmail(ctx, dto.Email)
	if err != nil {
		var httpError *apperr.HttpError
		if ok := errors.As(err, &httpError); !ok {
			return err
		}
	}

	if withSameEmail != nil {
		return apperr.NewHttpError(http.StatusConflict, "email already in use")
	}

	if dto.CNPJ != nil {
		withSameCnpj, err := s.userService.FindByCNPJ(ctx, *dto.CNPJ)
		if err != nil {
			var httpError *apperr.HttpError
			if ok := errors.As(err, &httpError); !ok {
				return err
			}
		}

		if withSameCnpj != nil {
			return apperr.NewHttpError(http.StatusConflict, "cnpj already in use")
		}
	}

	if dto.CPF != nil {
		withSameCpf, err := s.userService.FindByCPF(ctx, *dto.CPF)
		if err != nil {
			var httpError *apperr.HttpError
			if ok := errors.As(err, &httpError); !ok {
				return err
			}
		}

		if withSameCpf != nil {
			return apperr.NewHttpError(http.StatusConflict, "cpf already in use")
		}
	}

	hashed, err := s.bcryptService.Hash(dto.Password)
	if err != nil {
		return err
	}

	var userId int

	if dto.CNPJ != nil {
		id, err := s.userService.CreateShopkeeper(ctx, user.ShopkeeperUserDTO{
			Fullname: dto.Fullname,
			CNPJ:     dto.CNPJ,
			Email:    dto.Email,
			Password: hashed,
		})
		if err != nil {
			return err
		}
		userId = id
	}

	if dto.CPF != nil {
		id, err := s.userService.CreateCommon(ctx, user.CommonUserDTO{
			Fullname: dto.Fullname,
			CPF:      dto.CPF,
			Email:    dto.Email,
			Password: hashed,
		})
		if err != nil {
			return err
		}
		userId = id
	}

	err = s.wallService.Create(ctx, userId, 0)
	if err != nil {
		return err
	}

	return nil
}

func (s *authSvc) Login(ctx context.Context, dto LoginDTO) (string, error) {
	user, err := s.userService.FindByEmail(ctx, dto.Email)
	if err != nil {
		var httpError *apperr.HttpError
		if ok := errors.As(err, &httpError); ok {
			return "", apperr.NewHttpError(http.StatusUnauthorized, "invalid email or password")
		}

		return "", err
	}

	isValidPwd := s.bcryptService.Compare(dto.Password, user.Password)
	if !isValidPwd {
		return "", apperr.NewHttpError(http.StatusUnauthorized, "invalid email or password")
	}

	token, err := s.jwtService.GenerateToken(user.ID, time.Minute*30)
	if err != nil {
		return "", err
	}

	return token, nil
}

func NewAuthService(
	usrSvc user.UserService,
	wSvc wallet.WalletService,
	bcrSvc BcryptService,
	jwtSvc JWTService) AuthService {

	return &authSvc{
		userService:   usrSvc,
		wallService:   wSvc,
		bcryptService: bcrSvc,
		jwtService:    jwtSvc,
	}
}
