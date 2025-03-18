package repository

import (
	"github.com/google/uuid"
	"github.com/yangirxd/store-app/basket/domain"
	"gorm.io/gorm"
)

type BasketRepository interface {
	CreateBasket(basket *domain.Basket) error
	GetBasketByUserEmail(userEmail string) (*domain.Basket, error)
	AddItem(basketItem *domain.BasketItem) error
	RemoveItem(basketID, itemID uuid.UUID) error
	UpdateItem(basketItem *domain.BasketItem) error
	ClearBasket(basketID uuid.UUID) error
	FindItemByID(basketID, itemID uuid.UUID) (*domain.BasketItem, error)
}

type PostgresBasketRepository struct {
	db *gorm.DB
}

func NewPostgresBasketRepository(db *gorm.DB) *PostgresBasketRepository {
	return &PostgresBasketRepository{db: db}
}

func (r *PostgresBasketRepository) CreateBasket(basket *domain.Basket) error {
	return r.db.Create(basket).Error
}

func (r *PostgresBasketRepository) GetBasketByUserEmail(userEmail string) (*domain.Basket, error) {
	var basket domain.Basket
	if err := r.db.Preload("Items").Where("user_email = ?", userEmail).First(&basket).Error; err != nil {
		return nil, err
	}
	return &basket, nil
}

func (r *PostgresBasketRepository) AddItem(basketItem *domain.BasketItem) error {
	return r.db.Create(basketItem).Error
}

func (r *PostgresBasketRepository) RemoveItem(basketID, itemID uuid.UUID) error {
	return r.db.Where("basket_id = ? AND id = ?", basketID, itemID).Delete(&domain.BasketItem{}).Error
}

func (r *PostgresBasketRepository) UpdateItem(basketItem *domain.BasketItem) error {
	return r.db.Save(basketItem).Error
}

func (r *PostgresBasketRepository) ClearBasket(basketID uuid.UUID) error {
	if err := r.db.Where("basket_id = ?", basketID).Delete(&domain.BasketItem{}).Error; err != nil {
		return err
	}
	return r.db.Delete(&domain.Basket{ID: basketID}).Error
}
func (r *PostgresBasketRepository) FindItemByID(basketID, itemID uuid.UUID) (*domain.BasketItem, error) {
	var item domain.BasketItem
	if err := r.db.Where("basket_id = ? AND id = ?", basketID, itemID).First(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}
