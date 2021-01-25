package app

import (
	"banking/domain"
	"banking/logger"
	"banking/service"
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"
)

func Start() {
	logger.Info("connecting on database")
	db, err := sql.Open("postgres", "")
	if err != nil {
		logger.Panic(err.Error())
	}
	logger.Info("connected on database")

	customerHandler := CustomerHandler{
		service: service.NewDefaultCustomerService(domain.NewCustomerRepositoryDb(db)),
	}

	router := mux.NewRouter()
	router.HandleFunc("/customers", customerHandler.getAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id}", customerHandler.getCustomer).Methods(http.MethodGet)

	logger.Info("starting the application")
	if err := http.ListenAndServe("localhost:8080", router); err != nil {
		logger.Panic(err.Error())
	}
}
