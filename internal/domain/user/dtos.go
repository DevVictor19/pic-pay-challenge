package user

type CommonUserDTO struct {
	Fullname string
	CPF      string
	Email    string
	Password string
}

type ShopkeeperUserDTO struct {
	Fullname string
	CNPJ     string
	Email    string
	Password string
}
