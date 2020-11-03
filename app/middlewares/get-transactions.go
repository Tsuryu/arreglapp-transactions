package middlewares

import (
	"net/http"

	"github.com/Tsuryu/arreglapp-transactions/app/db/transactions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// GetTransactions : fetch a transactions by username
func GetTransactions(context *gin.Context) {
	condition := bson.M{}

	username := context.Query("username")
	if username != "" {
		condition["username"] = username
	}

	reference := context.Query("reference")
	if reference != "" {
		condition["reference"] = reference
	}

	transactionsList, err := transactions.CommonFindBy(condition)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Transactions have not been found"})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, transactionsList)
}
