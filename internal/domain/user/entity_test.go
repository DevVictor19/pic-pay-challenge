package user

import (
	"testing"
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
	if err != nil {
		t.Errorf("expected no error for valid common user, got: %v", err)
	}
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
	if err != nil {
		t.Errorf("expected no error for valid shopkeeper user, got: %v", err)
	}
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
	if err == nil {
		t.Error("expected error for invalid fullname")
	}
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
	if err == nil {
		t.Error("expected error for invalid role")
	}
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
	if err == nil {
		t.Error("expected error for invalid CPF in common user")
	}
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
	if err == nil {
		t.Error("expected error for invalid CNPJ in shopkeeper user")
	}
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
	if err == nil {
		t.Error("expected error for invalid email")
	}
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
	if err == nil {
		t.Error("expected error for invalid password")
	}
}

func TestIsValidFullname_WithValidName(t *testing.T) {
	val := "John Doe"
	err := isValidFullname(val)
	if err != nil {
		t.Error("should not return error for valid fullname")
	}
}

func TestIsValidFullname_WithNameTooShort(t *testing.T) {
	val := "Jo"
	err := isValidFullname(val)
	if err == nil {
		t.Error("should return error if fullname is less than 3 characters")
	}
}

func TestIsValidFullname_WithEmptyString(t *testing.T) {
	val := ""
	err := isValidFullname(val)
	if err == nil {
		t.Error("should return error if fullname is empty")
	}
}

func TestIsValidFullname_WithSpacesOnly(t *testing.T) {
	val := "   "
	err := isValidFullname(val)
	if err == nil {
		t.Error("should return error if fullname contains only spaces")
	}
}

func TestIsValidRole_WithUserCommon(t *testing.T) {
	role := Common
	err := isValidRole(role)
	if err != nil {
		t.Error("should not return error for Common role")
	}
}

func TestIsValidRole_WithUserShopkeeper(t *testing.T) {
	role := Shopkeeper
	err := isValidRole(role)
	if err != nil {
		t.Error("should not return error for Shopkeeper role")
	}
}

func TestIsValidRole_WithInvalidRole(t *testing.T) {
	var invalidRole UserRole = "ADMIN" // supondo que ADMIN não é válido
	err := isValidRole(invalidRole)
	if err == nil {
		t.Error("should return error for invalid role")
	}
}

func TestIsValidCPF_WithValidInput(t *testing.T) {
	val := "12345678901"
	result := isValidCPF(val)
	if result != nil {
		t.Error("should not return error if CPF has exactly 11 numeric digits")
	}
}

func TestIsValidCPF_WithLessThan11Digits(t *testing.T) {
	val := "1234567890" // 10 digits
	result := isValidCPF(val)
	if result == nil {
		t.Error("should return error if CPF has less than 11 digits")
	}
}

func TestIsValidCPF_WithMoreThan11Digits(t *testing.T) {
	val := "123456789012" // 12 digits
	result := isValidCPF(val)
	if result == nil {
		t.Error("should return error if CPF has more than 11 digits")
	}
}

func TestIsValidCPF_WithNonNumericCharacters(t *testing.T) {
	val := "12345abc901"
	result := isValidCPF(val)
	if result == nil {
		t.Error("should return error if CPF contains non-numeric characters")
	}
}

func TestIsValidCNPJ_WithValidInput(t *testing.T) {
	val := "06532946000185"
	result := isValidCNPJ(val)
	if result != nil {
		t.Error("should not return error if CNPJ is valid")
	}
}

func TestIsValidCNPJ_WithInvalidChars(t *testing.T) {
	val := "06532946000abc"
	result := isValidCNPJ(val)
	if result == nil {
		t.Error("should return error if CNPJ is invalid")
	}
}

func TestIsValidCNPJ_WithInvalidLength(t *testing.T) {
	val := "17.901.294/0001-25"
	result := isValidCNPJ(val)
	if result == nil {
		t.Error("should return error if CNPJ is invalid")
	}
}

func TestIsValidEmail_WithInvalidFormat(t *testing.T) {
	val := "invalid@email..com"
	result := isValidEmail(val)
	if result == nil {
		t.Error("should return error if email is invalid")
	}
}

func TestIsValidEmail_WithValidFormat(t *testing.T) {
	val := "valid@email.com"
	result := isValidEmail(val)
	if result != nil {
		t.Error("should not return error if email is valid")
	}
}

func TestIsValidPassword_WithTooShortPassword(t *testing.T) {
	val := "12345"
	result := isValidPassword(val)
	if result == nil {
		t.Error("should return error if password is at least 6 chars")
	}
}

func TestIsValidPassword_WithMinimumLength(t *testing.T) {
	val := "123456"
	result := isValidPassword(val)
	if result != nil {
		t.Error("should not return error if password is at least 6 chars")
	}
}
