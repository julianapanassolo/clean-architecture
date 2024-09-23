package routers

import (
	"clean-architecture/api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	orderController := controllers.NewOrderController()

	// Endpoint GET /order
	router.GET("/order", orderController.ListOrders)
}
