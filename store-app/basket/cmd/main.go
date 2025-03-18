package main

import (
	"github.com/yangirxd/store-app/basket/api"
	"github.com/yangirxd/store-app/basket/db"
	"github.com/yangirxd/store-app/basket/repository"
	"github.com/yangirxd/store-app/basket/service"
	"log"
)

func main() {
	basketDB, err := db.InitBasketDB()
	if err != nil {
		log.Fatal("failed to initialize database: ", err)
	}

	basketRepo := repository.NewPostgresBasketRepository(basketDB)
	basketService := service.NewBasketService(basketRepo)

	r := api.SetupRouter(basketService)

	log.Println("Server running on http://localhost:8083")
	if err := r.Run(":8083"); err != nil {
		log.Fatal("failed to start server:", err)
	}
}
