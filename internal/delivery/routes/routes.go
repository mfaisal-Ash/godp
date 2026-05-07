package routes

import (
	"godp/internal/delivery/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(locationHandler *handler.LocationHandler) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")

	location := api.Group("/locations")
	{
		location.POST("/", locationHandler.Create)
		location.POST("/nearby", locationHandler.FindNearby)
	}

	return r
}
