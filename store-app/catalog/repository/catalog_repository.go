package repository

import (
	"github.com/google/uuid"
	"github.com/yangirxd/store-app/catalog/domain"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *domain.Product) error
	FindByID(id uuid.UUID) (*domain.Product, error)
	FindAll() ([]*domain.Product, error)
	Update(product *domain.Product) error
	Delete(id uuid.UUID) error
}

type PostgresProductRepository struct {
	db *gorm.DB
}

func NewPostgresProductRepository(db *gorm.DB) *PostgresProductRepository {
	return &PostgresProductRepository{db: db}
}

func (r *PostgresProductRepository) Create(product *domain.Product) error {
	return r.db.Create(product).Error
}

func (r *PostgresProductRepository) FindByID(id uuid.UUID) (*domain.Product, error) {
	var product domain.Product
	if err := r.db.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *PostgresProductRepository) FindAll() ([]*domain.Product, error) {
	var products []*domain.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *PostgresProductRepository) Update(product *domain.Product) error {
	return r.db.Save(product).Error
}

func (r *PostgresProductRepository) Delete(id uuid.UUID) error {
	return r.db.Where("id = ?", id).Delete(&domain.Product{}).Error
}
