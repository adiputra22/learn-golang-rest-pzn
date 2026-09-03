package main

import (
	"adiputra22/learn-golang-restapi-pzn/app"
	"adiputra22/learn-golang-restapi-pzn/controller"
	"adiputra22/learn-golang-restapi-pzn/helper"
	"adiputra22/learn-golang-restapi-pzn/middleware"
	"adiputra22/learn-golang-restapi-pzn/repository"
	"adiputra22/learn-golang-restapi-pzn/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
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
