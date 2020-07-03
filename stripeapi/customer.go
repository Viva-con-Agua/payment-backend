package stripeapi

import (
	"payment-backend/utils"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
)

func CreateCustomer(email string, name string, locale string) (cu *stripe.Customer, err error) {
	stripe.Key = utils.Config.Key
	params := &stripe.CustomerParams{
		Email:            stripe.String(string(email)),
		Name:             stripe.String(string(name)),
		PreferredLocales: []*string{stripe.String(locale)},
	}
	cu, err = customer.New(params)
	return cu, err
}

func GetCustomerByEmail(email string) (cu *stripe.Customer, err error) {
	stripe.Key = utils.Config.Key
	params := &stripe.CustomerListParams{}
	params.Filters.AddFilter("email", "", email)
	response := customer.List(params)
	var cu_list []*stripe.Customer
	for response.Next() {
		cu_list = append(cu_list, response.Customer())
	}
	if cu_list == nil {
		return nil, err
	} else {
		return cu_list[0], err
	}
}
