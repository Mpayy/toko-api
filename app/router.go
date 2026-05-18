package app

import (
	"github.com/Mpayy/toko-api/controller"
	"github.com/Mpayy/toko-api/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(productController controller.ProductsController, categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/products", productController.CreateProduct)
	router.GET("/api/products", productController.GetAllProducts)
	router.GET("/api/products/:productId", productController.GetProductByID)
	router.PUT("/api/products/:productId", productController.UpdateProduct)
	router.DELETE("/api/products/:productId", productController.DeleteProduct)

	router.POST("/api/categories", categoryController.CreateCategory)
	router.GET("/api/categories", categoryController.GetAllCategory)
	router.GET("/api/categories/:categoryId", categoryController.GetCategoryById)
	router.PUT("/api/categories/:categoryId", categoryController.UpdateCategory)
	router.DELETE("/api/categories/:categoryId", categoryController.DeleteCategory)

	router.PanicHandler = exception.ErrorHandler

	return router
}
