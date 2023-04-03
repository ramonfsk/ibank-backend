package domain

import (
	"strings"
	"time"

	"github.com/ramonfsk/ibank/server/dto"
	"github.com/ramonfsk/ibank/server/errs"
	"github.com/ramonfsk/ibank/server/utils"
)

type Transaction struct {
	ID         string  `db:"id"`
	BankID     string  `db:"bank_id,omitempty"`
	AccountID  string  `json:"account_id,omitempty"`
	Agency     string  `db:"agency"`
	Number     string  `db:"account_number"`
	CheckDigit string  `db:"check_digit"`
	Type       string  `db:"type,omitempty"`
	CreatedAt  string  `db:"created_at"`
	Value      float64 `db:"value"`
}

type TransactionRepository interface {
	FindAll() ([]Transaction, *errs.AppError)
	FindByID(string) (*Transaction, *errs.AppError)
	RegisterNewTransaction(Transaction) (*Transaction, *errs.AppError)
	FindAccount(string) (*Account, *errs.AppError)
	FindAccountWithoutID(string, string, string) (*Account, *errs.AppError)
}

func (t Transaction) IsWithdrawal() bool {
	s := strings.ToLower(t.Type)
	return s == utils.WITHDRAW
}

func (t Transaction) ToDTO() dto.TransactionResponse {
	return dto.TransactionResponse{
		ID:         t.ID,
		BankID:     t.BankID,
		Agency:     t.Agency,
		Number:     t.Number,
		CheckDigit: t.CheckDigit,
		Type:       t.Type,
		CreatedAt:  time.Now().Format(time.RFC3339),
		Value:      t.Value,
	}
}
