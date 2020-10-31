package transactions

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Tsuryu/arreglapp-commons/app/models"
	"github.com/Tsuryu/arreglapp-commons/app/utils"
	"go.mongodb.org/mongo-driver/bson"
)

// FindByUsername : get user transactions
func FindByUsername(username string) ([]models.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"username": username}

	var results []models.Transaction
	// var metadata = make(map[string]interface{})

	cur, err := Collection.Find(ctx, condition)
	if err != nil {
		log.Println(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result models.Transaction
		err := cur.Decode(&result)
		if err != nil {
			fmt.Println(err)
		}
		result.SecurityCode = ""
		for index, detail := range result.Details {
			result.Details[index].Metadata = utils.GetObjectFromPrimitiveD(detail.Metadata)
		}
		results = append(results, result)
	}
	if err := cur.Err(); err != nil {
		fmt.Println(err)
	}

	return results, nil
}
