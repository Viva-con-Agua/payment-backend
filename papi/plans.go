package papi

import (
	"github.com/Viva-con-Agua/vcago/verr"
	"github.com/plutov/paypal/v3"
)

//PlanInsert inserts a billing Plan to the paypal api
func PlanInsert(c *paypal.Client, p *paypal.BillingPlan) (*paypal.CreateBillingResp, *verr.APIError) {
	resp, err := c.CreateBillingPlan(*p)
	if err != nil {
		return nil, verr.NewAPIError(err).InternalServerError()
	}
	return resp, nil
}
