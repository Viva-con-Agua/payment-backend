package stripeapi

import (
	"payment-backend/models"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/setupintent"
)

func SetupIntent(cu_id string, pm_type string) (co_data *models.CheckoutData, err error) {
	params := &stripe.SetupIntentParams{
		PaymentMethodTypes: []*string{
			stripe.String(pm_type),
		},
		Customer: stripe.String(cu_id),
	}
	si, err := setupintent.New(params)
	if err != nil {
		return nil, err
	}
	co_data.ClientSecret = si.ClientSecret
	return co_data, err

}
