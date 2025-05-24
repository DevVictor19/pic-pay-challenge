package auth

type SignupDTO struct {
	Fullname string  `json:"fullname"`
	CPF      *string `json:"cpf"`
	CNPJ     *string `json:"cnpj"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
}

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
