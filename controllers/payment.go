package controllers

import (
	"log"
	"net/http"
	"os"
	"payment-backend/models"

	"github.com/labstack/echo/v4"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/paymentintent"
)

func PaymentIntentCard(c echo.Context) (err error) {
	// create body as models.Role
	body := new(models.CreditCard)
	// save data to body
	if err = c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// validate body
	if err = c.Validate(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	stripe.Key = os.Getenv("STRIPE_KEY")

	cu_params := &stripe.CustomerParams{
		Email:            stripe.String(string(body.Email)),
		Name:             stripe.String(string(body.Name)),
		PreferredLocales: []*string{stripe.String(body.Locale)},
	}
	cu, err := customer.New(cu_params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	params := &stripe.PaymentIntentParams{
		Amount:   &body.Amount,
		Currency: &body.Currency,
		Customer: stripe.String(cu.ID),
		PaymentMethodTypes: []*string{
			stripe.String("card"),
		},
	}
	pi, _ := paymentintent.New(params)
	data := models.CheckoutData{
		ClientSecret: pi.ClientSecret,
	}
	return c.JSON(http.StatusOK, &data)
}

func PaymentIntentIBAN(c echo.Context) (err error) {
	log.Print(os.Getenv("GOT IBAN REQUEST"))
	// create body as models.Role
	body := new(models.Sepa)
	// save data to body
	if err = c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// validate body
	if err = c.Validate(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	stripe.Key = os.Getenv("STRIPE_KEY")

	cu_params := &stripe.CustomerParams{
		Email:            stripe.String(string(body.Email)),
		Name:             stripe.String(string(body.Name)),
		PreferredLocales: []*string{stripe.String(body.Locale)},
	}
	cu, err := customer.New(cu_params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	params := &stripe.PaymentIntentParams{
		Amount:           &body.Amount,
		Currency:         &body.Currency,
		SetupFutureUsage: stripe.String(string(stripe.PaymentIntentSetupFutureUsageOffSession)),
		Customer:         stripe.String(cu.ID),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"sepa_debit",
		}),
	}
	// Verify your integration in this guide by including this parameter
	params.AddMetadata("integration_check", "sepa_debit_accept_a_payment")
	pi, _ := paymentintent.New(params)
	data := models.CheckoutData{
		ClientSecret: pi.ClientSecret,
	}
	return c.JSON(http.StatusOK, &data)
}

func SuccessPayment(c echo.Context) (err error) {
	body := new(models.Payment)
	if err = c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// validate body
	if err = c.Validate(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, &body)

}
