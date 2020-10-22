package transactions

import (
	"context"
	"fmt"
	"time"

	"github.com/Tsuryu/arreglapp-commons/app/models"
	"go.mongodb.org/mongo-driver/bson"
)

// Update : adds detail to a transaction
func Update(transaction *models.Transaction) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// map string, interface
	register := make(map[string]interface{})
	if !transaction.Active {
		register["active"] = transaction.Active
	}
	updateString := bson.M{
		"$set": register,
	}

	filter := bson.M{
		"trace_id": bson.M{
			"$eq": transaction.TraceID,
		},
		"security_code": bson.M{
			"$eq": transaction.SecurityCode,
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
