package auth

type SignupDTO struct {
	Fullname string  `json:"fullname" validate:"required,max=100"`
	CPF      *string `json:"cpf,omitempty" validate:"omitempty,len=11"`
	CNPJ     *string `json:"cnpj,omitempty" validate:"omitempty,len=14`
	Email    string  `json:"email" validate:"required,email,max=100"`
	Password string  `json:"password" validate:"required,gte=6,max=100"`
}

type LoginDTO struct {
	Email    string `json:"email" validate:"required,email,max=100"`
	Password string `json:"password" validate:"required,gte=6,max=100"`
}
