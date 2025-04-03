package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/yangirxd/store-app/orders/middleware"
	"github.com/yangirxd/store-app/orders/service"
)

// @title Orders API
// @version 1.0
// @description This is an orders service using DDD and Gin with JWT authentication and Kafka
// @host localhost:8084
// @BasePath /orders
func SetupRouter(orderService *service.OrderService) *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api/v1")
	{
		protected := api.Group("", middleware.OrderMiddleware())
		{
			protected.POST("/orders", createOrderHandler(orderService))
			protected.GET("/orders/:orderID", getOrderHandler(orderService))
			protected.GET("/orders", getOrdersHandler(orderService))
		}
	}
	return r
}
