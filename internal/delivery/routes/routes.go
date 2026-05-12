package routes

import (
	"godp/internal/delivery/handler"
	"godp/internal/delivery/middleware"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Auth     *handler.AuthHandler
	Location *handler.LocationHandler
	Outfit   *handler.OutfitHandler
	Favorite *handler.FavoriteHandler
	Store    *handler.StoreHandler
}

func SetupRoutes(r *gin.Engine, h Handlers) {
	api := r.Group("/api")

	auth := api.Group("/auth")
	{
		auth.POST("/register", h.Auth.Register)
		auth.POST("/login", h.Auth.Login)
	}

	guest := api.Group("/guest")
	{
		guest.GET("/outfits/popular", h.Outfit.PopularGuest)
	}

	public := api.Group("")
	public.Use(middleware.OptionalAuth())
	{
		public.GET("/outfits", h.Outfit.FindAll)
		public.GET("/outfits/categories", h.Outfit.Categories)
		public.GET("/outfits/recommendations", h.Outfit.Recommendation)
		public.GET("/outfits/match-place", h.Outfit.MatchByPlace)
		public.GET("/outfits/:id", h.Outfit.Detail)
		public.GET("/locations", h.Location.FindAll)
		public.GET("/locations/nearby", h.Location.FindNearby)
		public.GET("/stores/nearby", h.Store.Nearby)
	}

	member := api.Group("/member")
	member.Use(middleware.AuthRequired())
	{
		member.GET("/profile", h.Auth.Profile)
		member.GET("/favorites", h.Favorite.FindByUser)
		member.POST("/favorites", h.Favorite.Add)
		member.DELETE("/favorites/:outfit_id", h.Favorite.Remove)
	}

	admin := api.Group("/admin")
	admin.Use(middleware.AuthRequired())
	{
		admin.POST("/locations", h.Location.Create)
		admin.POST("/outfits", h.Outfit.Create)
		admin.POST("/outfit-locations", h.Outfit.AttachLocation)
		admin.POST("/stores", h.Store.Create)
		admin.POST("/store-outfits", h.Store.AttachOutfit)
	}
}
