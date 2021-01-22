package domain

import (
	"banking/errs"
	"banking/logger"
	"database/sql"
	_ "github.com/lib/pq"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func NewCustomerRepositoryDb() *CustomerRepositoryDb {
	db, err := sql.Open("postgres", "")
	if err != nil {
		panic(err)
	}
	return &CustomerRepositoryDb{client: db}
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {

	var rows *sql.Rows
	var err error
	if status == "" {
		findAllSQL := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers"
		rows, err = d.client.Query(findAllSQL)
	} else {
		findAllSQL := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE status = $1"
		rows, err = d.client.Query(findAllSQL, status)
	}

	if err != nil {
		logger.Error("error while querying customers: " + err.Error())
		return nil, errs.UnexpectedError("Unexpected Error")
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		if err := rows.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status); err != nil {
			logger.Error("error while scanning customers: " + err.Error())
			return nil, errs.UnexpectedError("Unexpected Error")
		}
		customers = append(customers, c)
	}

	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	findByIdSQL := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE customer_id = $1"

	row := d.client.QueryRow(findByIdSQL, id)

	var c Customer
	if err := row.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status); err != nil {
		if err == sql.ErrNoRows {
			logger.Error("customer not found: " + err.Error())
			return nil, errs.NotFoundError("customer not found")
		}
		logger.Error("error while scanning customers: " + err.Error())
		return nil, errs.UnexpectedError("unexpected error")
	}

	return &c, nil
}
