package middlewareutils

import (
	commonMiddlewares "github.com/Tsuryu/arreglapp-commons/app/middlewares"
	"github.com/Tsuryu/arreglapp-transactions/app/middlewares"
	"github.com/gin-gonic/gin"
)

var middlewarelist = map[string][]gin.HandlerFunc{
	"health": {
		commonMiddlewares.Healthcheck,
	},
	"postTransaction": {
		middlewares.PostTransaction,
		middlewares.SendEmail,
	},
	"postTransactionDetail": {
		middlewares.PostTransactionDetail,
		middlewares.CloseTransaction,
	},
	"getTransaction": {
		middlewares.GetTransaction,
	},
	"getTransactions": {
		middlewares.GetTransactions,
	},
}

// GetMiddlewares : get array of middlewares by name
func GetMiddlewares(name string) []gin.HandlerFunc {
	return middlewarelist[name]
}
