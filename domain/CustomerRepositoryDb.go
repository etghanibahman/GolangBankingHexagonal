package domain

import (
	"RouterBasics/storage"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type CustomerRepositoryDb struct {
	client *gorm.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {

	customers := make([]Customer, 0)
	d.client.Statement.Exec("select customer_id, name, city, zipcode, date_of_birth, status from customers").Find(&customers)
	//d.client.Select("customer_id", "name", "city", "zipcode", "date_of_birth", "status").Find(&customers)

	log.Writer().Write([]byte(customers[len(customers)-1].Name))
	log.Writer().Write([]byte(customers[len(customers)-1].DateofBirth))
	return customers, nil
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

// func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
// 	var err error
// 	customers := make([]Customer, 0)

// 	if status == "" {
// 		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
// 		err = d.client.Select(&customers, findAllSql)
// 	} else {
// 		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
// 		err = d.client.Select(&customers, findAllSql, status)
// 	}

// 	if err != nil {
// 		logger.Error("Error while querying customers table " + err.Error())
// 		return nil, errs.NewUnexpectedError("Unexpected database error")
// 	}

// 	return customers, nil
// }
