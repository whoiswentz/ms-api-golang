package app

import (
	"banking/dto"
	"banking/errs"
	"banking/mocks/service"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

var customerHandler CustomerHandler
var router *mux.Router
var mockService *service.MockCustomerService

func setup(t *testing.T) func() {
	controller := gomock.NewController(t)
	mockService = service.NewMockCustomerService(controller)

	customerHandler = CustomerHandler{service: mockService}
	router = mux.NewRouter()

	return func() {
		router = nil
		defer controller.Finish()
	}
}

func TestShouldReturnCustomerWithStatusCode200(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	dummyCustomers := []dto.CustomerResponse{
		{"1", "Vinicios", "Curitiba", "666", "", ""},
		{"2", "Vagner", "Foz do Iguaçu", "777", "", ""},
	}

	mockService.EXPECT().GetAllCustomer("").Return(dummyCustomers, nil)

	router.HandleFunc("/customer", customerHandler.getAllCustomer)
	request, err := http.NewRequest(http.MethodGet, "/customer", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func TestShouldReturnStatusCode500WIthErrorMessage(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	mockService.EXPECT().GetAllCustomer("").Return(nil, errs.UnexpectedError("unexpected error"))

	router.HandleFunc("/customer", customerHandler.getAllCustomer)
	request, err := http.NewRequest(http.MethodGet, "/customer", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
}