package user

import (
	"errors"
	"regexp"
	"strings"
	"time"
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
	if u.Role == USER_COMMON {
		err := isValidCPF(u.CPF)
		if err != nil {
			return err
		}
	}
	if u.Role == USER_SHOPKEEPER {
		err := isValidCNPJ(u.CNPJ)
		if err != nil {
			return err
		}
	}
	if err := isValidCPF(u.CPF); err != nil {
		return err
	}
	if err := isValidCNPJ(u.CNPJ); err != nil {
		return err
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
	if !strings.Contains(str, " ") {
		return errors.New("fullname must contain at least a space separating first and last name")
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
		if ch < '0' || ch > '9' {
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
		if ch < '0' || ch > '9' {
			return errors.New("CNPJ must contain only numeric digits")
		}
	}
	return nil
}

func isValidEmail(str string) error {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
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
