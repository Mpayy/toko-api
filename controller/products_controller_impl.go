package controller

import (
	"net/http"
	"strconv"

	"github.com/Mpayy/toko-api/helper"
	"github.com/Mpayy/toko-api/model/web"
	"github.com/Mpayy/toko-api/service"
	"github.com/julienschmidt/httprouter"
)

type ProductsControllerImpl struct {
	ProductsService service.ProductService
}

func NewProductsController(productsService service.ProductService) ProductsController {
	return &ProductsControllerImpl{
		ProductsService: productsService,
	}
}

func (controller *ProductsControllerImpl) CreateProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productCreateRequest := web.ProductCreateRequest{}
	helper.ReadRequestBody(request, &productCreateRequest)

	product := controller.ProductsService.CreateProduct(request.Context(), productCreateRequest)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "success",
		Data:   product,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *ProductsControllerImpl) GetAllProducts(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	products := controller.ProductsService.GetAllProducts(request.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "success",
		Data:   products,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *ProductsControllerImpl) GetProductByID(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("productId")
	productId, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	product := controller.ProductsService.GetProductById(request.Context(), productId)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "success",
		Data:   product,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *ProductsControllerImpl) UpdateProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productUpdateRequest := web.ProductUpdateRequest{}
	helper.ReadRequestBody(request, &productUpdateRequest)

	id := params.ByName("productId")
	productId, err := strconv.Atoi(id)
	helper.PanicIfError(err)
	productUpdateRequest.Id = productId

	product := controller.ProductsService.UpdateProduct(request.Context(), productUpdateRequest)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "success",
		Data:   product,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *ProductsControllerImpl) DeleteProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("productId")
	productId, err := strconv.Atoi(id)
	helper.PanicIfError(err)
	controller.ProductsService.DeleteProduct(request.Context(), productId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "success",
	}

	helper.WriteResponseBody(writer, webResponse)
}
