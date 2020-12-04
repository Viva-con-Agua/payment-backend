package models

import (
	"github.com/google/uuid"
	"github.com/plutov/paypal/v3"
)

type (
	//PaypalBillingPlan represents the essential billing plan informations for frontend subscription handling
	PaypalBillingPlan struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Interval    string `json:"interval"`
		Amount      string `json:"amount"`
		Currency    string `json:"currency"`
		ProductID   string `json:"product_id"`
	}
	//PaypalBillingPlanDB represents the database models for storing PaypalBillingPlan.
	PaypalBillingPlanDB struct {
		ID          string `bson:"_id" json:"id" validate:"required"`
		Name        string `bson:"name" json:"name"`
		Description string `bson:"description" json:"description"`
		Amount      string `bson:"amount" json:"amount" validate:"required"`
		Currency    string `bson:"currency" json:"currency" validate:"required"`
		Interval    string `bson:"interval" json:"interval" validate:"required"`
		ProductID   string `bson:"product_id" json:"product_id" validate:"required"`
		PlanID      string `bson:"plan_id" json:"plan_id" validate:"required"`
	}
)

//PaypalBillingPlanDB creates a new PaypalBillingPlanDB based on b.
func (b *PaypalBillingPlan) PaypalBillingPlanDB(p string) *PaypalBillingPlanDB {
	return &PaypalBillingPlanDB{
		ID:          uuid.New().String(),
		Name:        b.Name,
		Description: b.Description,
		PlanID:      p,
		ProductID:   b.ProductID,
		Amount:      b.Amount,
		Currency:    b.Currency,
		Interval:    b.Interval,
	}
}

//SubscriptionPlan converts b into an paypal SubscriptionPlan
func (b *PaypalBillingPlan) SubscriptionPlan() *paypal.SubscriptionPlan {
	var interval paypal.IntervalUnit
	if b.Interval == "month" {
		interval = paypal.IntervalUnitMonth
	} else {
		interval = paypal.IntervalUnitYear
	}
	money := &paypal.Money{
		Value:    b.Amount,
		Currency: b.Currency,
	}

	return &paypal.SubscriptionPlan{
		ProductId:   b.ProductID,
		Name:        b.Name,
		Status:      paypal.SubscriptionPlanStatusActive,
		Description: "description",
		BillingCycles: []paypal.BillingCycle{
			paypal.BillingCycle{
				Frequency: paypal.Frequency{
					IntervalUnit:  interval,
					IntervalCount: 1,
				},
				TenureType:  "REGULAR",
				Sequence:    1,
				TotalCycles: 0,
				PricingScheme: paypal.PricingScheme{
					FixedPrice: *money,
				},
			},
		},
		PaymentPreferences: &paypal.PaymentPreferences{
			AutoBillOutstanding: false,
			SetupFee: &paypal.Money{
				Value:    "0.0",
				Currency: "EUR",
			},
			SetupFeeFailureAction: paypal.SetupFeeFailureActionContinue,
		},
		Taxes: &paypal.Taxes{
			Percentage: "0",
			Inclusive:  true,
		},
	}

}
