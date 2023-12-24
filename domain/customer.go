package domain

type Customer struct {
	Id          string `gorm:"column:customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateofBirth string `gorm:"column:date_of_birth"`
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
