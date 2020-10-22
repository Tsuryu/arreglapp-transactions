package middlewares

import (
	"github.com/Tsuryu/arreglapp-commons/app/models"
	"github.com/Tsuryu/arreglapp-commons/app/utils"
	"github.com/gin-gonic/gin"
)

// SendEmail : send email with otp
func SendEmail(context *gin.Context) {
	transaction := context.Keys["transaction"].(models.Transaction)

	if transaction.Email != "" && transaction.SecurityCode != "" {
		go utils.SendEmail(transaction.Email, transaction.SecurityCode)
	}
}
