package domain

import (
	"strings"

	"github.ibm.com/rfnascimento/ibank/dto"
	"github.ibm.com/rfnascimento/ibank/errs"
)

type Transaction struct {
	ID         string  `db:"id"`
	BankID     string  `db:"bank_id,omitempty"`
	AccountID  string  `json:"account_id,omitempty"`
	Agency     string  `db:"agency"`
	Number     string  `db:"number"`
	CheckDigit string  `db:"check_digit"`
	Type       string  `db:"type,omitempty"`
	CreatedAt  string  `db:"created_at"`
	Value      float64 `db:"value"`
}

type TransactionRepository interface {
	// GetByPeriod(string) ([]Transaction, *errs.AppError)
	GetAllTransactionsByAccount(transaction Transaction) ([]Transaction, *errs.AppError)
	SaveTransaction(transaction Transaction) (*Transaction, *errs.AppError)
	FindAccount(string, string, string) (*Account, *errs.AppError)
}

const (
	WITHDRAW = "withdraw"
	DEPOSIT  = "deposit"
)

func (t Transaction) IsWithdrawal() bool {
	s := strings.ToLower(t.Type)
	return s == WITHDRAW
}

func (t Transaction) ToDTO() dto.TransactionResponse {
	return dto.TransactionResponse{
		ID:         t.ID,
		BankID:     t.BankID,
		Agency:     t.Agency,
		Number:     t.Number,
		CheckDigit: t.CheckDigit,
		Type:       t.Type,
		CreatedAt:  t.CreatedAt,
		Value:      t.Value,
	}
}
