package repository

import (
	"github.com/google/uuid"
	"github.com/yangirxd/store-app/orders/domain"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order *domain.Order) error
	GetOrderByID(orderID uuid.UUID) (*domain.Order, error)
	GetOrdersByUserEmail(userEmail string) ([]domain.Order, error)
}

type PostgresOrderRepository struct {
	db *gorm.DB
}

func NewPostgresOrderRepository(db *gorm.DB) *PostgresOrderRepository {
	return &PostgresOrderRepository{db: db}
}

func (r *PostgresOrderRepository) CreateOrder(order *domain.Order) error {
	return r.db.Create(order).Error
}

func (r *PostgresOrderRepository) GetOrderByID(orderID uuid.UUID) (*domain.Order, error) {
	var order domain.Order
	if err := r.db.Preload("Items").First(&order, "id = ?", orderID).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *PostgresOrderRepository) GetOrdersByUserEmail(userEmail string) ([]domain.Order, error) {
	var orders []domain.Order
	if err := r.db.Preload("Items").Where("user_email = ?", userEmail).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
