package dao

import (
	"context"
	"payment-backend/models"
	"strings"
	"time"

	"github.com/Viva-con-Agua/vcago/verr"
)

//ProductInsertOne inserts
func ProductInsertOne(p *models.Product) (*models.Product, *verr.APIError) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var coll = DB.Collection("products")
	if _, err := coll.InsertOne(ctx, p); err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, verr.NewAPIError(err).Conflict("product_duplicate")
		}
		return nil, verr.NewAPIError(err).InternalServerError()
	}
	return p, nil
}
