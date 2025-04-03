package domain

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Basket struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserEmail string    `gorm:"not null;unique"` // Связь с пользователем через email
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	Items     []BasketItem
}

type BasketItem struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	BasketID  uuid.UUID `gorm:"not null"`
	ProductID uuid.UUID `gorm:"not null"` // Ссылка на продукт из catalog
	Quantity  int       `gorm:"not null;default:1"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
}

func NewBasket(userEmail string) *Basket {
	return &Basket{
		ID:        uuid.New(),
		UserEmail: userEmail,
		CreatedAt: time.Now(),
	}
}

// NewBasketItem создает новый элемент корзины
func NewBasketItem(basketID, productID uuid.UUID, quantity int) (*BasketItem, error) {
	if quantity <= 0 {
		return nil, fmt.Errorf("quantity must be positive")
	}

	return &BasketItem{
		ID:        uuid.New(),
		BasketID:  basketID,
		ProductID: productID,
		Quantity:  quantity,
		CreatedAt: time.Now(),
	}, nil
}
