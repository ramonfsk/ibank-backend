package domain

import (
	"github.com/ramonfsk/ibank-backend/server/dto"
	"github.com/ramonfsk/ibank-backend/server/errs"
	"github.com/ramonfsk/ibank-backend/server/utils"
)

type User struct {
	ID        string `db:"id"`
	Name      string `db:"name"`
	Birthdate string `db:"birthdate"`
	Password  string `db:"password"`
	Email     string `db:"email"`
	Document  string `db:"document"`
	Phone     string `db:"phone"`
	Status    int    `db:"status"`
	IsAdmin   bool   `db:"is_admin"`
}

type UserRepository interface {
	FindAll(string) ([]User, *errs.AppError)
	FindByID(string) (*User, *errs.AppError)
	RegisterNewUser(dto.UserRequest, Account) (*Account, *errs.AppError)
}

func (u User) ToDTO() dto.UserResponse {
	return dto.UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Birthdate: u.Birthdate,
		Password:  u.Password,
		Email:     u.Email,
		Document:  u.Document,
		Phone:     u.Phone,
		Status:    utils.StatusAsText(u.Status),
		IsAdmin:   u.IsAdmin,
	}
}
