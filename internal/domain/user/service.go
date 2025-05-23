package user

import (
	"net/http"
	"time"

	"github.com/DevVictor19/pic-pay-challenge/internal/infra/apperr"
)

type UserService interface {
	CreateCommon(fullname, cpf, email, password string) (int, error)
	CreateShopkeeper(fullname, cnpj, email, password string) (int, error)
	FindByCPF(cpf string) (*User, error)
	FindByCNPJ(cnpj string) (*User, error)
	FindByEmail(email string) (*User, error)
	FindByID(id string) (*User, error)
}

type userSvc struct {
	userRepo *UserRepository
}

func (s *userSvc) CreateCommon(fullname, cpf, email, password string) (int, error) {
	usrRepo := *s.userRepo

	now := time.Now()

	user := User{
		Fullname:  fullname,
		Role:      Common,
		CPF:       cpf,
		Email:     email,
		Password:  password,
		UpdatedAt: now,
		CreatedAt: now,
	}

	err := user.Validate()
	if err != nil {
		return 0, apperr.NewHttpError(http.StatusUnprocessableEntity, err.Error())
	}

	userId, err := usrRepo.Save(user)
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (s *userSvc) CreateShopkeeper(fullname, cnpj, email, password string) (int, error) {
	usrRepo := *s.userRepo

	now := time.Now()

	user := User{
		Fullname:  fullname,
		Role:      Shopkeeper,
		CNPJ:      cnpj,
		Email:     email,
		Password:  password,
		UpdatedAt: now,
		CreatedAt: now,
	}

	err := user.Validate()
	if err != nil {
		return 0, apperr.NewHttpError(http.StatusUnprocessableEntity, err.Error())
	}

	userId, err := usrRepo.Save(user)
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (s *userSvc) FindByCPF(cpf string) (*User, error) {
	usrRepo := *s.userRepo

	usr, err := usrRepo.FindByCPF(cpf)
	if err != nil {
		return nil, err
	}

	if usr == nil {
		return nil, apperr.NewHttpError(http.StatusNotFound, "user not found")
	}

	return usr, nil
}

func (s *userSvc) FindByCNPJ(cnpj string) (*User, error) {
	usrRepo := *s.userRepo

	usr, err := usrRepo.FindByCNPJ(cnpj)
	if err != nil {
		return nil, err
	}

	if usr == nil {
		return nil, apperr.NewHttpError(http.StatusNotFound, "user not found")
	}

	return usr, nil
}

func (s *userSvc) FindByEmail(email string) (*User, error) {
	usrRepo := *s.userRepo

	usr, err := usrRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if usr == nil {
		return nil, apperr.NewHttpError(http.StatusNotFound, "user not found")
	}

	return usr, nil
}

func (s *userSvc) FindByID(id string) (*User, error) {
	usrRepo := *s.userRepo

	usr, err := usrRepo.FindByID(id)
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
