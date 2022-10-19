package service

import (
	"github.ibm.com/rfnascimento/ibank/domain"
	"github.ibm.com/rfnascimento/ibank/dto"
	"github.ibm.com/rfnascimento/ibank/errs"
)

type TransactionService interface {
	// GetByPeriod(days string) ([]domain.Transaction, *errs.AppError)
	GetAllTransactionsByAccount(dto.TransactionRequest) ([]dto.TransactionResponse, *errs.AppError)
	MakeTransaction(dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
	GetAccount(agency string, number string, checkDigit string) (*domain.Account, *errs.AppError)
}

type DefaultTransactionService struct {
	repo domain.TransactionRepository
}

// func (ts DefaultTransactionService) GetByPeriod(days string) ([]domain.Transaction, *errs.AppError) {
// 	return ts.repo.GetByPeriod(days)
// }

func (ts DefaultTransactionService) GetAllTransactionsByAccount(req dto.TransactionRequest) ([]dto.TransactionResponse, *errs.AppError) {
	// incoming request validation
	err := req.ValidateGetAllTransactions()
	if err != nil {
		return nil, err
	}
	// if all is well, build the domain object & save the transaction
	transactionBuilted := domain.Transaction{
		Agency:     req.Agency,
		Number:     req.Number,
		CheckDigit: req.CheckDigit,
	}

	transactions, appErr := ts.repo.GetAllTransactionsByAccount(transactionBuilted)
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

func (ts DefaultTransactionService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	// incoming request validation
	err := req.ValidateMakeTransaction()
	if err != nil {
		return nil, err
	}
	// get account with agency, number and check digit
	account, err := ts.GetAccount(req.Agency, req.Number, req.CheckDigit)
	// server side validation for checking the available balance in the account
	if req.IsTransactionTypeWithdrawal() {
		if err != nil {
			return nil, err
		}

		if !account.CanWithdraw(req.Value) {
			return nil, errs.NewValidationError("Insufficient balance in account")
		}
	}
	// if all is well, build the domain object & save the transaction
	transactionBuilted := domain.Transaction{
		BankID:     req.BankID,
		AccountID:  account.ID,
		Agency:     req.Agency,
		Number:     req.Number,
		CheckDigit: req.CheckDigit,
		Type:       req.Type,
		Value:      req.Value,
	}

	transaction, appErr := ts.repo.SaveTransaction(transactionBuilted)
	if appErr != nil {
		return nil, appErr
	}

	response := transaction.ToDTO()

	return &response, nil
}

func (ts DefaultTransactionService) GetAccount(agency string, number string, checkDigit string) (*domain.Account, *errs.AppError) {
	account, err := ts.repo.FindAccount(agency, number, checkDigit)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func NewTransactionService(repo domain.TransactionRepository) DefaultTransactionService {
	return DefaultTransactionService{
		repo: repo,
	}
}
