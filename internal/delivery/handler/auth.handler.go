package handler

import (
	"net/http"

	"godp/internal/domain"
	"godp/internal/usecase"
	"godp/pkg/response"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct{ authUsecase usecase.AuthUsecase }

func NewAuthHandler(authUsecase usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUsecase: authUsecase}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req domain.User
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.authUsecase.Register(&req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.Created(c, "register success", user)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	token, user, err := h.authUsecase.Login(req.Email, req.Password)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, err.Error())
		return
	}
	response.OK(c, "login success", gin.H{"token": token, "user": user})
}

func (h *AuthHandler) Profile(c *gin.Context) {
	id := c.GetUint("user_id")
	user, err := h.authUsecase.Profile(id)
	if err != nil {
		response.Error(c, http.StatusNotFound, "user not found")
		return
	}
	response.OK(c, "profile", user)
}
