package domain

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Order struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserEmail string    `gorm:"not null"`
	Total     float64   `gorm:"not null;default:0.0"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	Items     []OrderItem
}

type OrderItem struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	OrderID   uuid.UUID `gorm:"not null"`
	ProductID uuid.UUID `gorm:"not null"`
	Quantity  int       `gorm:"not null;default:1"`
	Price     float64   `gorm:"not null;default:0.0"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
}

// NewOrder создает новый заказ
func NewOrder(userEmail string) *Order {
	return &Order{
		ID:        uuid.New(),
		UserEmail: userEmail,
		CreatedAt: time.Now(),
	}
}

// NewOrderItem создает новый элемент заказа
func NewOrderItem(orderID, productID uuid.UUID, quantity int, price float64) (*OrderItem, error) {
	if quantity <= 0 {
		return nil, fmt.Errorf("quantity must be positive")
	}
	if price < 0 {
		return nil, fmt.Errorf("price must be non-negative")
	}
	return &OrderItem{
		ID:        uuid.New(),
		OrderID:   orderID,
		ProductID: productID,
		Quantity:  quantity,
		Price:     price,
		CreatedAt: time.Now(),
	}, nil
}

// AddItem добавляет элемент в заказ и обновляет общую сумму
func (o *Order) AddItem(item *OrderItem) {
	o.Items = append(o.Items, *item)
	o.Total += float64(item.Quantity) * item.Price
}
