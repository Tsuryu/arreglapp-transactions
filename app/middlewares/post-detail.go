package middlewares

import (
	"net/http"

	"github.com/Tsuryu/arreglapp-transactions/app/db/details"
	"github.com/Tsuryu/arreglapp-transactions/app/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// PostTransactionDetail : create new transaction
func PostTransactionDetail(context *gin.Context) {
	transactionDetail := models.TransactionDetail{}
	context.ShouldBindBodyWith(&transactionDetail, binding.JSON)
	traceID := context.Param("id")

	modifiedCount, err := details.Insert(traceID, context.GetHeader("security-code"), &transactionDetail)
	if modifiedCount == 0 {
		context.JSON(http.StatusNotFound, gin.H{"message": "Failed to add details", "transaction_detail": transactionDetail})
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create transaction", "transaction_detail": transactionDetail})
		context.Abort()
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "created", "transaction_detail": transactionDetail})
}
