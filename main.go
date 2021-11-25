package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/komblog/restful-api-go/app"
	"github.com/komblog/restful-api-go/controller"
	"github.com/komblog/restful-api-go/helper"
	"github.com/komblog/restful-api-go/repository"
	"github.com/komblog/restful-api-go/service"
)

func main() {
	db := app.Connection()
	validator := validator.New()

	repository := repository.NewUserRepository()
	service := service.NewUserService(repository, db, validator)
	controller := controller.NewUserController(service)

	router := httprouter.New()

	router.GET("/api/user", controller.FindAll)
	router.GET("/api/user/:userId", controller.FindById)
	router.POST("/api/user", controller.Create)
	router.PUT("/api/user/:userId", controller.Update)
	router.DELETE("/api/user/:userId", controller.Delete)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	errServer := server.ListenAndServe()
	helper.PanicIfError(errServer)

}
