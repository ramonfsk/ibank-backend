package service

import "github.ibm.com/rfnascimento/ibank/auth/dto"

type DefaultAuthService struct {
	repository      domain.AuthRepository
	rolePermissions domain.RolePermissions
}

type AuthService interface {
	Login(dto.LoginRequest) (*string, error)
	Verify(urlParams map[string]string) (bool, error)
}

func (s DefaultAuthService) Login(request dto.LoginRequest) (*string, error) {

}
