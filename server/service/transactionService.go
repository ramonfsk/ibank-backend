package service

import (
	"github.com/ramonfsk/ibank-backend/server/domain"
	"github.com/ramonfsk/ibank-backend/server/dto"
	"github.com/ramonfsk/ibank-backend/server/errs"
)

type TransactionService interface {
	GetAllTransactions() ([]dto.TransactionResponse, *errs.AppError)
	GetTransaction(id string) (*dto.TransactionResponse, *errs.AppError)
	NewTransaction(dto.NewTransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}

type DefaultTransactionService struct {
	repo domain.TransactionRepository
}

func (s DefaultTransactionService) GetAllTransactions() ([]dto.TransactionResponse, *errs.AppError) {
	transactions, err := s.repo.FindAll()

	if len(transactions) == 0 {
		return nil, errs.NewValidationError("No have transactions for this bank on database")
	}

	response := make([]dto.TransactionResponse, 0)
	for _, transaction := range transactions {
		response = append(response, transaction.ToDTO())
	}

	return response, err
}

func (s DefaultTransactionService) GetTransaction(id string) (*dto.TransactionResponse, *errs.AppError) {
	transaction, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	response := transaction.ToDTO()

	return &response, nil
}

func (s DefaultTransactionService) NewTransaction(req dto.NewTransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	// incoming request validation
	err := req.ValidateMakeTransaction()
	if err != nil {
		return nil, err
	}
	// get account with agency, number and check digit
	account, err := s.getAccountWithoutID(req.Agency, req.Number, req.CheckDigit)
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
	var bankId string
	if req.BankID == "" {
		bankId = "1"
	} else {
		bankId = req.BankID
	}

	transactionBuilted := domain.Transaction{
		BankID:     bankId,
		AccountID:  account.ID,
		Agency:     req.Agency,
		Number:     req.Number,
		CheckDigit: req.CheckDigit,
		Type:       req.Type,
		Value:      req.Value,
	}

	transaction, appErr := s.repo.RegisterNewTransaction(transactionBuilted)
	if appErr != nil {
		return nil, appErr
	}

	response := transaction.ToDTO()

	return &response, nil
}

func (s DefaultTransactionService) getAccountWithoutID(agency string, number string, checkDigit string) (*domain.Account, *errs.AppError) {
	account, err := s.repo.FindAccountWithoutID(agency, number, checkDigit)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func NewTransactionService(repository domain.TransactionRepository) DefaultTransactionService {
	return DefaultTransactionService{repo: repository}
}
