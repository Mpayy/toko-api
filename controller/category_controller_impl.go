package controller

import (
	"net/http"
	"strconv"

	"github.com/Mpayy/toko-api/helper"
	"github.com/Mpayy/toko-api/model/web"
	"github.com/Mpayy/toko-api/service"
	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) CreateCategory(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadRequestBody(request, &categoryCreateRequest)

	category := controller.CategoryService.CreateCategory(request.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusCreated,
		Status: "success",
		Data:   category,
	}
	helper.WriteResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) GetAllCategory(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categories := controller.CategoryService.GetAllCategories(request.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "success",
		Data:   categories,
	}
	helper.WriteResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) GetCategoryById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("categoryId")
	categoryId, err := strconv.Atoi(id)
	helper.PanicIfError(err)
	category := controller.CategoryService.GetCategoryById(request.Context(), categoryId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "success",
		Data:   category,
	}
	helper.WriteResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) UpdateCategory(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadRequestBody(request, &categoryUpdateRequest)

	id := params.ByName("categoryId")
	categoryId, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = categoryId

	category := controller.CategoryService.UpdateCategory(request.Context(), categoryUpdateRequest)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "success",
		Data:   category,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) DeleteCategory(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("categoryId")
	categoryId, err := strconv.Atoi(id)
	helper.PanicIfError(err)
	controller.CategoryService.DeleteCategory(request.Context(), categoryId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "success",
	}
	helper.WriteResponseBody(writer, webResponse)
}
