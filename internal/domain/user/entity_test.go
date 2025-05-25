package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserValidate(t *testing.T) {
	t.Run("Valid Common User", func(t *testing.T) {
		cpf := "12345678901"
		user := User{
			Fullname: "John Doe",
			Role:     Common,
			CPF:      &cpf,
			Email:    "john@example.com",
			Password: "password123",
		}
		err := user.Validate()
		assert.NoError(t, err)
	})

	t.Run("Valid Shopkeeper User", func(t *testing.T) {
		cnpj := "12345678000199"
		user := User{
			Fullname: "Jane Shop",
			Role:     Shopkeeper,
			CNPJ:     &cnpj,
			Email:    "jane@shop.com",
			Password: "strongpass",
		}
		err := user.Validate()
		assert.NoError(t, err)
	})

	t.Run("Invalid Fullname", func(t *testing.T) {
		cpf := "12345678901"
		user := User{
			Fullname: "Jo",
			Role:     Common,
			CPF:      &cpf,
			Email:    "john@example.com",
			Password: "password123",
		}
		err := user.Validate()
		assert.Error(t, err)
	})

	t.Run("Invalid Role", func(t *testing.T) {
		cpf := "12345678901"
		user := User{
			Fullname: "John Doe",
			Role:     "INVALID_ROLE",
			CPF:      &cpf,
			Email:    "john@example.com",
			Password: "password123",
		}
		err := user.Validate()
		assert.Error(t, err)
	})

	t.Run("Invalid CPF for Common User", func(t *testing.T) {
		cpf := "12345abc901"
		user := User{
			Fullname: "John Doe",
			Role:     Common,
			CPF:      &cpf,
			Email:    "john@example.com",
			Password: "password123",
		}
		err := user.Validate()
		assert.Error(t, err)
	})

	t.Run("Invalid CNPJ for Shopkeeper User", func(t *testing.T) {
		cnpj := "12a45678000199"
		user := User{
			Fullname: "Jane Shop",
			Role:     Shopkeeper,
			CNPJ:     &cnpj,
			Email:    "jane@shop.com",
			Password: "strongpass",
		}
		err := user.Validate()
		assert.Error(t, err)
	})

	t.Run("Invalid Email", func(t *testing.T) {
		cpf := "12345678901"
		user := User{
			Fullname: "John Doe",
			Role:     Common,
			CPF:      &cpf,
			Email:    "invalid-email",
			Password: "password123",
		}
		err := user.Validate()
		assert.Error(t, err)
	})

	t.Run("Invalid Password", func(t *testing.T) {
		cpf := "12345678901"
		user := User{
			Fullname: "John Doe",
			Role:     Common,
			CPF:      &cpf,
			Email:    "john@example.com",
			Password: "123",
		}
		err := user.Validate()
		assert.Error(t, err)
	})
}

func TestIsValidFullname(t *testing.T) {
	tests := []struct {
		name string
		val  string
		want bool
	}{
		{"Valid Name", "John Doe", true},
		{"Too Short", "Jo", false},
		{"Empty String", "", false},
		{"Spaces Only", "   ", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := isValidFullname(tt.val)
			if tt.want {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestIsValidRole(t *testing.T) {
	tests := []struct {
		name string
		role UserRole
		want bool
	}{
		{"Common Role", Common, true},
		{"Shopkeeper Role", Shopkeeper, true},
		{"Invalid Role", "ADMIN", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := isValidRole(tt.role)
			if tt.want {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestIsValidCPF(t *testing.T) {
	tests := []struct {
		name string
		val  string
		want bool
	}{
		{"Valid Input", "12345678901", true},
		{"Less Than 11 Digits", "1234567890", false},
		{"More Than 11 Digits", "123456789012", false},
		{"Non Numeric Characters", "12345abc901", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := isValidCPF(tt.val)
			if tt.want {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestIsValidCNPJ(t *testing.T) {
	tests := []struct {
		name string
		val  string
		want bool
	}{
		{"Valid Input", "06532946000185", true},
		{"Invalid Characters", "06532946000abc", false},
		{"Invalid Length", "17.901.294/0001-25", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := isValidCNPJ(tt.val)
			if tt.want {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		name string
		val  string
		want bool
	}{
		{"Valid Format", "valid@email.com", true},
		{"Invalid Format", "invalid@email..com", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := isValidEmail(tt.val)
			if tt.want {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestIsValidPassword(t *testing.T) {
	tests := []struct {
		name string
		val  string
		want bool
	}{
		{"Too Short", "12345", false},
		{"Minimum Length", "123456", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := isValidPassword(tt.val)
			if tt.want {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
