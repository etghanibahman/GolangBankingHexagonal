package app

import (
	"RouterBasics/domain"
	"RouterBasics/logger"
	"RouterBasics/service"
	"RouterBasics/storage"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// set SERVER_ADDRESS=localhost
// set SERVER_PORT=8082
// export SERVER_ADDRESS=localhost
// SERVER_ADDRESS=localhost SERVER_PORT=8082 go run main.go
//	go run SERVER_ADDRESS=localhost SERVER_PORT=8082 main.go
// func sanityCheck() {
// 	logger.Info(os.Getenv("GOPATH"))
// 	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
// 		logger.Fatal("Enivironment variable is not defined...")
// 	}
// }

func Start() {

	//sanityCheck()

	router := mux.NewRouter()

	//wiring
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	dbClient := getDbClient()
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)
	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandler{service.NewAccountService(accountRepositoryDb)}

	//define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)
	//Starting server
	// address := os.Getenv("SERVER_ADDRESS")
	// port := os.Getenv("SERVER_PORT")
	// log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
	err := http.ListenAndServe("localhost:8000", router)

	logger.Fatal(err.Error())
}

func getDbClient() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Fatal(err.Error())
	}
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	client, err := storage.NewConnection(config)
	if err != nil {
		logger.Fatal("could not load the database")
	}
	return client
}
