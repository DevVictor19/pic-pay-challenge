package accounting

import (
	"errors"
	"time"
)

type UserService interface {
	Create(role UserRole, fullname, cpf, email, password string) error
	FindByCPF(cpf string) (*User, error)
}

type UserSvc struct {
	userRepo *UserRepository
	wallSvc  *WalletService
}

func (s *UserSvc) Create(role UserRole, fullname, cpf, email, password string) error {
	usrRepo := *s.userRepo
	wallSvc := *s.wallSvc

	withSameEmail, err := usrRepo.FindByEmail(email)
	if err != nil {
		// TODO: melhorar tratamento de erros
		return err
	}

	if withSameEmail != nil {
		// TODO: melhorar tratamento de erros
		return errors.New("email já cadastrado")
	}

	withSameCpf, err := usrRepo.FindByCPF(cpf)
	if err != nil {
		// TODO: melhorar tratamento de erros
		return err
	}

	if withSameCpf != nil {
		// TODO: melhorar tratamento de erros
		return errors.New("cpf já cadastrado")
	}

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

	// TODO: validar dados do user antes de salvar
	userId, err := usrRepo.Save(user)
	if err != nil {
		// TODO: melhorar tratamento de erros
		return err
	}

	if err := wallSvc.Create(userId, 0); err != nil {
		// TODO: melhorar tratamento de erros
		return err
	}

	return nil
}

func (s *UserSvc) FindByCPF(cpf string) (*User, error) {
	usrRepo := *s.userRepo

	usr, err := usrRepo.FindByCPF(cpf)
	if err != nil {
		// TODO: melhorar tratamento de erros
		return nil, err
	}

	if usr == nil {
		// TODO: melhorar tratamento de erros
		return nil, errors.New("usuário não encontrado")
	}

	return usr, nil
}

var userSvc *UserSvc

func NewUserService(userRepo *UserRepository, wallSvc *WalletService) UserService {
	if userSvc == nil {
		userSvc = &UserSvc{
			userRepo,
			wallSvc,
		}
	}
	return userSvc
}
