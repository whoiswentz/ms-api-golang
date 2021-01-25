package domain

import (
	"banking/errs"
	"banking/logger"
	"database/sql"
	"strconv"
)

type AccountRepositoryDb struct {
	client *sql.DB
}

func NewAccountRepositoryDb(client *sql.DB) *AccountRepositoryDb {
	return &AccountRepositoryDb{client: client}
}

func (r AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO account (customer_id, opening_date, account_type, amount, status) VALUES (?,?,?,?,?)"

	result, err := r.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.UnexpectedError("unexpected error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last id: " + err.Error())
		return nil, errs.UnexpectedError("unexpected error")
	}

	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

