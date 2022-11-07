package domain

import (
	"github.ibm.com/rfnascimento/ibank/server/dto"
	"github.ibm.com/rfnascimento/ibank/server/errs"
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

func (a Account) Validate() *errs.AppError {
	if a.Balance < 5000 {
		return errs.NewValidationError("To open a new account, you need to deposit atleast 5000.00")
	}

	return nil
}

func (a Account) CanWithdraw(value float64) bool {
	return a.Balance >= value
}

func (a Account) ToDTO() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		Agency:        a.Agency,
		NumberAccount: a.Number,
		CheckDigit:    a.CheckDigit,
	}
}
