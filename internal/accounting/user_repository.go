package accounting

type UserRepository interface {
	Save(u User) (int, error)
	FindByCPF(cpf string) (*User, error)
	FindByEmail(email string) (*User, error)
}
