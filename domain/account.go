package domain

import (
	"github.com/ashishjuyal/banking/dto"
	"github.com/ashishjuyal/banking/errs"
)

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{a.AccountId}
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
