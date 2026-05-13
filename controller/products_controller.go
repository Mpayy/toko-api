package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ProductsController interface {
	CreateProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetAllProducts(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetProductByID(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
