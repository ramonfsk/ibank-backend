package service

import (
	"github.com/ramonfsk/ibank-backend/server/domain"
	"github.com/ramonfsk/ibank-backend/server/dto"
	"github.com/ramonfsk/ibank-backend/server/errs"
)

type AccountService interface {
	GetAllAccounts(status string) ([]dto.AccountResponse, *errs.AppError)
	GetAccount(id string) (*dto.AccountResponse, *errs.AppError)
	GetAllTransactionByAccount(string) ([]dto.TransactionResponse, *errs.AppError)
	GetAllTransactionsByAccountWithPeriod(string, string, string) ([]dto.TransactionResponse, *errs.AppError)
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

func (s DefaultAccountService) GetAllTransactionByAccount(id string) ([]dto.TransactionResponse, *errs.AppError) {
	transactions, appErr := s.repo.FindAllTransactionsByID(id)
	if appErr != nil {
		return nil, appErr
	} else if len(transactions) == 0 {
		return nil, errs.NewValidationError("No have transactions for this account on database")
	}

	response := make([]dto.TransactionResponse, 0)
	for _, transaction := range transactions {
		response = append(response, transaction.ToDTO())
	}

	return response, nil
}

func (s DefaultAccountService) GetAllTransactionsByAccountWithPeriod(id string, startDate string, endDate string) ([]dto.TransactionResponse, *errs.AppError) {
	transactions, appErr := s.repo.FindAllTransactionsByAccountIDWithPeriod(id, startDate, endDate)
	if appErr != nil {
		return nil, appErr
	} else if len(transactions) == 0 {
		return nil, errs.NewValidationError("No have transactions for this account on database")
	}

	response := make([]dto.TransactionResponse, 0)
	for _, transaction := range transactions {
		response = append(response, transaction.ToDTO())
	}

	return response, nil
}

func NewAccountService(repository domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repository}
}
