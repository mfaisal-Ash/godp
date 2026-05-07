package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OK(c *gin.Context, message string, data any) {
	c.JSON(http.StatusOK, gin.H{"message": message, "data": data})
}

func Created(c *gin.Context, message string, data any) {
	c.JSON(http.StatusCreated, gin.H{"message": message, "data": data})
}

func Error(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}
