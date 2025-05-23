package user

import (
	"context"
	"net/http"
	"time"

	"github.com/DevVictor19/pic-pay-challenge/internal/infra/apperr"
)

type UserService interface {
	CreateCommon(ctx context.Context, dto CommonUserDTO) (int, error)
	CreateShopkeeper(ctx context.Context, dto ShopkeeperUserDTO) (int, error)
	FindByCPF(ctx context.Context, cpf string) (*User, error)
	FindByCNPJ(ctx context.Context, cnpj string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByID(ctx context.Context, id string) (*User, error)
}

type userSvc struct {
	userRepo *UserRepository
}

func (s *userSvc) CreateCommon(ctx context.Context, dto CommonUserDTO) (int, error) {
	usrRepo := *s.userRepo

	now := time.Now()

	user := User{
		Fullname:  dto.Fullname,
		Role:      Common,
		CPF:       dto.CPF,
		Email:     dto.Email,
		Password:  dto.Password,
		UpdatedAt: now,
		CreatedAt: now,
	}

	err := user.Validate()
	if err != nil {
		return 0, apperr.NewHttpError(http.StatusUnprocessableEntity, err.Error())
	}

	userId, err := usrRepo.Save(ctx, user)
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (s *userSvc) CreateShopkeeper(ctx context.Context, dto ShopkeeperUserDTO) (int, error) {
	usrRepo := *s.userRepo

	now := time.Now()

	user := User{
		Fullname:  dto.Fullname,
		Role:      Shopkeeper,
		CNPJ:      dto.CNPJ,
		Email:     dto.Email,
		Password:  dto.Password,
		UpdatedAt: now,
		CreatedAt: now,
	}

	err := user.Validate()
	if err != nil {
		return 0, apperr.NewHttpError(http.StatusUnprocessableEntity, err.Error())
	}

	userId, err := usrRepo.Save(ctx, user)
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (s *userSvc) FindByCPF(ctx context.Context, cpf string) (*User, error) {
	usrRepo := *s.userRepo

	usr, err := usrRepo.FindByCPF(ctx, cpf)
	if err != nil {
		return nil, err
	}

	if usr == nil {
		return nil, apperr.NewHttpError(http.StatusNotFound, "user not found")
	}

	return usr, nil
}

func (s *userSvc) FindByCNPJ(ctx context.Context, cnpj string) (*User, error) {
	usrRepo := *s.userRepo

	usr, err := usrRepo.FindByCNPJ(ctx, cnpj)
	if err != nil {
		return nil, err
	}

	if usr == nil {
		return nil, apperr.NewHttpError(http.StatusNotFound, "user not found")
	}

	return usr, nil
}

func (s *userSvc) FindByEmail(ctx context.Context, email string) (*User, error) {
	usrRepo := *s.userRepo

	usr, err := usrRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if usr == nil {
		return nil, apperr.NewHttpError(http.StatusNotFound, "user not found")
	}

	return usr, nil
}

func (s *userSvc) FindByID(ctx context.Context, id string) (*User, error) {
	usrRepo := *s.userRepo

	usr, err := usrRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if usr == nil {
		return nil, apperr.NewHttpError(http.StatusNotFound, "user not found")
	}

	return usr, nil
}

var userSvcRef *userSvc

func NewUserService(userRepo *UserRepository) UserService {
	if userSvcRef == nil {
		userSvcRef = &userSvc{
			userRepo,
		}
	}
	return userSvcRef
}
