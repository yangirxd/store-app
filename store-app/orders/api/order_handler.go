package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yangirxd/store-app/orders/api/dto"
	_ "github.com/yangirxd/store-app/orders/docs"
	"github.com/yangirxd/store-app/orders/service"
	"net/http"
)

// @Summary Create a new order
// @Description Create a new order for a user (requires authentication)
// @Tags orders
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param input body dto.CreateOrderRequest true "Order data"
// @Success 201 {object} domain.Order
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/orders [post]
func createOrderHandler(orderService *service.OrderService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userEmail := c.GetString("userEmail")
		if userEmail == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user email not found in token"})
			return
		}
		var req dto.CreateOrderRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if req.UserEmail != userEmail {
			c.JSON(http.StatusForbidden, gin.H{"error": "user email in request does not match authenticated user"})
			return
		}

		// Преобразуем []dto.OrderItem в []service.BasketItem
		items := make([]service.BasketItem, len(req.Items))
		for i, item := range req.Items {
			items[i] = service.BasketItem{
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
			}
		}

		order, err := orderService.CreateOrder(req.UserEmail, items)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, order)
	}
}

// @Summary Get order by ID
// @Description Get an order by its ID (requires authentication)
// @Tags orders
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param orderID path string true "Order ID"
// @Success 200 {object} domain.Order
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Order not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/orders/{orderID} [get]
func getOrderHandler(orderService *service.OrderService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userEmail := c.GetString("userEmail")
		if userEmail == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user email not found in token"})
			return
		}
		orderID, err := uuid.Parse(c.Param("orderID"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order ID"})
			return
		}
		order, err := orderService.GetOrder(orderID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
			return
		}
		if order.UserEmail != userEmail {
			c.JSON(http.StatusForbidden, gin.H{"error": "order does not belong to user"})
			return
		}
		c.JSON(http.StatusOK, order)
	}
}

// @Summary Get all orders for user
// @Description Get all orders for the authenticated user (requires authentication)
// @Tags orders
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {array} domain.Order
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/orders [get]
func getOrdersHandler(orderService *service.OrderService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userEmail := c.GetString("userEmail")
		if userEmail == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user email not found in token"})
			return
		}
		orders, err := orderService.GetOrders(userEmail)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, orders)
	}
}
