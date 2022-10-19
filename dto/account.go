package dto

import "github.ibm.com/rfnascimento/ibank/errs"

type NewAccountRequest struct {
	UserID      string  `json:"user_id"`
	OpeningDate string  `json:"opening_date"`
	Agency      string  `json:"agency"`
	Number      string  `json:"number"`
	CheckDigit  string  `json:"check_digit"`
	PIN         string  `json:"pin"`
	Balance     float64 `json:"balance"`
	Status      string  `json:"status"`
}

type NewAccountResponse struct {
	ID string `json:"id"`
}

type AccountResponse struct {
	ID          string  `json:"id,omitempty"`
	UserID      string  `json:"user_id"`
	OpeningDate string  `json:"opening_date"`
	Agency      string  `json:"agency"`
	Number      string  `json:"number"`
	CheckDigit  string  `json:"check_digit"`
	PIN         string  `json:"pin"`
	Balance     float64 `json:"balance"`
	Status      string  `json:"status"`
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Balance < 5000 {
		return errs.NewValidationError("To open a new account, you need to deposit atleast 5000.00")
	}

	return nil
}
