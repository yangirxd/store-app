package db

import (
	"github.com/yangirxd/store-app/auth/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitAuthDB() (*gorm.DB, error) {
	// Настройка подключения к PostgreSQL
	dsn := "host=localhost user=postgres password=postgres dbname=auth_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error; err != nil {
		log.Fatal("failed to create uuid-ossp extension:", err)
	}

	// Миграция схемы
	if err := db.AutoMigrate(&domain.User{}); err != nil {
		log.Fatal("failed to auto migrate user:", err)
	}

	return db, nil
}
