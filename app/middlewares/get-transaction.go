package middlewares

import (
	"net/http"

	"github.com/Tsuryu/arreglapp-transactions/app/db/transactions"
	"github.com/gin-gonic/gin"
)

// GetTransaction : fetch a transaction
func GetTransaction(context *gin.Context) {
	transaction, err := transactions.FindBy(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Transaction " + context.Param("id") + " has not been found"})
		context.Abort()
		return
	}

	transaction.SecurityCode = ""

	context.JSON(http.StatusOK, transaction)
}
