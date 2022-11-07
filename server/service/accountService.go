package service

import (
	"github.ibm.com/rfnascimento/ibank/server/domain"
)

type AccountService interface {
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func NewAccountService(repository domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repository}
}
