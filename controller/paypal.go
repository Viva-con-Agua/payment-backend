package controller

import (
	"net/http"
	"payment-backend/models"
	"payment-backend/papi"

	"github.com/Viva-con-Agua/vcago/verr"
	"github.com/labstack/echo"
)

func PaypalCreateBilling(c echo.Context) error {
	body := new(models.Billing)
	if apiErr := verr.JSONValidate(c, body); apiErr != nil {
		return c.JSON(apiErr.Code, apiErr.Body)
	}
	c, apiErr := papi.NewClient()
	p, apiErr := papi.PlanInsert(c, body.Paypal())

	return c.JSON(http.StatusCreated, _)
}
