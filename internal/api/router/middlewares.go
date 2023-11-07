package router

import (
	"github.com/gin-gonic/gin"
	"reforce.pattern/internal/api/middlewares"
	"reforce.pattern/internal/api/response"
)

// Middlewares добавление cors, а также проверки авторизации пользователя
func (r *router) Middlewares() {
	r.engine.Use(middlewares.Cors)
	r.engine.Use(func(ctx *gin.Context) {
		if err := middlewares.Authorization(ctx, r.mdb); err != nil {
			response.Unauthorized(ctx, err)
			ctx.Abort()
		}
	})
}
