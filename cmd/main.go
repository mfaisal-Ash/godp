package main

import (
	"log"

	"godp/config"
	"godp/internal/delivery/handler"
	"godp/internal/delivery/routes"
	"godp/internal/repository"
	"godp/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found, using environment/default config")
	}
	config.ConnectDatabase()

	userRepo := repository.NewUserRepository(config.DB)
	locationRepo := repository.NewLocationRepository(config.DB)
	favoriteRepo := repository.NewFavoriteRepository(config.DB)
	storeRepo := repository.NewStoreRepository(config.DB)

	authUsecase := usecase.NewAuthUsecase(userRepo)
	locationUsecase := usecase.NewLocationUsecase(locationRepo)
	favoriteUsecase := usecase.NewFavoriteUsecase(favoriteRepo)
	storeUsecase := usecase.NewStoreUsecase(storeRepo)

	r := gin.Default()
	routes.SetupRoutes(r, routes.Handlers{
		Auth:     handler.NewAuthHandler(authUsecase),
		Location: handler.NewLocationHandler(locationUsecase),
		Favorite: handler.NewFavoriteHandler(favoriteUsecase),
		Store:    handler.NewStoreHandler(storeUsecase),
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
