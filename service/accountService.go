package service

import (
	"time"

	"github.ibm.com/rfnascimento/ibank/domain"
	"github.ibm.com/rfnascimento/ibank/dto"
	"github.ibm.com/rfnascimento/ibank/errs"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (as DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	account := domain.Account{
		ID:          "",
		UserID:      req.UserID,
		OpeningDate: time.Now().Format(time.RFC3339),
		Agency:      req.Agency,
		Number:      req.Number,
		CheckDigit:  req.CheckDigit,
		PIN:         req.PIN,
		Balance:     req.Balance,
		Status:      1,
	}

	newAccount, err := as.repo.Save(account)
	if err != nil {
		return nil, err
	}

	response := newAccount.ToNewAccountResponseDTO()

	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{
		repo: repo,
	}
}
