package db

import (
	"github.com/yangirxd/store-app/catalog/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitCatalogDB() (*gorm.DB, error) {
	// Настройка подключения к PostgreSQL
	dsn := "host=localhost user=postgres password=postgres dbname=catalog_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error; err != nil {
		log.Fatal("failed to create uuid-ossp extension:", err)
	}

	if err := db.AutoMigrate(&domain.Product{}); err != nil {
		log.Fatal("failed to auto migrate user:", err)
	}

	return db, nil
}
