package domain

import (
	"github.ibm.com/rfnascimento/ibank/dto"
	"github.ibm.com/rfnascimento/ibank/errs"
)

type Account struct {
	ID          string  `db:"id"`
	UserID      string  `db:"user_id"`
	OpeningDate string  `db:"opening_date"`
	Agency      string  `db:"agency"`
	Number      string  `db:"number"`
	CheckDigit  string  `db:"check_digit"`
	PIN         string  `db:"pin"`
	Balance     float64 `db:"balance"`
	Status      int     `db:"status"`
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}

func (a Account) ToNewAccountResponseDTO() dto.NewAccountResponse {
	return dto.NewAccountResponse{ID: a.ID}
}

func (a Account) ToDTO() dto.AccountResponse {
	return dto.AccountResponse{}
}

func (a Account) CanWithdraw(value float64) bool {
	return a.Balance >= value
}
