package app

import (
	"banking/domain"
	"banking/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()

	customerHandler := CustomerHandler{
		service: service.NewDefaultCustomerService(domain.NewCustomerRepositoryDb()),
	}
	router.HandleFunc("/customers", customerHandler.getAllCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
