package db

import (
	"github.com/yangirxd/store-app/auth/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func InitAuthDB() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error; err != nil {
		log.Fatal("failed to create uuid-ossp extension:", err)
	}

	if err := db.AutoMigrate(&domain.User{}); err != nil {
		log.Fatal("failed to auto migrate user:", err)
	}

	return db, nil
}
