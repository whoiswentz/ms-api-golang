package app

import (
	"banking/logger"
	"database/sql"
	"os"
)

var dbURL = os.Getenv("DB_URL")
var dbEngine = os.Getenv("DB_ENGINE")

func init() {
	if dbURL == "" {
		logger.Panic("db url must be provided")
	}
	if dbEngine == "" {
		logger.Panic("db engine must be provided")
	}
}

func Connect() (*sql.DB, error) {
	logger.Info("connecting on database")
	db, err := sql.Open(dbEngine, dbURL)
	if err != nil {
		logger.Error("error while connection on db")
		return nil, err
	}
	logger.Info("connected on database")
	return db, nil
}