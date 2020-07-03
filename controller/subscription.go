package controller

import (
	"log"
	"net/http"
	"payment-backend/models"
	"payment-backend/utils"

	"github.com/labstack/echo"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/paymentmethod"
	"github.com/stripe/stripe-go/price"
	"github.com/stripe/stripe-go/setupintent"
	"github.com/stripe/stripe-go/sub"
)

func CheckCustomer(billing *models.Billing) (customer_id *stripe.Customer, err error) {

	stripe.Key = utils.Config.Key
	cu_search_params := &stripe.CustomerListParams{}
	cu_search_params.Filters.AddFilter("email", "", billing.Email)
	log.Print(cu_search_params)
	cu_response := customer.List(cu_search_params)
	log.Print(cu_response)
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

	cu, err := CheckCustomer(body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	log.Print(cu)

	intent_params := &stripe.SetupIntentParams{
		PaymentMethodTypes: []*string{
			stripe.String(body.Type),
		},
		Customer: stripe.String(cu.ID),
	}
	si, err := setupintent.New(intent_params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	data := models.CheckoutData{
		ClientSecret: si.ClientSecret,
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
	pm := pm_list[0]
	default_params := &stripe.CustomerParams{
		InvoiceSettings: &stripe.CustomerInvoiceSettingsParams{
			DefaultPaymentMethod: stripe.String(pm.ID),
		},
	}

	cu, _ = customer.Update(cu.ID, default_params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, i)
	}

	//product_params := &stripe.ProductParams{
	//	Name: stripe.String("Supporting Membership"),
	//	Type: stripe.String(string(stripe.ProductTypeService)),
	//}
	//log.Print("customer", cu)
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

	p, _ := price.New(price_params)
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
	s, _ := sub.New(params)
	return c.JSON(http.StatusCreated, s)
}
