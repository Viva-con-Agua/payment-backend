package dao

import (
	"context"
	"os"
	"payment-backend/models"
	"strings"
	"time"

	"github.com/Viva-con-Agua/vcago/verr"
	"github.com/plutov/paypal/v3"
	"go.mongodb.org/mongo-driver/bson"
)

//CreatePaypalBillingPlan handles paypal subscription.
func CreatePaypalBillingPlan(b *models.PaypalBillingPlan) (*models.PaypalBillingPlanDB, *verr.APIError) {
	//search billing plan in database
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	coll := DB.Collection("paypal_billingplan")
	filter := bson.M{"amount": b.Amount, "currency": b.Currency, "interval": b.Interval, "product_id": b.ProductID}
	paypalDB := new(models.PaypalBillingPlanDB)
	err := coll.FindOne(ctx, filter).Decode(&paypalDB)
	if err != nil {
		if strings.Contains(err.Error(), "no documents in result") {
			//create new billing plan
			c, err := paypal.NewClient(os.Getenv("PAYPAL_CLIENT_ID"), os.Getenv("PAYPAL_SECRET"), os.Getenv("PAYPAL_BASE_URL"))
			if err != nil {
				return nil, verr.NewAPIError(err).InternalServerError()
			}
			_, err = c.GetAccessToken()
			if err != nil {
				return nil, verr.NewAPIError(err).InternalServerError()
			}
			resp, err := c.CreateSubscriptionPlan(*b.SubscriptionPlan())
			if err != nil {
				return nil, verr.NewAPIError(err).InternalServerError()
			}
			paypalDB = b.PaypalBillingPlanDB(resp.ID)
			_, err = coll.InsertOne(ctx, paypalDB)
			if err != nil {
				return nil, verr.NewAPIError(err).InternalServerError()
			}
		} else {
			return nil, verr.NewAPIError(err).InternalServerError()

		}
	}
	return paypalDB, nil

}
