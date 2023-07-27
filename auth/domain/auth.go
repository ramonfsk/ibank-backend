package domain

import (
	"github.com/ramonfsk/ibank-backend/auth/errs"
)

type User struct {
	ID        string `db:"id"`
	Name      string `db:"name,omitempty"`
	Birthdate string `db:"birthdate,omitempty"`
	Password  string `db:"password"`
	Email     string `db:"email"`
	Document  string `db:"document,omitempty"`
	Phone     string `db:"phone,omitempty"`
	Status    int    `db:"status,omitempty"`
	IsAdmin   bool   `db:"is_admin,omitempty"`
}

type AuthRepository interface {
	FindBy(email string, password string) (*User, *errs.AppError)
}
