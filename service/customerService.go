package service

import "banking/domain"

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repository domain.CustomerRepository
}

func NewDefaultCustomerService(repository domain.CustomerRepository) *DefaultCustomerService {
	return &DefaultCustomerService{repository: repository}
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repository.FindAll()
}