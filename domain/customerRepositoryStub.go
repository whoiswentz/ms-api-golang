package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func NewCustomerRepositoryStub() *CustomerRepositoryStub {
	return &CustomerRepositoryStub{customers: []Customer{
		{"1", "Vinicios", "Curitiba", "666", "", ""},
		{"2", "Vagner", "Foz do Iguaçu", "777", "", ""},
	}}
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}
