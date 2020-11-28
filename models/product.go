package models

import (
	"github.com/google/uuid"
	"github.com/plutov/paypal/v3"
	"github.com/stripe/stripe-go/v71"
)

type (
	//ProductCreate represents the inital product json.
	ProductCreate struct {
		Name        string `json:"name"`
		Description string `json:"description" validate:"required"`
		ImageURL    string `json:"image_url"`
		HomeURL     string `json:"home_url"`
	}
	//Product is the internal database representation of payment products
	Product struct {
		ID          string         `bson:"_id" json:"id" validate:"required"`
		Name        string         `bson:"name" json:"name"`
		Description string         `bson:"description" json:"description" validate:"required"`
		ImageURL    string         `json:"image_url"`
		HomeURL     string         `json:"home_url"`
		Paypal      paypal.Product `bson:"paypal" json:"paypal" validate:"required"`
		Stripe      stripe.Product `bson:"stripe" json:"stripe" validate:"required"`
	}
)

//Paypal creates a paypal.Product from p.
func (p *ProductCreate) Paypal() *paypal.Product {
	return &paypal.Product{
		Name:        p.Name,
		Description: p.Description,
		Type:        paypal.ProductTypeDigital,
		Category:    paypal.ProductCategorySoftware,
		ImageUrl:    p.ImageURL,
		HomeUrl:     p.HomeURL,
	}
}

//Product returns a new Product based on p. The function simply adds a new UUID.
func (p *ProductCreate) Product() *Product {
	return &Product{
		ID:          uuid.New().String(),
		Name:        p.Name,
		Description: p.Description,
		ImageURL:    p.ImageURL,
		HomeURL:     p.HomeURL,
	}
}
