package main

import (
	"shift-app-go-api/controller"
	"shift-app-go-api/db"
	"shift-app-go-api/repository"
	"shift-app-go-api/router"
	"shift-app-go-api/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUseCase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
