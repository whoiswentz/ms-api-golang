package domain

type CustomerRepositoryStub struct {
	Customers []Customer
}

func NewCustomerRepositoryStub() *CustomerRepositoryStub {
	return &CustomerRepositoryStub{Customers: []Customer{
		{"1", "Vinicios", "Curitiba", "666", "", ""},
		{"2", "Vagner", "Foz do Iguaçu", "777", "", ""},
	}}
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.Customers, nil
}
