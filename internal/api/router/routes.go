package router

import "reforce.pattern/internal/api/handlers"

// Routes инициализация всех путей(api) сервера
func (r *router) Routes() {
	// группировка апи с определением базового пути
	group := r.engine.Group("")

	// пинг
	group.GET("ping", handlers.Ping)
}
