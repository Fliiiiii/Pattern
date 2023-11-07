package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Success успешный ответ запроса без данных
func Success(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"error":   "",
	})
	ctx.Abort()
}

// SuccessData успешный ответ запроса всего с 1 полем ответа
func SuccessData(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"error":   "",
		"data":    data,
	})
	ctx.Abort()
}

// SuccessMultipleData успешный ответ запроса с неограниченным числом дополнительный полей ответа
func SuccessMultipleData(ctx *gin.Context, data map[string]any) {
	data["success"] = true
	data["error"] = ""
	ctx.JSON(http.StatusOK, data)
	ctx.Abort()
}
