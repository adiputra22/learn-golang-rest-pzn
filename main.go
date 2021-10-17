package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"learn-golang-restapi-pzn/app"
	"learn-golang-restapi-pzn/controller"
	"learn-golang-restapi-pzn/helper"
	"learn-golang-restapi-pzn/middleware"
	"learn-golang-restapi-pzn/repository"
	"learn-golang-restapi-pzn/service"
	"net/http"
)

func main() {
	db := app.NewDB()

	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
