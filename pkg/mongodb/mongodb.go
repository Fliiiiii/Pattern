package mongodb

import (
	"github.com/maratIbatulin/mongodb/mongo"
	"reforce.pattern/config"
	"reforce.pattern/pkg/logger"
)

var cfg = config.CFG.MongoDB

func Init() *Collections {
	connection := mongo.Connection().Default(cfg.App)
	connection.Hosts(cfg.Hosts).Pools(cfg.PoolLimits.Min, cfg.PoolLimits.Max)
	logger.Info("Создание подключения к mongo, попытка соединения")
	//connection.Auth(connect.Auth{
	//	DB:       "admin",
	//	UserName: cfg.User,
	//	Password: cfg.Password,
	//})

	db, err := mongo.ConnectDB(connection, cfg.DB)
	if err != nil {
		logger.Fatal("Ошибка при подключении к mongo: %s", err.Error())
	}
	logger.Info("Успешное подключение, mongo готова к работе")

	return &Collections{
		ContactPersons: db.Collection("ContactPersons"),
	}
}
