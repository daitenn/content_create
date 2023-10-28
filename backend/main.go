package main

import (
	"go-restapi/controller"
	"go-restapi/db"
	"go-restapi/repository"
	"go-restapi/router"
	"go-restapi/usecase"
	"go-restapi/validator"
)

func main() {
	db := db.NewDB()
	userValidator := validator.NewUserValidator()
	contentValidator := validator.NewContentValidate()
	userRepository := repository.NewUserRepository(db)
	contentRepository := repository.NewContentRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	contentUsecase := usecase.NewContentUsecase(contentRepository, contentValidator)
	userController := controller.NewUserController(userUsecase)
	contentController := controller.NewContentController(contentUsecase)
	e := router.NewRouter(userController, contentController)
	e.Logger.Fatal(e.Start(":8080"))
}
