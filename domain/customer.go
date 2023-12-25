package domain

import "RouterBasics/errs"

type Customer struct {
	Id          string `gorm:"column:customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateofBirth string `gorm:"column:date_of_birth"`
	Status      string
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
