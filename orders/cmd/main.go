package main

import (
	"context"
	"github.com/google/uuid"
	"github.com/yangirxd/store-app/orders/api"
	"github.com/yangirxd/store-app/orders/db"
	"github.com/yangirxd/store-app/orders/kafka"
	"github.com/yangirxd/store-app/orders/repository"
	"github.com/yangirxd/store-app/orders/service"
	"log"
)

type MockCatalogService struct{}

func (m *MockCatalogService) GetProductPrice(productID uuid.UUID) (float64, error) {
	// Мок для тестов, в реальном проекте нужно делать HTTP-запрос к catalog
	return 10.0, nil
}

func main() {
	ordersDB, err := db.InitOrdersDB()
	if err != nil {
		log.Fatal("failed to initialize database: ", err)
	}

	// Настройка Kafka
	brokers := []string{"kafka:9099"}
	kafkaProducer := kafka.NewProducer(brokers)
	defer kafkaProducer.Close()

	kafkaConsumer := kafka.NewConsumer(brokers, "orders.created", "orders-group")
	//defer kafkaConsumer.Close()

	orderRepo := repository.NewPostgresOrderRepository(ordersDB)
	catalogService := &MockCatalogService{}
	orderService := service.NewOrderService(orderRepo, kafkaProducer, catalogService)

	go func() {
		kafkaConsumer.Consume(context.Background(), orderService.ProcessOrderCreatedEvent)
	}()

	r := api.SetupRouter(orderService)

	if err := r.Run(":8084"); err != nil {
		log.Fatal("failed to start server:", err)
	}
}
