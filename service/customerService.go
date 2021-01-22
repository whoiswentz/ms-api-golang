package service

import (
	"banking/domain"
	"banking/dto"
	"banking/errs"
)

type CustomerService interface {
	GetAllCustomer(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repository domain.CustomerRepository
}

func NewDefaultCustomerService(repository domain.CustomerRepository) *DefaultCustomerService {
	return &DefaultCustomerService{repository: repository}
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}

	customers, err := s.repository.FindAll("")
	if err != nil {
		return nil, err
	}

	response := make([]dto.CustomerResponse, 0)
	for _, customer := range customers {
		response = append(response, customer.ToDto())
	}
	return response, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	customer, err := s.repository.ById(id)
	if err != nil {
		return nil, err
	}

	response := customer.ToDto()
	return &response, nil
}
