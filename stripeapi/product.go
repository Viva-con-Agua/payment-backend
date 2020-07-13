package stripeapi

import (
	"errors"
	"payment-backend/utils"
	"time"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/price"
	"github.com/stripe/stripe-go/sub"
)

func GetDate() (t time.Time) {
	var time_now = time.Now()
	month := time_now.Month()
	year := time_now.Year()
	now_time := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	start_time := now_time.AddDate(0, 1, 0)
	return start_time

}

func CreatePrice(amount int64, pm_id string, interval string) (p *stripe.Price, err error) {

	var freq string
	if interval == "month" {
		freq = string(stripe.PriceRecurringIntervalMonth)
	} else if interval == "year" {
		freq = string(stripe.PriceRecurringIntervalYear)
	} else {
		err = errors.New("no interval")
		return nil, err
	}
	stripe.Key = utils.Config.Key
	params := &stripe.PriceParams{
		Nickname:   stripe.String("Standard Monthly"),
		Product:    stripe.String(pm_id),
		UnitAmount: stripe.Int64(amount),
		Currency:   stripe.String(string(stripe.CurrencyEUR)),
		Recurring: &stripe.PriceRecurringParams{
			Interval:  stripe.String(freq),
			UsageType: stripe.String(string(stripe.PriceRecurringUsageTypeLicensed)),
		},
	}

	p, err = price.New(params)
	return p, err
}

func SubProduct(cu_id string, p_id string) (s *stripe.Subscription, err error) {

	stripe.Key = utils.Config.Key
	start_time := GetDate()
	billing_time := start_time.AddDate(0, 0, 4)
	params := &stripe.SubscriptionParams{
		Customer: stripe.String(cu_id),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Price:    stripe.String(p_id),
				Quantity: stripe.Int64(1),
			},
		},
		TrialEnd:           stripe.Int64(billing_time.Unix()),
		BillingCycleAnchor: stripe.Int64(billing_time.Unix()),
	}
	params.AddExpand("latest_invoice.payment_intent")
	s, err = sub.New(params)
	return s, err

}
