package db

import (
	"github.com/yangirxd/store-app/basket/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitBasketDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=basket_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error; err != nil {
		log.Fatal("failed to create uuid-ossp extension:", err)
	}

	if err := db.AutoMigrate(&domain.Basket{}, &domain.BasketItem{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	return db, nil
}
