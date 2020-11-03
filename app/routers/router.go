package routers

import (
	"os"
	"strconv"

	middlewareutils "github.com/Tsuryu/arreglapp-transactions/app/middlewares/utils"
	"github.com/gin-gonic/gin"
)

// Router : app routes
func Router() {
	router := gin.Default()
	router.GET("/health", middlewareutils.GetMiddlewares("health")...)
	router.POST("/transaction/:id/detail", middlewareutils.GetMiddlewares("postTransactionDetail")...)
	router.POST("/transaction", middlewareutils.GetMiddlewares("postTransaction")...)
	router.GET("/transaction", middlewareutils.GetMiddlewares("getTransactions")...)
	router.GET("/transaction/:id", middlewareutils.GetMiddlewares("getTransaction")...)

	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err == nil {
		router.Run(":" + strconv.Itoa(port))
	} else {
		router.Run()
	}
}
