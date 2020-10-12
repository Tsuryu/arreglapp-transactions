package middlewares

import (
	"github.com/Tsuryu/arreglapp-transactions/app/db/transactions"
	"github.com/Tsuryu/arreglapp-transactions/app/models"
	"github.com/gin-gonic/gin"
)

// CloseTransaction : finish a transaction to make it not updatable
func CloseTransaction(context *gin.Context) {
	if context.GetHeader("last-update") != "" {
		transaction := models.Transaction{}
		transaction.TraceID = context.Param("id")
		transaction.SecurityCode = context.GetHeader("security-code")
		transactions.Update(&transaction)
	}
}
