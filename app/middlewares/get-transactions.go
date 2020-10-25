package middlewares

import (
	"net/http"

	"github.com/Tsuryu/arreglapp-transactions/app/db/transactions"
	"github.com/gin-gonic/gin"
)

// GetTransactions : fetch a transactions by username
func GetTransactions(context *gin.Context) {
	transactions, err := transactions.FindByUsername(context.Query("username"))
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Transactions by " + context.Query("username") + " have not been found"})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, transactions)
}
