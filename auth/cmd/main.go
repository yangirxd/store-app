package main

import (
	"github.com/yangirxd/store-app/auth/api"
	"github.com/yangirxd/store-app/auth/db"
	"github.com/yangirxd/store-app/auth/repository"
	"github.com/yangirxd/store-app/auth/service"
	"log"
)

func main() {
	authDB, err := db.InitAuthDB()
	if err != nil {
		log.Fatal("failed to initialize database: ", err)
	}

	// Инициализация зависимостей
	userRepo := repository.NewPostgresUserRepository(authDB)
	authService := service.NewAuthService(userRepo)

	// Настройка роутера
	r := api.SetupRouter(authService)

	// Запуск сервера
	if err := r.Run(":8085"); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
