package user

import (
	"errors"
	"regexp"
	"strings"
	"time"
	"unicode"
)

type UserRole string

const (
	USER_COMMON     UserRole = "common"
	USER_SHOPKEEPER UserRole = "shopkeeper"
)

type User struct {
	ID        int       `json:"id"`
	Fullname  string    `json:"fullname"`
	Role      UserRole  `json:"role"`
	CPF       string    `json:"cpf"`
	CNPJ      string    `json:"cnpj"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) Validate() error {
	if err := isValidFullname(u.Fullname); err != nil {
		return err
	}
	if err := isValidRole(u.Role); err != nil {
		return err
	}
	switch u.Role {
	case USER_COMMON:
		if err := isValidCPF(u.CPF); err != nil {
			return err
		}
	case USER_SHOPKEEPER:
		if err := isValidCNPJ(u.CNPJ); err != nil {
			return err
		}
	default:
		return errors.New("unsupported user role")
	}
	if err := isValidEmail(u.Email); err != nil {
		return err
	}
	if err := isValidPassword(u.Password); err != nil {
		return err
	}
	return nil
}

func isValidFullname(str string) error {
	if len(strings.TrimSpace(str)) < 3 {
		return errors.New("fullname must be at least 3 characters")
	}
	return nil
}

func isValidRole(r UserRole) error {
	if r != USER_COMMON && r != USER_SHOPKEEPER {
		return errors.New("role must be common or shopkeeper")
	}
	return nil
}

func isValidCPF(str string) error {
	if len(str) != 11 {
		return errors.New("CPF must have exactly 11 digits")
	}
	for _, ch := range str {
		if !unicode.IsDigit(ch) {
			return errors.New("CPF must contain only numeric digits")
		}
	}
	return nil
}

func isValidCNPJ(str string) error {
	if len(str) != 14 {
		return errors.New("CNPJ must have exactly 14 digits")
	}
	for _, ch := range str {
		if !unicode.IsDigit(ch) {
			return errors.New("CNPJ must contain only numeric digits")
		}
	}
	return nil
}

func isValidEmail(str string) error {
	re := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
	if !re.MatchString(str) {
		return errors.New("invalid email format")
	}
	return nil
}

func isValidPassword(str string) error {
	if len(str) < 6 {
		return errors.New("password must be at least 6 characters long")
	}
	return nil
}
