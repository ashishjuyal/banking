package app

import (
	"encoding/json"
	"github.com/ashishjuyal/banking/service"
	"github.com/gorilla/mux"
	"net/http"
)

type CustomerHandlers struct {
	service service.CustomerService
}

// swagger:operation GET /customers Customers list
//
// List customers.
//
// This will return all customers, active and inactive
//
// ---
// produces:
// - application/json
// responses:
//   '200':
//     description: gets all customers from database
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/customerResponse"
//   '500':
//     description: when database down
func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	status := r.URL.Query().Get("status")

	customers, err := ch.service.GetAllCustomer(status)

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customers)
	}
}

// swagger:route GET /customers/{customer_id} Customers get
//
// Get a single customer.
//
// This will return a single customer, by Id
//
//     Produces:
//     - application/json
//
//     Responses:
//       200: customerResponse
//       500: description: when database not available
func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
