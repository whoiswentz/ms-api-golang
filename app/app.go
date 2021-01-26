package app

import (
	"banking/domain"
	"banking/logger"
	"banking/service"
	"github.com/gorilla/mux"
	"net/http"
)



func Start() {
	db, err := Connect()
	if err != nil {
		logger.Panic(err.Error())
	}

	customerHandler := CustomerHandler{service: service.NewCustomerService(domain.NewCustomerRepositoryDb(db))}
	accountHandler := AccountHandler{service: service.NewAccountService(domain.NewAccountRepositoryDb(db))}
	authHandler := AuthHandler{service: service.NewDefaultAuthService(domain.NewAuthRepositoryDb(db))}

	router := mux.NewRouter()
	//router.Use(NewAuthMiddleware(service.NewDefaultAuthService(domain.NewAuthRepositoryDb(db))).AuthorizationHandler())

	router.HandleFunc("/customers", customerHandler.getAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id}", customerHandler.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id}/account", accountHandler.newAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id}/account/{account_id}", accountHandler.MakeTransaction).Methods(http.MethodPost)
	router.HandleFunc("/auth/login", authHandler.Login).Methods(http.MethodPost)

	logger.Info("starting the application")
	if err := http.ListenAndServe("localhost:8080", router); err != nil {
		logger.Panic(err.Error())
	}
}
