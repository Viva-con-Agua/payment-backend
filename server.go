package main

import (
	"log"
	"os"
	"payment-backend/controller"
	"strings"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
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
	godotenv.Load()

	log.Print(os.Getenv("STRIPE_KEY"))
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: strings.Split(os.Getenv("ALLOW_ORIGINS"), ","),
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Validator = &CustomValidator{validator: validator.New()}
	e.POST("/v1/payment/card", controller.PaymentIntentCard)
	e.POST("/v1/payment/iban", controller.PaymentIntentIBAN)
	e.POST("/v1/payment/success", controller.SuccessPayment)
	e.POST("/v1/payment/default", controller.AddDefaultPayment)
	e.POST("/v1/payment/subscription", controller.Subscription)
	e.Logger.Fatal(e.Start(":1323"))
}
