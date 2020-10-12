package middlewares

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/Tsuryu/arreglapp-transactions/app/db/transactions"
	"github.com/Tsuryu/arreglapp-transactions/app/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// PostTransaction : create new transaction
func PostTransaction(context *gin.Context) {
	transaction := models.Transaction{}
	context.ShouldBindBodyWith(&transaction, binding.JSON)

	if context.GetHeader("with-security-code") != "" {
		transaction.SecurityCode = fmt.Sprintf("%02d", rand.Intn(999999-0))
	}

	err := transactions.Insert(&transaction)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create transaction", "transaction": transaction})
		context.Abort()
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "created", "transaction": transaction})
}
