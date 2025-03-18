package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/yangirxd/store-app/basket/docs"
	"github.com/yangirxd/store-app/basket/middleware"
	"github.com/yangirxd/store-app/basket/service"
	"net/http"
)

// @title Basket API
// @version 1.0
// @description This is a basket service using DDD and Gin with JWT authentication
// @host localhost:8083
// @BasePath /
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

type CreateBasketRequest struct {
	UserEmail string `json:"userEmail" binding:"required,email"`
}

type AddItemRequest struct {
	ProductID uuid.UUID `json:"productID" binding:"required"`
	Quantity  int       `json:"quantity" binding:"required,gt=0"`
}

type UpdateItemRequest struct {
	Quantity int `json:"quantity" binding:"required,gt=0"`
}

// @Summary Create a new basket
// @Description Create a new basket for a user (requires authentication)
// @Tags baskets
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param input body CreateBasketRequest true "Basket data"
// @Success 201 {object} domain.Basket
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/baskets [post]
func createBasketHandler(basketService *service.BasketService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateBasketRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		basket, err := basketService.CreateBasket(req.UserEmail)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, basket)
	}
}

// @Summary Get basket
// @Description Get the basket for the authenticated user (requires authentication)
// @Tags baskets
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} domain.Basket
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Basket not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/baskets [get]
func getBasketHandler(basketService *service.BasketService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userEmail := c.GetString("userEmail")
		if userEmail == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user email not found in token"})
			return
		}
		basket, err := basketService.GetBasket(userEmail)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "basket not found"})
			return
		}
		c.JSON(http.StatusOK, basket)
	}
}

// @Summary Add item to basket
// @Description Add an item to the basket (requires authentication)
// @Tags baskets
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param input body AddItemRequest true "Item data"
// @Success 200 {string} string "Item added"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/baskets/items [post]
func addItemHandler(basketService *service.BasketService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userEmail := c.GetString("userEmail")
		if userEmail == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user email not found in token"})
			return
		}
		var req AddItemRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := basketService.AddItem(userEmail, req.ProductID, req.Quantity); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "item added"})
	}
}

// @Summary Remove item from basket
// @Description Remove an item from the basket (requires authentication)
// @Tags baskets
// @Param Authorization header string true "Bearer token"
// @Param itemID path string true "Item ID"
// @Success 200 {string} string "Item removed"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Item not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/baskets/items/{itemID} [delete]
func removeItemHandler(basketService *service.BasketService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userEmail := c.GetString("userEmail")
		if userEmail == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user email not found in token"})
			return
		}
		itemID, err := uuid.Parse(c.Param("itemID"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid item ID"})
			return
		}
		if err := basketService.RemoveItem(userEmail, itemID); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "item removed"})
	}
}

// @Summary Update item quantity
// @Description Update the quantity of an item in the basket (requires authentication)
// @Tags baskets
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param itemID path string true "Item ID"
// @Param input body UpdateItemRequest true "Updated quantity"
// @Success 200 {string} string "Item updated"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Item not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/baskets/items/{itemID} [put]
func updateItemHandler(basketService *service.BasketService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userEmail := c.GetString("userEmail")
		if userEmail == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user email not found in token"})
			return
		}
		itemID, err := uuid.Parse(c.Param("itemID"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid item ID"})
			return
		}
		var req UpdateItemRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := basketService.UpdateItem(userEmail, itemID, req.Quantity); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "item updated"})
	}
}

// @Summary Clear basket
// @Description Clear all items from the basket (requires authentication)
// @Tags baskets
// @Param Authorization header string true "Bearer token"
// @Success 200 {string} string "Basket cleared"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Basket not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/baskets [delete]
func clearBasketHandler(basketService *service.BasketService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userEmail := c.GetString("userEmail")
		if userEmail == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user email not found in token"})
			return
		}
		if err := basketService.ClearBasket(userEmail); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "basket not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "basket cleared"})
	}
}
