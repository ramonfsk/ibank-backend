package service

import (
	"github.ibm.com/rfnascimento/ibank/auth/dto"
	"github.ibm.com/rfnascimento/ibank/auth/errs"
)

type DefaultAuthService struct {
	repository      domain.AuthRepository
	rolePermissions domain.RolePermissions
}

type AuthService interface {
	Login(dto.LoginRequest) (*string, *errs.AppError)
	Verify(urlParams map[string]string) (bool, error)
}

func (s DefaultAuthService) Login(request dto.LoginRequest) (*string, *errs.AppError) {
	login, err := s.repository.FindBy(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	token, err := login.GenerateToken()
	if err != nil {
		return nil, err
	}

	return token
}
