package main

import (
	"payment-backend/controller"
	"payment-backend/utils"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
func main() {
	utils.LoadConfig()
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Validator = &CustomValidator{validator: validator.New()}
	e.POST("/api/v1/payment/card", controller.PaymentIntentCard)
	e.Logger.Fatal(e.Start(":1323"))
}
