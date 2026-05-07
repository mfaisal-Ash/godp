package main

import (
	"godp/config"
	"godp/internal/delivery/handler"
	"godp/internal/delivery/routes"
	"godp/internal/repository"
	"godp/internal/usecase"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	config.ConnectDatabase()
	locationRepo := repository.NewLocationRepository(config.DB)
	locationUsecase := usecase.NewLocationUsecase(locationRepo)
	locationHandler := handler.NewLocationHandler(locationUsecase)

	router := routes.SetupRouter(locationHandler)
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
