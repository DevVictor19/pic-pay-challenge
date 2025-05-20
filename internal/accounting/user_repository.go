package accounting

type userRepository interface {
	save(u User) (int, error)
	findByCPF(cpf string) (*User, error)
	findByEmail(email string) (*User, error)
}
