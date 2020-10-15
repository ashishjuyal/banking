package domain

import "github.com/ashishjuyal/banking/errs"

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ById(string) (*Customer, *errs.AppError)
}
