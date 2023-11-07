package http

import (
	"github.com/gin-gonic/gin"
	"reforce.pattern/config"
)

var cfg = config.CFG.Server

func Init() *gin.Engine {
	if cfg.Development == false {
		gin.SetMode(gin.ReleaseMode)
	}

	return gin.Default()
}
