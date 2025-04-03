package main

import (
	"github.com/yangirxd/store-app/catalog/api"
	"github.com/yangirxd/store-app/catalog/db"
	"github.com/yangirxd/store-app/catalog/repository"
	"github.com/yangirxd/store-app/catalog/service"
	"log"
)

func main() {
	catalogDB, err := db.InitCatalogDB()
	if err != nil {
		log.Fatal("failed to initialize database: ", err)
	}

	productRepo := repository.NewPostgresProductRepository(catalogDB)
	catalogService := service.NewCatalogService(productRepo)

	r := api.SetupRouter(catalogService)

	if err := r.Run(":8081"); err != nil {
		log.Fatal("failed to start server:", err)
	}
}
