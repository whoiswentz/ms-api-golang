package app

import (
	"banking/domain"
	"banking/logger"
	"banking/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	logger.Info("Starting the application")

	router := mux.NewRouter()

	customerHandler := CustomerHandler{
		service: service.NewDefaultCustomerService(domain.NewCustomerRepositoryDb()),
	}

	router.HandleFunc("/customers", customerHandler.getAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id}", customerHandler.getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
