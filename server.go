package main

import (
	"log"
	"os"
	"payment-backend/controllers"
	"payment-backend/dao"
	"strings"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	dao.Connect()
	log.Print(os.Getenv("STRIPE_KEY"))
	e := echo.New()
	m := middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     strings.Split(os.Getenv("ALLOW_ORIGINS"), ","),
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	})
	log.Print(strings.Split(os.Getenv("ALLOW_ORIGINS"), ","))
	e.Use(m)
	e.Validator = &CustomValidator{validator: validator.New()}
	e.POST("/v1/payment/card", controllers.PaymentIntentCard)
	e.POST("/v1/payment/iban", controllers.PaymentIntentIBAN)
	e.POST("/v1/payment/success", controllers.SuccessPayment)
	e.POST("/v1/payment/default", controllers.AddDefaultPayment)
	e.POST("/v1/payment/subscription", controllers.Subscription)
	e.POST("/v1/payment/paypal/subscription", controllers.CreatePaypalBillingPlan)
	if port, ok := os.LookupEnv("REPO_PORT"); ok {
		e.Logger.Fatal(e.Start(":" + port))
	} else {
		e.Logger.Fatal(e.Start(":1323"))
	}
}
