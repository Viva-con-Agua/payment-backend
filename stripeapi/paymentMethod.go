package stripeapi

import (
	"errors"
	"payment-backend/models"
	"payment-backend/utils"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentmethod"
	"github.com/stripe/stripe-go/setupintent"
)

func SetupIntent(cu_id string, pm_type string) (co_data *models.CheckoutData, err error) {

	stripe.Key = utils.Config.Key
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
	co_data = new(models.CheckoutData)
	co_data.ClientSecret = si.ClientSecret
	return co_data, err

}

func GetPaymentMethod(cu_id string, pm_type string) (pm_method *stripe.PaymentMethod, err error) {
	stripe.Key = utils.Config.Key

	pm_params := &stripe.PaymentMethodListParams{
		Customer: stripe.String(cu_id),
		Type:     stripe.String(pm_type),
	}
	var pm_list []*stripe.PaymentMethod
	i := paymentmethod.List(pm_params)
	for i.Next() {
		pm_list = append(pm_list, i.PaymentMethod())
	}
	if pm_list == nil {
		err = errors.New("No Payment Method")
		return nil, err
	}
	return pm_list[0], nil
}
