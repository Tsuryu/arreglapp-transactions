package transactions

import (
	"context"
	"time"

	"github.com/Tsuryu/arreglapp-transactions/app/models"
	"go.mongodb.org/mongo-driver/bson"
)

// FindBy : fetch a transaction by id
func FindBy(id string) (*models.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"trace_id": id}

	var result *models.Transaction

	err := Collection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
