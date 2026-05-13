package app

import (
	"github.com/Mpayy/toko-api/controller"
	"github.com/Mpayy/toko-api/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(productController controller.ProductsController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/products", productController.CreateProduct)
	router.GET("/api/products", productController.GetAllProducts)
	router.GET("/api/products/:productId", productController.GetProductByID)
	router.PUT("/api/products/:productId", productController.UpdateProduct)
	router.DELETE("/api/products/:productId", productController.DeleteProduct)

	router.PanicHandler = exception.ErrorHandler

	return router
}
