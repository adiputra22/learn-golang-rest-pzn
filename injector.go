//go:build wireinject
// +build wireinject

package main

import (
	"adiputra22/learn-golang-restapi-pzn/app"
	"adiputra22/learn-golang-restapi-pzn/controller"
	"adiputra22/learn-golang-restapi-pzn/middleware"
	"adiputra22/learn-golang-restapi-pzn/repository"
	"adiputra22/learn-golang-restapi-pzn/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepositoryImpl,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)

	return &http.Server{}
}
