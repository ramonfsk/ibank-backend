package dto

import (
	"fmt"
	"regexp"
	"strings"

	"github.ibm.com/rfnascimento/ibank/errs"
)

const (
	WITHDRAW = "withdraw"
	DEPOSIT  = "deposit"
	CHECK    = "check"
	TRANSFER = "transfer"
)

type TransactionRequest struct {
	BankID     string  `json:"bank_id,omitempty"`
	Agency     string  `json:"agency"`
	Number     string  `json:"number"`
	CheckDigit string  `json:"check_digit"`
	Type       string  `json:"type"`
	Value      float64 `json:"value"`
}

type TransactionResponse struct {
	ID         string  `json:"id,omitempty"`
	BankID     string  `json:"bank_id,omitempty"`
	Agency     string  `json:"agency"`
	Number     string  `json:"number"`
	CheckDigit string  `json:"check_digit"`
	Type       string  `json:"type"`
	CreatedAt  string  `json:"created_at"`
	Value      float64 `json:"value"`
}

func (t TransactionRequest) IsTransactionTypeWithdrawal() bool {
	s := strings.ToLower(t.Type)
	return s == WITHDRAW
}

func (t TransactionRequest) ValidateMakeTransaction() *errs.AppError {
	typeErr := t.IsValidType()
	accountAgencyErr := t.IsValidAccountAgency()
	accountNumberErr := t.IsValidAccountNumber()
	accountCheckDigitErr := t.IsValidAccountCheckDigit()
	valueErr := t.IsValidValue()
	if typeErr != nil {
		return typeErr
	} else if t.IsValidAccountAgency() != nil {
		return accountAgencyErr
	} else if t.IsValidAccountNumber() != nil {
		return accountNumberErr
	} else if t.IsValidAccountCheckDigit() != nil {
		return accountCheckDigitErr
	} else if t.IsValidValue() != nil {
		return valueErr
	}

	return nil
}

func (t TransactionRequest) ValidateGetAllTransactions() *errs.AppError {
	accountAgencyErr := t.IsValidAccountAgency()
	accountNumberErr := t.IsValidAccountNumber()
	accountCheckDigitErr := t.IsValidAccountCheckDigit()
	if t.IsValidAccountAgency() != nil {
		return accountAgencyErr
	} else if t.IsValidAccountNumber() != nil {
		return accountNumberErr
	} else if t.IsValidAccountCheckDigit() != nil {
		return accountCheckDigitErr
	}

	return nil
}

func (t TransactionRequest) IsValidType() *errs.AppError {
	s := strings.ToLower(t.Type)
	switch s {
	case WITHDRAW:
		return nil
	case DEPOSIT:
		return nil
	case CHECK:
		return nil
	case TRANSFER:
		return nil
	default:
		return errs.NewValidationError("Type" + t.Type + " of transaction is invalid!")
	}
}

func (t TransactionRequest) IsValidAccountAgency() *errs.AppError {
	regex, _ := regexp.Compile(`[0-9]+`)

	if !regex.MatchString(t.Agency) {
		return errs.NewValidationError("Account agency" + t.Agency + " is invalid!")
	}

	return nil
}

func (t TransactionRequest) IsValidAccountNumber() *errs.AppError {
	regex, _ := regexp.Compile(`[0-9]+`)

	if !regex.MatchString(t.Number) {
		return errs.NewValidationError("Account number" + t.Number + " is invalid!")
	}

	return nil
}

func (t TransactionRequest) IsValidAccountCheckDigit() *errs.AppError {
	regex, _ := regexp.Compile(`[0-9]+`)

	if !regex.MatchString(t.CheckDigit) {
		return errs.NewValidationError("Account check digit" + t.CheckDigit + " is invalid!")
	}

	return nil
}

func (t TransactionRequest) IsValidValue() *errs.AppError {
	if t.Value < 0.01 {
		return errs.NewValidationError(fmt.Sprintf("Value R$ %.2f of transaction is invalid!", t.Value))
	}

	return nil
}
