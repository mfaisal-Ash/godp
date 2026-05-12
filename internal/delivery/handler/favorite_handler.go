package handler

import (
	"net/http"
	"strconv"

	"godp/internal/delivery/middleware"
	"godp/internal/usecase"
	"godp/pkg/response"

	"github.com/gin-gonic/gin"
)

type FavoriteHandler struct{ favoriteUsecase usecase.FavoriteUsecase }

func NewFavoriteHandler(favoriteUsecase usecase.FavoriteUsecase) *FavoriteHandler {
	return &FavoriteHandler{favoriteUsecase: favoriteUsecase}
}

func (h *FavoriteHandler) Add(c *gin.Context) {
	var req struct {
		OutfitID uint `json:"outfit_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.favoriteUsecase.Add(middleware.UserID(c), req.OutfitID); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Created(c, "favorite added", req)
}

func (h *FavoriteHandler) FindByUser(c *gin.Context) {
	data, err := h.favoriteUsecase.FindByUser(middleware.UserID(c))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.OK(c, "favorites", data)
}

func (h *FavoriteHandler) Remove(c *gin.Context) {
	id64, _ := strconv.ParseUint(c.Param("outfit_id"), 10, 64)
	if err := h.favoriteUsecase.Remove(middleware.UserID(c), uint(id64)); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.OK(c, "favorite removed", nil)
}
