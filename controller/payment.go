package controller

import (
	"net/http"
	"payment-backend/models"
	"payment-backend/utils"

	"github.com/labstack/echo"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/paymentintent"
)

func PaymentIntentCard(c echo.Context) (err error) {
	// create body as models.Role
	body := new(models.Money)
	// save data to body
	if err = c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// validate body
	if err = c.Validate(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	stripe.Key = utils.Config.Key

	params := &stripe.PaymentIntentParams{
		Amount:   &body.Amount,
		Currency: &body.Currency,
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
	// create body as models.Role
	body := new(models.Money)
	// save data to body
	if err = c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// validate body
	if err = c.Validate(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	stripe.Key = utils.Config.Key
	cu_params := &stripe.CustomerParams{}
	cu, _ := customer.New(cu_params)
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
