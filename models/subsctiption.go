package models

import "github.com/plutov/paypal/v3"

type (
	//Billing stripe billing information
	Billing struct {
		Email    string `json:"email"`
		Name     string `json:"name"`
		Interval string `json:"interval"`
		Locale   string `json:"locale"`
		Amount   int64  `json:"amount"`
		Currency string `json:"currency"`
		Type     string `json:"type"`
		Product  string `json:"product"`
	}

	//ResponseMessage message for response
	ResponseMessage struct {
		Message string `json:"message"`
	}

	//PaypalProduct represents an model for creating an product by using paypal api.
	PaypalProduct struct {
		Name        string `json:"name"`
		Description string `json:"description" validate:"required"`
		Type        string `bson:"type" json:"type" validate:"required"`
		Category    string `json:"category" bson:"category"`
		ImageURL    string `json:"image_url"`
		HomeURL     string `json:"home_url"`
	}
)

func (b *Billing) Paypal() *paypal.BillingPlan {
	return &paypal.BillingPlan{
		ProductID: b.Product,
		Name:      b.Name,
	}
}
