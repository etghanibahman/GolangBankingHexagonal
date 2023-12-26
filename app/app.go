package app

import (
	"RouterBasics/domain"
	"RouterBasics/logger"
	"RouterBasics/service"
	"net/http"

	"github.com/gorilla/mux"
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
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	//define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	//Starting server
	// address := os.Getenv("SERVER_ADDRESS")
	// port := os.Getenv("SERVER_PORT")
	// log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
	err := http.ListenAndServe("localhost:8000", router)

	logger.Fatal(err.Error())
}
