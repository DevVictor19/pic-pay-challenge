package accounting

import "time"

type UserRole string

const (
	USER_COMMON     UserRole = "common"
	USER_SHOPKEEPER UserRole = "shopkeeper"
)

type User struct {
	ID        int       `json:"id"`
	Fullname  string    `json:"fullname"`
	Role      UserRole  `json:"role"`
	CPF       string    `json:"cpf"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
