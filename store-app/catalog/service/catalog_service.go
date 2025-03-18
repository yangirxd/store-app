package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/yangirxd/store-app/catalog/domain"
	"github.com/yangirxd/store-app/catalog/repository"
)

type CatalogService struct {
	productRepo repository.ProductRepository
}

func NewCatalogService(productRepo repository.ProductRepository) *CatalogService {
	return &CatalogService{productRepo: productRepo}
}

func (s *CatalogService) CreateProduct(name, description string, price float64, stock int) (*domain.Product, error) {
	product, err := domain.NewProduct(name, description, price, stock)
	if err != nil {
		return nil, err
	}

	if err := s.productRepo.Create(product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *CatalogService) GetProductByID(id string) (*domain.Product, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID: %v", err)
	}

	return s.productRepo.FindByID(uid)
}

func (s *CatalogService) GetAllProducts() ([]*domain.Product, error) {
	return s.productRepo.FindAll()
}

func (s *CatalogService) UpdateProduct(id, name, description string, price float64, stock int) (*domain.Product, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID: %v", err)
	}

	product, err := s.productRepo.FindByID(uid)
	if err != nil {
		return nil, err
	}

	product.Name = name
	product.Description = description
	product.Price = price
	product.Stock = stock
	if err := s.productRepo.Update(product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *CatalogService) DeleteProduct(id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid UUID: %v", err)
	}

	return s.productRepo.Delete(uid)
}
