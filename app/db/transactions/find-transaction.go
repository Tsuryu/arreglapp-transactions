package transactions

import (
	"context"
	"time"

	"github.com/Tsuryu/arreglapp-commons/app/models"
	"github.com/Tsuryu/arreglapp-commons/app/utils"
	"go.mongodb.org/mongo-driver/bson"
)

// FindBy : fetch a transaction by id
func FindBy(id string) (models.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"trace_id": id}

	var result models.Transaction

	err := Collection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return result, err
	}

	for index, detail := range result.Details {
		result.Details[index].Metadata = utils.GetObjectFromPrimitiveD(detail.Metadata)
	}

	return result, nil
}
