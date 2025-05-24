package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserValidate_ValidCommonUser(t *testing.T) {
	cpf := "12345678901"
	user := User{
		Fullname: "John Doe",
		Role:     Common,
		CPF:      &cpf, // válido para o teste
		Email:    "john@example.com",
		Password: "password123",
	}
	err := user.Validate()
	assert.NoError(t, err)
}

func TestUserValidate_ValidShopkeeperUser(t *testing.T) {
	cnpj := "12345678000199"
	user := User{
		Fullname: "Jane Shop",
		Role:     Shopkeeper,
		CNPJ:     &cnpj, // válido para o teste
		Email:    "jane@shop.com",
		Password: "strongpass",
	}
	err := user.Validate()
	assert.NoError(t, err)
}

func TestUserValidate_InvalidFullname(t *testing.T) {
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
}

func TestUserValidate_InvalidRole(t *testing.T) {
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
}

func TestUserValidate_InvalidCPFForCommonUser(t *testing.T) {
	cpf := "12345abc901"
	user := User{
		Fullname: "John Doe",
		Role:     Common,
		CPF:      &cpf, // inválido
		Email:    "john@example.com",
		Password: "password123",
	}
	err := user.Validate()
	assert.Error(t, err)
}

func TestUserValidate_InvalidCNPJForShopkeeperUser(t *testing.T) {
	cnpj := "12a45678000199"
	user := User{
		Fullname: "Jane Shop",
		Role:     Shopkeeper,
		CNPJ:     &cnpj, // inválido
		Email:    "jane@shop.com",
		Password: "strongpass",
	}
	err := user.Validate()
	assert.Error(t, err)
}

func TestUserValidate_InvalidEmail(t *testing.T) {
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
}

func TestUserValidate_InvalidPassword(t *testing.T) {
	cpf := "12345678901"
	user := User{
		Fullname: "John Doe",
		Role:     Common,
		CPF:      &cpf,
		Email:    "john@example.com",
		Password: "123", // senha muito curta
	}
	err := user.Validate()
	assert.Error(t, err)
}

func TestIsValidFullname_WithValidName(t *testing.T) {
	val := "John Doe"
	err := isValidFullname(val)
	assert.NoError(t, err)
}

func TestIsValidFullname_WithNameTooShort(t *testing.T) {
	val := "Jo"
	err := isValidFullname(val)
	assert.Error(t, err)
}

func TestIsValidFullname_WithEmptyString(t *testing.T) {
	val := ""
	err := isValidFullname(val)
	assert.Error(t, err)
}

func TestIsValidFullname_WithSpacesOnly(t *testing.T) {
	val := "   "
	err := isValidFullname(val)
	assert.Error(t, err)
}

func TestIsValidRole_WithUserCommon(t *testing.T) {
	role := Common
	err := isValidRole(role)
	assert.NoError(t, err)
}

func TestIsValidRole_WithUserShopkeeper(t *testing.T) {
	role := Shopkeeper
	err := isValidRole(role)
	assert.NoError(t, err)
}

func TestIsValidRole_WithInvalidRole(t *testing.T) {
	var invalidRole UserRole = "ADMIN" // supondo que ADMIN não é válido
	err := isValidRole(invalidRole)
	assert.Error(t, err)
}

func TestIsValidCPF_WithValidInput(t *testing.T) {
	val := "12345678901"
	err := isValidCPF(val)
	assert.NoError(t, err)
}

func TestIsValidCPF_WithLessThan11Digits(t *testing.T) {
	val := "1234567890" // 10 digits
	err := isValidCPF(val)
	assert.Error(t, err)
}

func TestIsValidCPF_WithMoreThan11Digits(t *testing.T) {
	val := "123456789012" // 12 digits
	err := isValidCPF(val)
	assert.Error(t, err)
}

func TestIsValidCPF_WithNonNumericCharacters(t *testing.T) {
	val := "12345abc901"
	err := isValidCPF(val)
	assert.Error(t, err)
}

func TestIsValidCNPJ_WithValidInput(t *testing.T) {
	val := "06532946000185"
	err := isValidCNPJ(val)
	assert.NoError(t, err)
}

func TestIsValidCNPJ_WithInvalidChars(t *testing.T) {
	val := "06532946000abc"
	err := isValidCNPJ(val)
	assert.Error(t, err)
}

func TestIsValidCNPJ_WithInvalidLength(t *testing.T) {
	val := "17.901.294/0001-25"
	err := isValidCNPJ(val)
	assert.Error(t, err)
}

func TestIsValidEmail_WithInvalidFormat(t *testing.T) {
	val := "invalid@email..com"
	err := isValidEmail(val)
	assert.Error(t, err)
}

func TestIsValidEmail_WithValidFormat(t *testing.T) {
	val := "valid@email.com"
	err := isValidEmail(val)
	assert.NoError(t, err)
}

func TestIsValidPassword_WithTooShortPassword(t *testing.T) {
	val := "12345"
	err := isValidPassword(val)
	assert.Error(t, err)
}

func TestIsValidPassword_WithMinimumLength(t *testing.T) {
	val := "123456"
	err := isValidPassword(val)
	assert.NoError(t, err)
}
