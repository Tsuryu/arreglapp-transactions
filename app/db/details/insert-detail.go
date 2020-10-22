package details

import (
	"context"
	"fmt"
	"time"

	"github.com/Tsuryu/arreglapp-commons/app/models"
	"go.mongodb.org/mongo-driver/bson"
)

// Insert : adds detail to a transaction
func Insert(traceID string, securityCode string, detail *models.TransactionDetail) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if detail.Date.IsZero() {
		detail.Date = time.Now()
	}

	updateString := bson.M{
		"$addToSet": bson.M{
			"details": bson.M{
				"$each": []models.TransactionDetail{(*detail)},
			},
		},
	}

	filter := bson.M{
		"trace_id": bson.M{
			"$eq": traceID,
		},
		"security_code": bson.M{
			"$eq": securityCode,
		},
		"active": bson.M{
			"$eq": true,
		},
	}

	result, err := Collection.UpdateOne(ctx, filter, updateString)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	return result.ModifiedCount, err
}
