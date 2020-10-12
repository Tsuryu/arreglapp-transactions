package routers

import (
	middlewareutils "github.com/Tsuryu/arreglapp-transactions/app/middlewares/utils"
	"github.com/gin-gonic/gin"
)

// Router : app routes
func Router() {
	router := gin.Default()
	router.GET("/health", middlewareutils.GetMiddlewares("health")...)
	router.POST("/transaction/:id/detail", middlewareutils.GetMiddlewares("postTransactionDetail")...)
	router.POST("/transaction", middlewareutils.GetMiddlewares("postTransaction")...)
	router.Run(":8080")
}
