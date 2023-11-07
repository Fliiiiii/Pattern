package main

import (
	"log"
	"reforce.pattern/internal/api/router"
	"reforce.pattern/pkg/http"
	"reforce.pattern/pkg/mongodb"
)

func main() {
	log.Println("Запуск сервиса по работе с объектами партнера в приложении PRM ")

	//создание статического подключения к mongo
	mdb := mongodb.Init()
	//инициализация http роутера
	engine := http.Init()

	//инициализация middlewares и routes для http сервераа
	server := router.Init(engine, mdb)
	server.Start()
	//test
}
