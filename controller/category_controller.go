package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CategoryController interface {
	CreateCategory(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetAllCategory(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetCategoryById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateCategory(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteCategory(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
