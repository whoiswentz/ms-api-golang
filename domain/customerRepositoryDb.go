package domain

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type CustomerRepositoryDb struct {

}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {
	db, err := sql.Open("postgres", "")
	if err != nil {
		log.Fatal("error connection to the database: ", err)
	}
	defer db.Close()
}