package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/yangirxd/store-app/catalog/docs"
	"github.com/yangirxd/store-app/catalog/middleware"
	"github.com/yangirxd/store-app/catalog/service"
)

// @title Catalog API
// @version 1.0
// @description This is a catalog service using DDD and Gin with JWT authentication
// @host localhost:8081
// @BasePath /catalog
func SetupRouter(catalogService *service.CatalogService) *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api/v1")
	{
		api.GET("/products", getAllProductsHandler(catalogService))
		api.GET("/products/:id", getProductHandler(catalogService))

		protected := api.Group("", middleware.CatalogMiddleware())
		{
			protected.POST("/products", createProductHandler(catalogService))
			protected.PUT("/products/:id", updateProductHandler(catalogService))
			protected.DELETE("/products/:id", deleteProductHandler(catalogService))
		}
	}

	return r
}
