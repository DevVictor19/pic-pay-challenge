package user

import (
	"fmt"
	"time"

	"github.com/DevVictor19/pic-pay-challenge/internal/infra/http_errors"
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
		Role:      USER_COMMON,
		CPF:       cpf,
		Email:     email,
		Password:  password,
		UpdatedAt: now,
		CreatedAt: now,
	}

	err := user.Validate()
	if err != nil {
		return 0, fmt.Errorf("invalid user data: %w", http_errors.ErrBadRequest)
	}

	userId, err := usrRepo.Save(user)
	if err != nil {
		return 0, fmt.Errorf("error saving user: %w", http_errors.ErrInternal)
	}

	return userId, nil
}

func (s *userSvc) CreateShopkeeper(fullname, cnpj, email, password string) (int, error) {
	usrRepo := *s.userRepo

	now := time.Now()

	user := User{
		Fullname:  fullname,
		Role:      USER_SHOPKEEPER,
		CNPJ:      cnpj,
		Email:     email,
		Password:  password,
		UpdatedAt: now,
		CreatedAt: now,
	}

	err := user.Validate()
	if err != nil {
		return 0, fmt.Errorf("invalid user data: %w", http_errors.ErrBadRequest)
	}

	userId, err := usrRepo.Save(user)
	if err != nil {
		return 0, fmt.Errorf("error saving user: %w", http_errors.ErrInternal)
	}

	return userId, nil
}

func (s *userSvc) FindByCPF(cpf string) (*User, error) {
	usrRepo := *s.userRepo

	usr, err := usrRepo.FindByCPF(cpf)
	if err != nil {
		return nil, fmt.Errorf("error finding user by cpf: %w", http_errors.ErrInternal)
	}

	if usr == nil {
		return nil, fmt.Errorf("user not found: %w", http_errors.ErrNotFound)
	}

	return usr, nil
}

func (s *userSvc) FindByCNPJ(cnpj string) (*User, error) {
	usrRepo := *s.userRepo

	usr, err := usrRepo.FindByCNPJ(cnpj)
	if err != nil {
		return nil, fmt.Errorf("error finding user by cnpj: %w", http_errors.ErrInternal)
	}

	if usr == nil {
		return nil, fmt.Errorf("user not found: %w", http_errors.ErrNotFound)
	}

	return usr, nil
}

func (s *userSvc) FindByEmail(email string) (*User, error) {
	usrRepo := *s.userRepo

	usr, err := usrRepo.FindByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("error finding user by email: %w", http_errors.ErrInternal)
	}

	if usr == nil {
		return nil, fmt.Errorf("user not found: %w", http_errors.ErrNotFound)
	}

	return usr, nil
}

func (s *userSvc) FindByID(id string) (*User, error) {
	usrRepo := *s.userRepo

	usr, err := usrRepo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("error finding user by id: %w", http_errors.ErrInternal)
	}

	if usr == nil {
		return nil, fmt.Errorf("user not found: %w", http_errors.ErrNotFound)
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
