package service

import (
	"github.com/ramonfsk/ibank-backend/auth/domain"
	"github.com/ramonfsk/ibank-backend/auth/dto"
	"github.com/ramonfsk/ibank-backend/auth/errs"
)

type DefaultAuthService struct {
	repository domain.AuthRepository
}

type AuthService interface {
	Login(dto.LoginRequest) (*dto.TokenResponse, *errs.AppError)
	Verify(urlParams map[string]string) (bool, *errs.AppError)
}

func (s DefaultAuthService) Login(request dto.LoginRequest) (*dto.TokenResponse, *errs.AppError) {
	login, err := s.repository.FindBy(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	token, err := login.GenerateToken()
	if err != nil {
		return nil, err
	}

	return &dto.TokenResponse{
		Token: *token,
	}, nil
}

// Sample URL String
// auth/verify?token=aaa.bbbb.cccc&routeName=getAllUsers
func (s DefaultAuthService) Verify(urlParams map[string]string) (bool, *errs.AppError) {
	return true, nil
}

func NewAuthService(repository domain.AuthRepository) DefaultAuthService {
	return DefaultAuthService{repository: repository}
}
