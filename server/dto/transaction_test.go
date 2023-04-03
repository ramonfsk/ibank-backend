package dto

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/ramonfsk/ibank/server/utils"
)

func TestSholdReturnErrorWhenTransactionTypeIsNotDepositOrWithdrawl(t *testing.T) {
	// Arrange
	request := NewTransactionRequest{
		Type: "invalid transaction type",
	}
	// Act
	appError := request.ValidateMakeTransaction()
	// Assert
	if appError.Message != "Type"+request.Type+" of transaction is invalid!" {
		t.Error("Invalid message while testing transaction type")
	}

	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while testing transaction type")
	}
}

func TestShouldReturnErrorWhenAmountIsLessThanZero(t *testing.T) {
	// Arrange
	request := NewTransactionRequest{
		BankID:     "1",
		Agency:     "100",
		Number:     "1000",
		CheckDigit: "1",
		Type:       utils.DEPOSIT,
		Value:      -100.0,
	}
	// Act
	appError := request.ValidateMakeTransaction()
	// Assert
	if appError.Message != fmt.Sprintf("Value R$ %.2f of transaction is invalid!", request.Value) {
		t.Error("Invalid message while testing transaction value")
	}

	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while testing transaction value")
	}
}
