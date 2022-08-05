package main

import (
	"fmt"
	"rest_api_golang_crud_sqlx/config"
	"rest_api_golang_crud_sqlx/database"
	"rest_api_golang_crud_sqlx/handler"
	"rest_api_golang_crud_sqlx/repository"
	"rest_api_golang_crud_sqlx/service"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Не удается загрузить config:", err)
	}

	db := new(database.DB).InitDatabase(&config)
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	router := handlers.SetupRouter()

	fmt.Printf("Сервер, работающий на порту:%v\n", config.Port)
	router.Run()
}
