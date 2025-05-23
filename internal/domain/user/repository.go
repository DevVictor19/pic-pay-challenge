package user

import "context"

type UserRepository interface {
	Save(ctx context.Context, u User) (int, error)
	FindByCPF(ctx context.Context, cpf string) (*User, error)
	FindByCNPJ(ctx context.Context, cnpj string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByID(ctx context.Context, id string) (*User, error)
}
