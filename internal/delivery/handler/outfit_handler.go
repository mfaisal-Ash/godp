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

type OutfitHandler struct{ outfitUsecase usecase.OutfitUsecase }

func NewOutfitHandler(outfitUsecase usecase.OutfitUsecase) *OutfitHandler {
	return &OutfitHandler{outfitUsecase: outfitUsecase}
}

func (h *OutfitHandler) Create(c *gin.Context) {
	var req domain.Outfit
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.outfitUsecase.Create(&req); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Created(c, "outfit created", req)
}

func (h *OutfitHandler) AttachLocation(c *gin.Context) {
	var req struct {
		OutfitID   uint `json:"outfit_id"`
		LocationID uint `json:"location_id"`
		Score      int  `json:"score"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if req.OutfitID == 0 || req.LocationID == 0 {
		response.Error(c, http.StatusBadRequest, "outfit_id and location_id are required")
		return
	}
	if err := h.outfitUsecase.AttachLocation(req.OutfitID, req.LocationID, req.Score); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Created(c, "outfit location match saved", req)
}

func (h *OutfitHandler) FindAll(c *gin.Context) {
	data, err := h.outfitUsecase.FindAll(c.Query("search"), c.Query("category"), c.Query("style"), c.Query("gender"), middleware.IsMember(c))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.OK(c, "outfits", data)
}

func (h *OutfitHandler) PopularGuest(c *gin.Context) {
	data, err := h.outfitUsecase.FindPopularGuest()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.OK(c, "popular outfits for guest", data)
}

func (h *OutfitHandler) Categories(c *gin.Context) {
	data, err := h.outfitUsecase.FindCategories()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.OK(c, "outfit categories", data)
}

func (h *OutfitHandler) MatchByPlace(c *gin.Context) {
	longitude, err := strconv.ParseFloat(c.Query("longitude"), 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "longitude is required and must be number")
		return
	}
	latitude, err := strconv.ParseFloat(c.Query("latitude"), 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "latitude is required and must be number")
		return
	}
	radius, _ := strconv.Atoi(c.DefaultQuery("radius_meter", "1000"))
	data, err := h.outfitUsecase.MatchByPlace(
		longitude,
		latitude,
		radius,
		c.Query("category"),
		c.Query("style"),
		c.Query("gender"),
		middleware.IsMember(c),
	)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.OK(c, "matched outfits by place", data)
}

func (h *OutfitHandler) Detail(c *gin.Context) {
	id64, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	data, err := h.outfitUsecase.FindByID(uint(id64))
	if err != nil {
		response.Error(c, http.StatusNotFound, "outfit not found")
		return
	}
	response.OK(c, "outfit detail", data)
}

func (h *OutfitHandler) Recommendation(c *gin.Context) {
	locationID64, _ := strconv.ParseUint(c.Query("location_id"), 10, 64)
	data, err := h.outfitUsecase.Recommend(uint(locationID64), middleware.IsMember(c))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.OK(c, "outfit recommendations", data)
}
