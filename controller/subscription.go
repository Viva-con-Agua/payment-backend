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
	"github.com/stripe/stripe-go/price"
	"github.com/stripe/stripe-go/sub"
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

	cu, err := CheckCustomer(body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	pm_params := &stripe.PaymentMethodListParams{
		Customer: stripe.String(cu.ID),
		Type:     stripe.String(body.Type),
	}
	var pm_list []*stripe.PaymentMethod
	i := paymentmethod.List(pm_params)
	for i.Next() {
		pm_list = append(pm_list, i.PaymentMethod())
	}
	if pm_list == nil {
		response := new(models.ResponseMessage)
		response.Message = "no payment method"
		return c.JSON(http.StatusBadRequest, response)
	}
	pm := pm_list[0]
	default_params := &stripe.CustomerParams{
		InvoiceSettings: &stripe.CustomerInvoiceSettingsParams{
			DefaultPaymentMethod: stripe.String(pm.ID),
		},
	}

	cu, err = customer.Update(cu.ID, default_params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	price_params := &stripe.PriceParams{
		Nickname:   stripe.String("Standard Monthly"),
		Product:    stripe.String("prod_HZW4PLYJeuxnyC"),
		UnitAmount: stripe.Int64(body.Amount),
		Currency:   stripe.String(string(stripe.CurrencyEUR)),
		Recurring: &stripe.PriceRecurringParams{
			Interval:  stripe.String(string(stripe.PriceRecurringIntervalMonth)),
			UsageType: stripe.String(string(stripe.PriceRecurringUsageTypeLicensed)),
		},
	}

	p, err := price.New(price_params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	params := &stripe.SubscriptionParams{
		Customer: stripe.String(cu.ID),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Price:    stripe.String(p.ID),
				Quantity: stripe.Int64(1),
			},
		},
	}

	params.AddExpand("latest_invoice.payment_intent")
	s, err := sub.New(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, s)
}
