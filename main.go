package main

import (
	"net/http"

	"github.com/Mpayy/toko-api/helper"
	_ "github.com/go-sql-driver/mysql"

	"github.com/Mpayy/toko-api/middleware"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:8080",
		Handler: authMiddleware,
	}
}

func main() {
	//db := app.NewDb()
	//validate := validator.New()
	//productsRepository := repository.NewProductRepository()
	//productService := service.NewProductService(productsRepository, db, validate)
	//productsController := controller.NewProductsController(productService)
	//router := app.NewRouter(productsController)
	//authMiddleware := middleware.NewMiddleware(router)

	//server := http.Server{
	//	Addr:    "localhost:8080",
	//	Handler: authMiddleware,
	//}

	server := InitializedServer()

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
