package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Mpayy/toko-api/app"
	"github.com/Mpayy/toko-api/controller"
	"github.com/Mpayy/toko-api/helper"
	"github.com/Mpayy/toko-api/middleware"
	"github.com/Mpayy/toko-api/repository"
	"github.com/Mpayy/toko-api/service"
	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDb()
	validate := validator.New()
	productsRepository := repository.NewProductRepository()
	productService := service.NewProductService(productsRepository, db, validate)
	productsController := controller.NewProductsController(productService)

	router := app.NewRouter(productsController)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: middleware.NewProductMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
