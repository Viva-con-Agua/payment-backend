package papi

import (
	"os"

	"github.com/Viva-con-Agua/vcago/verr"
	"github.com/plutov/paypal/v3"
)

//PlanInsert inserts a billing Plan to the paypal api
func PlanInsert(c *paypal.Client, p *paypal.SubscriptionPlan) (*paypal.CreateSubscriptionPlanResponse, *verr.APIError) {
	c, err := paypal.NewClient(os.Getenv("PAYPAL_CLIENT_ID"), os.Getenv("PAYPAL_SECRET"), os.Getenv("PAYPAL_BASE_URL"))
	if err != nil {
		return nil, verr.NewAPIError(err).InternalServerError()
	}
	_, err = c.GetAccessToken()
	if err != nil {
		return nil, verr.NewAPIError(err).InternalServerError()
	}
	resp, err := c.CreateSubscriptionPlan(*p)
	if err != nil {
		return nil, verr.NewAPIError(err).InternalServerError()
	}
	return resp, nil
}
