package service

import (
	"banking/domain"
	"banking/dto"
	"banking/errs"
	"time"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	MakeTransaction(r dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repository domain.AccountRepository
}

func NewAccountService(r domain.AccountRepository) *DefaultAccountService {
	return &DefaultAccountService{repository: r}
}

func (s DefaultAccountService) NewAccount(r dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	account := domain.Account{
		AccountId:   "",
		CustomerId:  r.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: r.AccountType,
		Amount:      r.Amount,
		Status:      "1",
	}

	newAccount, err := s.repository.Save(account)
	if err != nil {
		return nil, err
	}

	response := newAccount.ToDto()
	return &response, nil
}

func (s DefaultAccountService) MakeTransaction(r dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	if r.IsTransactionTypeWithdrawal() {
		account, err := s.repository.FindBy(r.AccountId)
		if err != nil {
			return nil, err
		}
		if !account.CanWithdraw(r.Amount) {
			return nil, errs.NewValidationError("insufficient balance in the account")
		}
	}

	t := domain.Transaction{
		AccountId:       r.AccountId,
		Amount:          r.Amount,
		TransactionType: r.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	transaction, err := s.repository.SaveTransaction(t)
	if err != nil {
		return nil, err
	}

	response := transaction.ToDto()
	return &response, nil
}