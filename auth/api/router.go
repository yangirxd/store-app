package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/yangirxd/store-app/auth/docs"
	"github.com/yangirxd/store-app/auth/service"
)

// @title Auth API
// @version 1.0
// @description This is an auth service using DDD and Gin
// @host localhost:8085
// @BasePath /auth
func SetupRouter(authService *service.AuthService) *gin.Engine {
	r := gin.Default()

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Группа API
	api := r.Group("/user/v1")
	{
		api.POST("/register", registerHandler(authService))
		api.POST("/login", loginHandler(authService))
	}

	return r
}
