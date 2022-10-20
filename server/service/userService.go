package service

import (
	"github.ibm.com/rfnascimento/ibank/server/domain"
	"github.ibm.com/rfnascimento/ibank/server/dto"
	"github.ibm.com/rfnascimento/ibank/server/errs"
)

type UserService interface {
	GetAllUsers(status string) ([]domain.User, *errs.AppError)
	GetUser(id string) (*dto.UserResponse, *errs.AppError)
}

type DefaultUserService struct {
	repo domain.UserRepository
}

func (s DefaultUserService) GetAllUsers(status string) ([]domain.User, *errs.AppError) {
	users, err := s.repo.FindAll(status)

	if len(users) == 0 {
		return nil, errs.NewValidationError("No have users for this bank on database")
	}

	return users, err
}

func (s DefaultUserService) GetUser(id string) (*dto.UserResponse, *errs.AppError) {
	u, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	response := u.ToDTO()

	return &response, nil
}

func NewUserService(repository domain.UserRepository) DefaultUserService {
	return DefaultUserService{repo: repository}
}
