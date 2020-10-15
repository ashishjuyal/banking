package app

import (
	"github.com/ashishjuyal/banking/domain"
	"github.com/ashishjuyal/banking/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {

	router := mux.NewRouter()

	//wiring
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))

}
