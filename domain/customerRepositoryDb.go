package domain

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
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

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {
	findAllSQL := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers"

	rows, err := d.client.Query(findAllSQL)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		if err := rows.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status); err != nil {
			log.Fatalf("Error while scanning customer table: %s", err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}

	return customers, nil
}