package dto

import "github.com/google/uuid"

type CreateOrderRequest struct {
	UserEmail string      `json:"userEmail" binding:"required,email"`
	Items     []OrderItem `json:"items" binding:"required"`
}

type OrderItem struct {
	ProductID uuid.UUID `json:"productID" binding:"required"`
	Quantity  int       `json:"quantity" binding:"required,gt=0"`
}
