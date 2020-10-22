package transactions

import (
	"context"
	"fmt"
	"time"

	"github.com/Tsuryu/arreglapp-commons/app/models"
	"github.com/Tsuryu/arreglapp-commons/app/utils"
)

// Insert : Creates new transaction
func Insert(transaction *models.Transaction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	transaction.Active = true
	random, err := utils.RandomHex(16)
	if err != nil {
		return err
	}
	transaction.TraceID = random

	if transaction.Details != nil {
		transaction.Details[0].Date = time.Now()
	}

	_, err = Collection.InsertOne(ctx, transaction)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}
