package middlewareutils

import (
	"github.com/Tsuryu/arreglapp-transactions/app/middlewares"
	"github.com/gin-gonic/gin"
)

var middlewarelist = map[string][]gin.HandlerFunc{
	"health": []gin.HandlerFunc{
		middlewares.Healthcheck,
	},
	"postTransaction": []gin.HandlerFunc{
		middlewares.PostTransaction,
	},
	"postTransactionDetail": []gin.HandlerFunc{
		middlewares.PostTransactionDetail,
		middlewares.CloseTransaction,
	},
}

// GetMiddlewares : get array of middlewares by name
func GetMiddlewares(name string) []gin.HandlerFunc {
	return middlewarelist[name]
}
