package domain

import (
	"time"

	"github.com/ramonfsk/ibank/server/dto"
	"github.com/ramonfsk/ibank/server/errs"
)

type TransactionRepositoryStub struct {
	transactions []Transaction
}

func (r TransactionRepositoryStub) GetByPeriod(days int) ([]Transaction, *errs.AppError) {
	return r.transactions, nil
}

func (r TransactionRepositoryStub) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	return &r.transactions[0], nil
}

func (d TransactionRepositoryStub) FindAccountByID(id string) (*dto.AccountResponse, *errs.AppError) {
	return &dto.AccountResponse{}, nil
}

func NewTransactionRepositoryStub() TransactionRepositoryStub {
	transaction := &Transaction{
		ID:         "1",
		BankID:     "1",
		AccountID:  "1",
		Agency:     "0001",
		Number:     "10001234",
		CheckDigit: "7",
		Type:       "deposit",
		CreatedAt:  time.Now().Format(time.RFC3339),
		Value:      100,
	}

	return TransactionRepositoryStub{
		transactions: []Transaction{*transaction},
	}
}
