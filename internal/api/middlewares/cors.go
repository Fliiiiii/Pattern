package middlewares

import (
	"github.com/gin-gonic/gin"
)

// Cors настройка заголовков доступа
func Cors(ctx *gin.Context) {
	//ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers",
		"Content-Type, Content-Length, Accept-Encoding,X-CSRF-Token,Authorization,accept,origin,Cache-Control,X-Requested-With")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS,GET")
	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(200)
		return
	}
}
