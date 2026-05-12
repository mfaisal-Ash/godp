package handler

import (
	"net/http"
	"strconv"

	"godp/internal/delivery/middleware"
	"godp/internal/domain"
	"godp/internal/usecase"
	"godp/pkg/response"

	"github.com/gin-gonic/gin"
)

type LocationHandler struct{ locationUsecase usecase.LocationUsecase }

func NewLocationHandler(locationUsecase usecase.LocationUsecase) *LocationHandler {
	return &LocationHandler{locationUsecase: locationUsecase}
}

func (h *LocationHandler) Create(c *gin.Context) {
	var req domain.Location
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.locationUsecase.Create(&req); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Created(c, "location created", req)
}

func (h *LocationHandler) FindAll(c *gin.Context) {
	var req domain.Location
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	data, err := h.locationUsecase.FindAll(middleware.IsMember(c))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.OK(c, "locations", data)

	if err := h.locationUsecase.Create(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "location created successfully",
		"data":    req,
	})
}

func (h *LocationHandler) FindNearby(c *gin.Context) {
	longitude, _ := strconv.ParseFloat(c.Query("longitude"), 64)
	latitude, _ := strconv.ParseFloat(c.Query("latitude"), 64)
	radius, _ := strconv.Atoi(c.DefaultQuery("radius_meter", "1000"))
	data, err := h.locationUsecase.FindNearby(longitude, latitude, radius, middleware.IsMember(c))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.OK(c, "nearby locations", data)
}
