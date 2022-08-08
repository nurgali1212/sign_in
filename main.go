package main

import (
	"fmt"
	"sign_in/config"
	"sign_in/database"
	"sign_in/handler"
	"sign_in/repository"
	"sign_in/service"
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
