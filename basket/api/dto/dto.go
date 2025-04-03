package dto

import "github.com/google/uuid"

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
