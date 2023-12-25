package domain

import (
	"RouterBasics/errs"
	"RouterBasics/storage"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type CustomerRepositoryDb struct {
	client *gorm.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {

	customers := make([]Customer, 0)
	var err *gorm.DB
	//d.client.Statement.Exec("select customer_id, name, city, zipcode, date_of_birth, status from customers").Find(&customers)
	if status == "" {
		err = d.client.Select("customer_id", "name", "city", "zipcode", "date_of_birth", "status").Find(&customers)
	} else {
		err = d.client.Select("customer_id", "name", "city", "zipcode", "date_of_birth", "status").Where("status = ?", status).Find(&customers)
	}

	if err.Error != nil {
		if err.Error.Error() == "record not found" {
			return nil, errs.NewNotFoundError("There is not any customer in the table!")
		} else {
			log.Println("Error while fetching customers" + err.Error.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}

	}
	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	var customer Customer

	err := d.client.Where("customer_id = ?", id).First(&customer)

	if err.Error != nil {
		if err.Error.Error() == "record not found" {
			log.Println("Customer not found error happened")
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			log.Println("Error while scanning customer" + err.Error.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}

	}
	return &customer, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("could not load the database")
	}

	return CustomerRepositoryDb{db}
}
