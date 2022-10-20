package domain

import (
	"github.ibm.com/rfnascimento/ibank/server/dto"
	"github.ibm.com/rfnascimento/ibank/server/errs"
)

type User struct {
	ID        string
	Name      string
	Birthdate string
	Password  string
	Email     string
	Document  string
	City      string
	Zipcode   string
	Phone     string
	Status    int
}

type UserRepository interface {
	FindAll(string) ([]User, *errs.AppError)
	FindByID(string) (*User, *errs.AppError)
}

const (
	ACTIVE   = "active"
	INACTIVE = "inactive"
)

func (u User) statusAsText() string {
	if u.Status == 0 {
		return INACTIVE
	}

	return ACTIVE
}

func (u User) ToDTO() dto.UserResponse {
	return dto.UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Birthdate: u.Birthdate,
		Password:  u.Password,
		Email:     u.Email,
		Document:  u.Document,
		City:      u.City,
		Zipcode:   u.Zipcode,
		Phone:     u.Phone,
		Status:    u.statusAsText(),
	}
}
