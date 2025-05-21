package user

type UserRepository interface {
	Save(u User) (int, error)
	FindByCPF(cpf string) (*User, error)
	FindByCNPJ(cnpj string) (*User, error)
	FindByEmail(email string) (*User, error)
	FindByID(id string) (*User, error)
}
