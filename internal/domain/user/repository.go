package user

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type UserRepository interface {
	Save(ctx context.Context, u User) (int, error)
	FindByCPF(ctx context.Context, cpf string) (*User, error)
	FindByCNPJ(ctx context.Context, cnpj string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByID(ctx context.Context, id int) (*User, error)
}

type userRepo struct {
	database     *sql.DB
	queryTimeout time.Duration
}

func (r *userRepo) Save(ctx context.Context, u User) (int, error) {
	query := `
		INSERT INTO users (fullname, role, cpf, cnpj, email, password, updated_at, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`

	ctx, cancel := context.WithTimeout(ctx, r.queryTimeout)
	defer cancel()

	db := r.database

	var userId int
	err := db.QueryRowContext(
		ctx,
		query,
		u.Fullname, u.Role, u.CPF, u.CNPJ, u.Email, u.Password, u.UpdatedAt, u.CreatedAt,
	).Scan(&userId)
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (r *userRepo) FindByCPF(ctx context.Context, cpf string) (*User, error) {
	query := `
		SELECT id, fullname, role, cpf, cnpj, email, password, updated_at, created_at
		FROM users
		WHERE cpf = $1
	`

	ctx, cancel := context.WithTimeout(ctx, r.queryTimeout)
	defer cancel()

	row := r.database.QueryRowContext(ctx, query, cpf)

	var u User
	var cpfPtr, cnpjPtr *string

	err := row.Scan(
		&u.ID,
		&u.Fullname,
		&u.Role,
		&cpfPtr,
		&cnpjPtr,
		&u.Email,
		&u.Password,
		&u.UpdatedAt,
		&u.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	u.CPF = cpfPtr
	u.CNPJ = cnpjPtr

	return &u, nil
}

func (r *userRepo) FindByCNPJ(ctx context.Context, cnpj string) (*User, error) {
	query := `
		SELECT id, fullname, role, cpf, cnpj, email, password, updated_at, created_at
		FROM users
		WHERE cnpj = $1
	`

	ctx, cancel := context.WithTimeout(ctx, r.queryTimeout)
	defer cancel()

	row := r.database.QueryRowContext(ctx, query, cnpj)

	var u User
	var cpfPtr, cnpjPtr *string

	err := row.Scan(
		&u.ID,
		&u.Fullname,
		&u.Role,
		&cpfPtr,
		&cnpjPtr,
		&u.Email,
		&u.Password,
		&u.UpdatedAt,
		&u.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	u.CPF = cpfPtr
	u.CNPJ = cnpjPtr

	return &u, nil
}

func (r *userRepo) FindByEmail(ctx context.Context, email string) (*User, error) {
	query := `
		SELECT id, fullname, role, cpf, cnpj, email, password, updated_at, created_at
		FROM users
		WHERE email = $1
	`

	ctx, cancel := context.WithTimeout(ctx, r.queryTimeout)
	defer cancel()

	row := r.database.QueryRowContext(ctx, query, email)

	var u User
	var cpfPtr, cnpjPtr *string

	err := row.Scan(
		&u.ID,
		&u.Fullname,
		&u.Role,
		&cpfPtr,
		&cnpjPtr,
		&u.Email,
		&u.Password,
		&u.UpdatedAt,
		&u.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	u.CPF = cpfPtr
	u.CNPJ = cnpjPtr

	return &u, nil
}

func (r *userRepo) FindByID(ctx context.Context, id int) (*User, error) {
	query := `
		SELECT id, fullname, role, cpf, cnpj, email, password, updated_at, created_at
		FROM users
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, r.queryTimeout)
	defer cancel()

	row := r.database.QueryRowContext(ctx, query, id)

	var u User
	var cpfPtr, cnpjPtr *string

	err := row.Scan(
		&u.ID,
		&u.Fullname,
		&u.Role,
		&cpfPtr,
		&cnpjPtr,
		&u.Email,
		&u.Password,
		&u.UpdatedAt,
		&u.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	u.CPF = cpfPtr
	u.CNPJ = cnpjPtr

	return &u, nil
}

var userRepoRef *userRepo

func NewUserRepository(database *sql.DB, qt time.Duration) UserRepository {
	if userRepoRef == nil {
		userRepoRef = &userRepo{
			database:     database,
			queryTimeout: qt,
		}
	}
	return userRepoRef
}
