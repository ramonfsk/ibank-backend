package service

import (
	"github.ibm.com/rfnascimento/ibank/server/domain"
	"github.ibm.com/rfnascimento/ibank/server/dto"
	"github.ibm.com/rfnascimento/ibank/server/errs"
)

type AccountService interface {
	GetAllAccounts(status string) ([]dto.AccountResponse, *errs.AppError)
	GetAccount(id string) (*dto.AccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) GetAllAccounts(status string) ([]dto.AccountResponse, *errs.AppError) {
	accounts, err := s.repo.FindAll(status)

	if len(accounts) == 0 {
		return nil, errs.NewValidationError("No have accounts for this bank on database")
	}

	response := make([]dto.AccountResponse, 0)
	for _, account := range accounts {
		response = append(response, account.ToDTO())
	}

	return response, err
}

func (s DefaultAccountService) GetAccount(id string) (*dto.AccountResponse, *errs.AppError) {
	account, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	response := account.ToDTO()

	return &response, nil
}

func NewAccountService(repository domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repository}
}
