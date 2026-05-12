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

type StoreHandler struct{ storeUsecase usecase.StoreUsecase }

func NewStoreHandler(storeUsecase usecase.StoreUsecase) *StoreHandler {
	return &StoreHandler{storeUsecase: storeUsecase}
}

func (h *StoreHandler) Create(c *gin.Context) {
	var req domain.Store
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.storeUsecase.Create(&req); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Created(c, "store created", req)
}

func (h *StoreHandler) AttachOutfit(c *gin.Context) {
	var req struct {
		StoreID  uint `json:"store_id"`
		OutfitID uint `json:"outfit_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if req.StoreID == 0 || req.OutfitID == 0 {
		response.Error(c, http.StatusBadRequest, "store_id and outfit_id are required")
		return
	}
	if err := h.storeUsecase.AttachOutfit(req.StoreID, req.OutfitID); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Created(c, "store outfit link saved", req)
}

func (h *StoreHandler) Nearby(c *gin.Context) {
	longitude, _ := strconv.ParseFloat(c.Query("longitude"), 64)
	latitude, _ := strconv.ParseFloat(c.Query("latitude"), 64)
	radius, _ := strconv.Atoi(c.DefaultQuery("radius_meter", "1000"))
	outfitID64, _ := strconv.ParseUint(c.DefaultQuery("outfit_id", "0"), 10, 64)
	data, err := h.storeUsecase.FindNearby(longitude, latitude, radius, uint(outfitID64), middleware.IsMember(c))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.OK(c, "nearby stores", data)
}
