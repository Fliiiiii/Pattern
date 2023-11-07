package router

import (
	"github.com/gin-gonic/gin"
	"reforce.pattern/config"
	"reforce.pattern/pkg/logger"
	"reforce.pattern/pkg/mongodb"
)

type router struct {
	engine *gin.Engine
	mdb    *mongodb.Collections
}

// Init инициализация роутера со всеми необходимыми подключениями к базам данных
func Init(server *gin.Engine, mdb *mongodb.Collections) *router {
	return &router{
		engine: server,
		mdb:    mdb,
	}
}

// Start запуск рабочего роутера
func (r *router) Start() {
	r.Middlewares()
	r.Routes()

	if err := r.engine.Run(config.CFG.Server.Port); err != nil {
		logger.Panic("Ошибка при запуске http сервера по адресу %s: %s", config.CFG.Server.Port, err.Error())
	}
}
