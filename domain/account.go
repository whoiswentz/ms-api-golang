package domain

import (
	"banking/dto"
	"banking/errs"
)

type Account struct {
	AccountId string
	CustomerId string
	OpeningDate string
	AccountType string
	Amount float64
	Status string
}

func (a Account) ToDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountId: a.AccountId}
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	FindBy(string) (*Account, *errs.AppError)
	SaveTransaction(Transaction) (*Transaction, *errs.AppError)
}

func (a Account) CanWithdraw(amount float64) bool {
	if a.Amount < amount {
		return false
	}
	return true
}