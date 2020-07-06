package controller

import (
	"errors"
	"log"
	"net/http"
	"payment-backend/models"
	"payment-backend/stripeapi"
	"payment-backend/utils"

	"github.com/labstack/echo"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/paymentmethod"
)

/**
* Get or create customer. Use billing informations.
* return stripe Customer
 */
func CheckCustomer(billing *models.Billing) (customer_id *stripe.Customer, err error) {
	stripe.Key = utils.Config.Key
	cu_search_params := &stripe.CustomerListParams{}
	cu_search_params.Filters.AddFilter("email", "", billing.Email)
	cu_response := customer.List(cu_search_params)
	var cu_list []*stripe.Customer
	for cu_response.Next() {
		cu_list = append(cu_list, cu_response.Customer())
	}
	if cu_list == nil {
		cu_params := &stripe.CustomerParams{
			Email:            stripe.String(string(billing.Email)),
			Name:             stripe.String(string(billing.Name)),
			PreferredLocales: []*string{stripe.String(billing.Locale)},
		}

		cu, err := customer.New(cu_params)
		//	cu_list = append(cu_list, cu)
		//	log.Print(cu, err)
		return cu, err
	} else {
		cu := cu_list[0]
		return cu, nil
	}
}

/**
* get payment method from stripe api via customer id and given type
 */
func GetPaymentMethod(cu_id string, pm_type string) (pm_method *stripe.PaymentMethod, err error) {
	stripe.Key = utils.Config.Key

	pm_params := &stripe.PaymentMethodListParams{
		Customer: stripe.String(cu_id),
		Type:     stripe.String(pm_type),
	}
	var pm_list []*stripe.PaymentMethod
	i := paymentmethod.List(pm_params)
	for i.Next() {
		pm_list = append(pm_list, i.PaymentMethod())
	}
	if pm_list == nil {
		err = errors.New("No Payment Method")
		return nil, err
	}
	return pm_list[0], nil
}

/**
* controller for payment
 */
func AddDefaultPayment(c echo.Context) (err error) {

	stripe.Key = utils.Config.Key
	//handle body as Subscription
	body := new(models.Billing)
	if err = c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err = c.Validate(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// create default response model
	response := new(models.ResponseMessage)

	//get or create customer. Need to be repleaced
	cu, err := stripeapi.GetCustomerByEmail(body.Email)
	if cu == nil {
		cu, err = stripeapi.CreateCustomer(body.Email, body.Name, body.Locale)
		if err != nil {
			log.Print("Error CheckCustomer: \n", err)
			response.Message = "Fail creating or get Customer"
			return c.JSON(http.StatusBadRequest, response)
		}
	}

	//add payment method
	data, err := stripeapi.SetupIntent(cu.ID, body.Type)
	if err != nil {
		log.Print("Error setupintent New: \n", err)
		response.Message = "Fail setup intent"
		return c.JSON(http.StatusBadRequest, response)
	}
	return c.JSON(http.StatusOK, data)
}

func Subscription(c echo.Context) (err error) {
	stripe.Key = utils.Config.Key
	//handle body as Subscription
	body := new(models.Billing)
	if err = c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err = c.Validate(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	cu, err := stripeapi.GetCustomerByEmail(body.Email)
	if cu == nil {
		log.Print("Error CheckCustomer: \n", err)
		response := new(models.ResponseMessage)
		response.Message = "Failed get Customer"
		return c.JSON(http.StatusBadRequest, response)
	}
	pm, err := stripeapi.GetPaymentMethod(cu.ID, body.Type)
	if pm == nil {
		log.Print("Error GetPaymentMethode: \n", err)
		response := new(models.ResponseMessage)
		return c.JSON(http.StatusBadRequest, response)
	}

	cu, err = stripeapi.SetDefaultPayment(cu.ID, pm.ID)
	if err != nil {
		response := new(models.ResponseMessage)
		response.Message = "Failed set default payment"
		return c.JSON(http.StatusBadRequest, response)

	}
	p, err := stripeapi.CreatePrice(body.Amount)
	if p != nil {
		log.Print("Error CreatePrice: \n", err)
		response := new(models.ResponseMessage)
		response.Message = "failed creating price"
		return c.JSON(http.StatusBadRequest, response)
	}
	s, err := stripeapi.SubProduct(cu.ID, p.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, s)
}
