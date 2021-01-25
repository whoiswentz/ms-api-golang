package domain

import (
	"banking/errs"
	"database/sql"
)

type AuthRepository interface {
	FindBy(string, string) (*Login, *errs.AppError)
}

type AuthRepositoryDb struct {
	client *sql.DB
}

func NewAuthRepositoryDb(client *sql.DB) *AuthRepositoryDb {
	return &AuthRepositoryDb{client: client}
}

func (r AuthRepositoryDb) FindBy(username string, password string) (*Login, *errs.AppError) {
	sqlVerify := `
		SELECT username, u.customer_id, role, array_agg(a.account_id) as account_numbers
		FROM users u
			LEFT JOIN accounts a ON a.customer_id = u.customer_id
		WHERE username = $1
		  AND password = $2
		GROUP BY a.customer_id, username;`

	var l Login
	if err := r.client.QueryRow(sqlVerify, username, password).Scan(&l.Username, &l.CustomerId, &l.Role,&l.Accounts); err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NotFoundError("user not found")
		}
		return nil, errs.UnexpectedError("unexpected error")
	}

	return &l, nil
}

