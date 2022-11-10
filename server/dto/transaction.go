package dto

import (
	"fmt"
	"regexp"
	"strings"

	"github.ibm.com/rfnascimento/ibank/server/errs"
	"github.ibm.com/rfnascimento/ibank/server/utils"
)

type NewTransactionRequest struct {
	BankID     string  `json:"bank_id,omitempty"`
	Agency     string  `json:"agency"`
	Number     string  `json:"account_number"`
	CheckDigit string  `json:"check_digit"`
	Type       string  `json:"type"`
	Value      float64 `json:"value"`
}

type TransactionResponse struct {
	ID         string  `json:"id,omitempty"`
	BankID     string  `json:"bank_id,omitempty"`
	Agency     string  `json:"agency"`
	Number     string  `json:"account_number"`
	CheckDigit string  `json:"check_digit"`
	Type       string  `json:"type"`
	CreatedAt  string  `json:"created_at"`
	Value      float64 `json:"value"`
}

func (t NewTransactionRequest) IsTransactionTypeWithdrawal() bool {
	s := strings.ToLower(t.Type)
	return s == utils.WITHDRAW
}

func (t NewTransactionRequest) ValidateMakeTransaction() *errs.AppError {
	typeErr := t.IsValidType()
	accountAgencyErr := t.IsValaccountIdAgency()
	accountNumberErr := t.IsValaccountIdNumber()
	accountCheckDigitErr := t.IsValaccountIdCheckDigit()
	valueErr := t.IsValidValue()
	if typeErr != nil {
		return typeErr
	} else if t.IsValaccountIdAgency() != nil {
		return accountAgencyErr
	} else if t.IsValaccountIdNumber() != nil {
		return accountNumberErr
	} else if t.IsValaccountIdCheckDigit() != nil {
		return accountCheckDigitErr
	} else if t.IsValidValue() != nil {
		return valueErr
	}

	return nil
}

func (t NewTransactionRequest) ValidateGetAllTransactions() *errs.AppError {
	accountAgencyErr := t.IsValaccountIdAgency()
	accountNumberErr := t.IsValaccountIdNumber()
	accountCheckDigitErr := t.IsValaccountIdCheckDigit()
	if t.IsValaccountIdAgency() != nil {
		return accountAgencyErr
	} else if t.IsValaccountIdNumber() != nil {
		return accountNumberErr
	} else if t.IsValaccountIdCheckDigit() != nil {
		return accountCheckDigitErr
	}

	return nil
}

func (t NewTransactionRequest) IsValidType() *errs.AppError {
	s := strings.ToLower(t.Type)
	switch s {
	case utils.WITHDRAW:
		return nil
	case utils.DEPOSIT:
		return nil
	case utils.CHECK:
		return nil
	case utils.TRANSFER:
		return nil
	default:
		return errs.NewValidationError("Type" + t.Type + " of transaction is invalid!")
	}
}

func (t NewTransactionRequest) IsValaccountIdAgency() *errs.AppError {
	regex, _ := regexp.Compile(`[0-9]+`)

	if !regex.MatchString(t.Agency) {
		return errs.NewValidationError("Account agency" + t.Agency + " is invalid!")
	}

	return nil
}

func (t NewTransactionRequest) IsValaccountIdNumber() *errs.AppError {
	regex, _ := regexp.Compile(`[0-9]+`)

	if !regex.MatchString(t.Number) {
		return errs.NewValidationError("Account number" + t.Number + " is invalid!")
	}

	return nil
}

func (t NewTransactionRequest) IsValaccountIdCheckDigit() *errs.AppError {
	regex, _ := regexp.Compile(`[0-9]+`)

	if !regex.MatchString(t.CheckDigit) {
		return errs.NewValidationError("Account check digit" + t.CheckDigit + " is invalid!")
	}

	return nil
}

func (t NewTransactionRequest) IsValidValue() *errs.AppError {
	if t.Value < 0.01 {
		return errs.NewValidationError(fmt.Sprintf("Value R$ %.2f of transaction is invalid!", t.Value))
	}

	return nil
}
