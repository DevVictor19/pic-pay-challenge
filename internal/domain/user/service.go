package user

import (
	"errors"
	"time"
)

type UserService interface {
	Create(role UserRole, fullname, cpf, email, password string) (int, error)
	FindByCPF(cpf string) (*User, error)
}

type userSvc struct {
	userRepo *UserRepository
}

func (s *userSvc) Create(role UserRole, fullname, cpf, email, password string) (int, error) {
	usrRepo := *s.userRepo

	now := time.Now()

	user := User{
		Fullname:  fullname,
		Role:      role,
		CPF:       cpf,
		Email:     email,
		Password:  password,
		UpdatedAt: now,
		CreatedAt: now,
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
		return nil, errors.New("usuário não encontrado")
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
