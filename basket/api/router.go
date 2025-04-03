package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/yangirxd/store-app/basket/docs"
	"github.com/yangirxd/store-app/basket/middleware"
	"github.com/yangirxd/store-app/basket/service"
)

// @title Basket API
// @version 1.0
// @description This is a basket service using DDD and Gin with JWT authentication
// @host localhost:8083
// @BasePath /basket
func SetupRouter(basketService *service.BasketService) *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api/v1")
	{
		protected := api.Group("", middleware.BasketMiddleware())
		{
			protected.POST("/baskets", createBasketHandler(basketService))
			protected.GET("/baskets", getBasketHandler(basketService))
			protected.POST("/baskets/items", addItemHandler(basketService))
			protected.DELETE("/baskets/items/:itemID", removeItemHandler(basketService))
			protected.PUT("/baskets/items/:itemID", updateItemHandler(basketService))
			protected.DELETE("/baskets", clearBasketHandler(basketService))
		}
	}

	return r
}
