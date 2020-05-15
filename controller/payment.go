package controller

import (
	"net/http"
	"payment-backend/models"
	"payment-backend/utils"

	"github.com/labstack/echo"
	"github.com/stripe/stripe-go"
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
