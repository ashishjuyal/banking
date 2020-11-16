package dto

import (
	"net/http"
	"testing"
)

func Test_should_return_error_when_transaction_type_is_not_deposit_or_withdrawl(t *testing.T) {
	// Arrange
	request := TransactionRequest{
		TransactionType: "invalid transaction type",
	}
	// Act
	appError := request.Validate()

	// Assert
	if appError.Message != "Transaction type can only be deposit or withdrawal" {
		t.Error(" Invalid message while testing transaction type")
	}

	if appError.Code != http.StatusUnprocessableEntity {
		t.Error(" Invalid code while testing transaction type")
	}
}

func Test_should_return_error_when_amount_is_less_than_zero(t *testing.T) {
	// Arrange
	request := TransactionRequest{TransactionType: DEPOSIT, Amount: -100}

	// Act
	appError := request.Validate()

	// Assert
	if appError.Message != "Amount cannot be less than zero" {
		t.Error("Invalid message while validating amount")
	}
	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while validating amount")
	}

}
