package domain

import (
	"banking/errs"
	"banking/logger"
	"context"
	"database/sql"
)

type AccountRepositoryDb struct {
	client *sql.DB
}

func NewAccountRepositoryDb(client *sql.DB) *AccountRepositoryDb {
	return &AccountRepositoryDb{client: client}
}

func (r AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := `INSERT INTO 
    	accounts (customer_id, opening_date, account_type, amount, status, pin) 
    	VALUES ($1,$2,$3,$4,$5, '11111') RETURNING account_id`

	if err := r.client.QueryRow(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status).Scan(&a.AccountId); err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.UnexpectedError("unexpected error")
	}

	return &a, nil
}

func (r AccountRepositoryDb) FindBy(customerId string) (*Account, *errs.AppError) {
	sqlSelect := `SELECT account_id, customer_id, opening_date, account_type, amount FROM accounts WHERE account_id = $1`
	row := r.client.QueryRow(sqlSelect, customerId)

	var a Account
	if err := row.Scan(&a.AccountId, &a.CustomerId, &a.OpeningDate, &a.AccountType, &a.Amount); err != nil {
		logger.Error("error while scanning the row: " + err.Error())
		return nil, errs.UnexpectedError("unexpected error")
	}

	return &a, nil
}

func (r AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	tx, err := r.client.BeginTx(context.Background(), nil)
	if err != nil {
		return nil, errs.UnexpectedError("")
	}

	result := tx.QueryRow(`INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) VALUES ($1, $2, $3, $4) RETURNING transaction_id`,
		t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)
	if err := result.Scan(&t.TransactionId); err != nil {
		logger.Error("error while scanning the row: " + err.Error())
		return nil, errs.UnexpectedError("unexpected error")
	}

	if t.IsWithdrawal() {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount - $1 WHERE account_id = $2 RETURNING account_id`, t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount + $1 WHERE account_id = $2 RETURNING account_id`, t.Amount, t.AccountId)
	}
	if err != nil {
		tx.Rollback()
		logger.Error("error while saving a transaction: " + err.Error())
		return nil, errs.UnexpectedError("unexpected error")
	}

	if err := tx.Commit(); err != nil {
		logger.Error("error while commiting transaction: " + err.Error())
		return nil, errs.UnexpectedError("unexpected error")
	}

	account, appError := r.FindBy(t.AccountId)
	if appError != nil {
		return nil, appError
	}

	t.Amount = account.Amount
	return &t, nil
}