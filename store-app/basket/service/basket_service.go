package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/yangirxd/store-app/basket/domain"
	"github.com/yangirxd/store-app/basket/repository"
)

type BasketService struct {
	basketRepo repository.BasketRepository
}

func NewBasketService(basketRepo repository.BasketRepository) *BasketService {
	return &BasketService{
		basketRepo: basketRepo,
	}
}

func (s *BasketService) CreateBasket(userEmail string) (*domain.Basket, error) {
	basket := domain.NewBasket(userEmail)
	if err := s.basketRepo.CreateBasket(basket); err != nil {
		return nil, err
	}
	return basket, nil
}

func (s *BasketService) GetBasket(userEmail string) (*domain.Basket, error) {
	return s.basketRepo.GetBasketByUserEmail(userEmail)
}

func (s *BasketService) AddItem(userEmail string, productID uuid.UUID, quantity int) error {
	basket, err := s.basketRepo.GetBasketByUserEmail(userEmail)
	if err != nil {
		return fmt.Errorf("basket not found: %v", err)
	}

	item, err := domain.NewBasketItem(basket.ID, productID, quantity)
	if err != nil {
		return err
	}
	return s.basketRepo.AddItem(item)
}

func (s *BasketService) RemoveItem(userEmail string, itemID uuid.UUID) error {
	basket, err := s.basketRepo.GetBasketByUserEmail(userEmail)
	if err != nil {
		return fmt.Errorf("basket not found: %v", err)
	}

	return s.basketRepo.RemoveItem(basket.ID, itemID)
}

func (s *BasketService) UpdateItem(userEmail string, itemID uuid.UUID, quantity int) error {
	basket, err := s.basketRepo.GetBasketByUserEmail(userEmail)
	if err != nil {
		return fmt.Errorf("basket not found: %v", err)
	}

	item, err := s.basketRepo.FindItemByID(basket.ID, itemID)
	if err != nil {
		return fmt.Errorf("item not found: %v", err)
	}
	item.Quantity = quantity

	return s.basketRepo.UpdateItem(item)
}

func (s *BasketService) ClearBasket(userEmail string) error {
	basket, err := s.basketRepo.GetBasketByUserEmail(userEmail)
	if err != nil {
		return fmt.Errorf("basket not found: %v", err)
	}

	return s.basketRepo.ClearBasket(basket.ID)
}
