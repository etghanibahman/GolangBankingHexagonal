package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Ali", City: "Stuttgart", Zipcode: "110012", DateofBirth: "2000-01-01", Status: "1"},
		{Id: "10012", Name: "Bahman", City: "Kuala lumpur", Zipcode: "898988", DateofBirth: "2000-02-02", Status: "1"},
	}
	return CustomerRepositoryStub{customers}
}
