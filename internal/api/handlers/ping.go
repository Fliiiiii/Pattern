package handlers

import (
	"github.com/gin-gonic/gin"
	"reforce.pattern/internal/api/response"
)

// Ping стандартный пинг для проверки доступности сервера
func Ping(ctx *gin.Context) {
	response.Success(ctx)
}
