package stripeapi

import (
	"payment-backend/utils"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/price"
	"github.com/stripe/stripe-go/sub"
)

func CreatePrice(amount int64) (p *stripe.Price, err error) {

	stripe.Key = utils.Config.Key
	params := &stripe.PriceParams{
		Nickname:   stripe.String("Standard Monthly"),
		Product:    stripe.String("prod_HZW4PLYJeuxnyC"),
		UnitAmount: stripe.Int64(amount),
		Currency:   stripe.String(string(stripe.CurrencyEUR)),
		Recurring: &stripe.PriceRecurringParams{
			Interval:  stripe.String(string(stripe.PriceRecurringIntervalMonth)),
			UsageType: stripe.String(string(stripe.PriceRecurringUsageTypeLicensed)),
		},
	}

	p, err = price.New(params)
	return p, err
}

func SubProduct(cu_id string, p_id string) (s *stripe.Subscription, err error) {

	stripe.Key = utils.Config.Key
	params := &stripe.SubscriptionParams{
		Customer: stripe.String(cu_id),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Price:    stripe.String(p_id),
				Quantity: stripe.Int64(1),
			},
		},
	}

	params.AddExpand("latest_invoice.payment_intent")
	s, err = sub.New(params)
	return s, err

}
