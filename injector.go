//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	"github.com/Mpayy/toko-api/app"
	"github.com/Mpayy/toko-api/controller"
	"github.com/Mpayy/toko-api/middleware"
	"github.com/Mpayy/toko-api/repository"
	"github.com/Mpayy/toko-api/service"
	// "github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	service.NewCategoryService,
	controller.NewCategoryController,
)

var productSet = wire.NewSet(
	repository.NewProductRepository,
	service.NewProductService,
	controller.NewProductsController,
)

// func ProvideValidator() *validator.Validate {
// 	return validator.New()
// }

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDb,
		app.NewLogger,
		app.NewValidator,
		categorySet,
		productSet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewMiddleware,
		NewServer,
	)
	return nil
}
