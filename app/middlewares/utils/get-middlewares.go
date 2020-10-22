package middlewareutils

import (
	commonMiddlewares "github.com/Tsuryu/arreglapp-commons/app/middlewares"
	"github.com/Tsuryu/arreglapp-transactions/app/middlewares"
	"github.com/gin-gonic/gin"
)

var middlewarelist = map[string][]gin.HandlerFunc{
	"health": []gin.HandlerFunc{
		commonMiddlewares.Healthcheck,
	},
	"postTransaction": []gin.HandlerFunc{
		middlewares.PostTransaction,
		middlewares.SendEmail,
	},
	"postTransactionDetail": []gin.HandlerFunc{
		middlewares.PostTransactionDetail,
		middlewares.CloseTransaction,
	},
	"getTransaction": []gin.HandlerFunc{
		middlewares.GetTransaction,
	},
}

// GetMiddlewares : get array of middlewares by name
func GetMiddlewares(name string) []gin.HandlerFunc {
	return middlewarelist[name]
}
