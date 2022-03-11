package dto

// Customer response details
// swagger:model customerResponse
type CustomerResponse struct {
	// Customer Id
	Id string `json:"customer_id"`
	// Full name of customer
	Name string `json:"full_name"`
	// City of customer
	City string `json:"city"`
	// Zipcode of customer city
	Zipcode string `json:"zipcode"`
	// Date of birth of customer in YYYY-MM-DD format
	DateofBirth string `json:"date_of_birth"`
	// Tells if the customer is active or not
	Status string `json:"status"`
}
