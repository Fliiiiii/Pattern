package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// BadRequest ошибка на сервере во время обработки запроса
func BadRequest(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"error":   err.Error(),
	})
}

// Unauthorized ответ при ошибке авторизации(не существует пользователь/нет токена/ошибка во время обработки)
func Unauthorized(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"success": false,
		"error":   err.Error(),
	})
}

// Forbidden у пользователя недостаточно прав для использования выбранной api
func Forbidden(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusForbidden, gin.H{
		"success": false,
		"error":   err.Error(),
	})
}

// NotAcceptable данные передаваемые пользователем невозможно принять
func NotAcceptable(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusNotAcceptable, gin.H{
		"success": false,
		"error":   err.Error(),
	})
}
