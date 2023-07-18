package main

import (
	"shift-app-go-api/controller"
	"shift-app-go-api/db"
	"shift-app-go-api/repository"
	"shift-app-go-api/router"
	"shift-app-go-api/usecase"
	"shift-app-go-api/validator"
)

func main() {
	db := db.NewDB()

	userValidator := validator.NewUserValidator()
	shiftValidator := validator.NewShiftValidator()

	userRepository := repository.NewUserRepository(db)
	shiftRepository := repository.NewShiftRepository(db)

	userUsecase := usecase.NewUserUseCase(userRepository, userValidator)
	shiftUsecase := usecase.NewShiftUsecase(shiftRepository, shiftValidator)

	userController := controller.NewUserController(userUsecase)
	shiftController := controller.NewShiftController(shiftUsecase)

	e := router.NewRouter(userController, shiftController)
	e.Logger.Fatal(e.Start(":8080"))
}
