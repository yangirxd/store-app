package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/yangirxd/store-app/orders/domain"
	"github.com/yangirxd/store-app/orders/kafka"
	"github.com/yangirxd/store-app/orders/repository"
)

type OrderService struct {
	orderRepo      repository.OrderRepository
	kafkaProducer  *kafka.Producer
	catalogService CatalogServiceClient
}

type CatalogServiceClient interface {
	GetProductPrice(productID uuid.UUID) (float64, error)
}

type BasketItem struct {
	ProductID uuid.UUID `json:"productID"`
	Quantity  int       `json:"quantity"`
}

func NewOrderService(orderRepo repository.OrderRepository, kafkaProducer *kafka.Producer, catalogService CatalogServiceClient) *OrderService {
	return &OrderService{
		orderRepo:      orderRepo,
		kafkaProducer:  kafkaProducer,
		catalogService: catalogService,
	}
}

func (s *OrderService) CreateOrder(userEmail string, items []BasketItem) (*domain.Order, error) {
	order := domain.NewOrder(userEmail)

	for _, item := range items {
		price, err := s.catalogService.GetProductPrice(item.ProductID)
		if err != nil {
			return nil, fmt.Errorf("failed to get product price: %v", err)
		}
		orderItem, err := domain.NewOrderItem(order.ID, item.ProductID, item.Quantity, price)
		if err != nil {
			return nil, err
		}
		order.AddItem(orderItem)
	}

	if err := s.orderRepo.CreateOrder(order); err != nil {
		return nil, err
	}

	eventData, _ := json.Marshal(order)
	if err := s.kafkaProducer.Produce(context.Background(), "orders.created", eventData); err != nil {
		// Логируем ошибку, но не прерываем выполнение
		fmt.Printf("Failed to produce order created event: %v\n", err)
	}

	return order, nil
}

func (s *OrderService) GetOrder(orderID uuid.UUID) (*domain.Order, error) {
	return s.orderRepo.GetOrderByID(orderID)
}

func (s *OrderService) GetOrders(userEmail string) ([]domain.Order, error) {
	return s.orderRepo.GetOrdersByUserEmail(userEmail)
}

func (s *OrderService) ProcessOrderCreatedEvent(data []byte) error {
	var event domain.Order
	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("failed to unmarshal basket cleared event: %v", err)
	}

	fmt.Printf("Created order for user %s\n", event.UserEmail)
	return nil
}
