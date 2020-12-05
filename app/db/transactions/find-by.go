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

// CommonFindBy : get user transactions
func CommonFindBy(condition bson.M) ([]models.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var results []models.Transaction

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
			if result.Details[index].Metadata != nil {
				result.Details[index].Metadata = utils.GetObjectFromPrimitiveD(detail.Metadata)
			}
		}
		results = append(results, result)
	}
	if err := cur.Err(); err != nil {
		fmt.Println(err)
	}

	return results, nil
}
