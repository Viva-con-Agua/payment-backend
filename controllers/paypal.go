package controllers

import (
	"net/http"
	"payment-backend/dao"
	"payment-backend/models"

	"github.com/Viva-con-Agua/vcago/verr"
	"github.com/labstack/echo/v4"
)

//CreatePaypalBillingPlan POST /v1/payment/paypal/subscribtion
func CreatePaypalBillingPlan(c echo.Context) (err error) {
	body := new(models.PaypalBillingPlan)
	if apiErr := verr.JSONValidate(c, body); apiErr != nil {
		return c.JSON(apiErr.Code, apiErr.Body)
	}
	resp, apiErr := dao.CreatePaypalBillingPlan(body)
	if apiErr != nil {
		apiErr.Log(c)
		return c.JSON(apiErr.Code, apiErr.Body)
	}
	return c.JSON(http.StatusCreated, resp)
}
