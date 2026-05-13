package service

import (
	"context"
	"database/sql"

	"github.com/Mpayy/toko-api/exception"
	"github.com/Mpayy/toko-api/helper"
	"github.com/Mpayy/toko-api/model/domain"
	"github.com/Mpayy/toko-api/model/web"
	"github.com/Mpayy/toko-api/repository"
	"github.com/go-playground/validator/v10"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewProductService(productRepository repository.ProductRepository, DB *sql.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *ProductServiceImpl) CreateProduct(ctx context.Context, requestCreate web.ProductCreateRequest) web.ProductResponse {
	err := service.Validate.Struct(requestCreate)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product := domain.Products{
		Name:  requestCreate.Name,
		Price: requestCreate.Price,
		Stock: requestCreate.Stock,
	}

	product = service.ProductRepository.CreateProduct(ctx, tx, product)

	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) GetAllProducts(ctx context.Context) []web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	products := service.ProductRepository.GetAllProducts(ctx, tx)
	return helper.ToProductResponses(products)
}

func (service *ProductServiceImpl) GetProductById(ctx context.Context, productId int) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.GetProductById(ctx, tx, productId)
	helper.PanicIfError(err)

	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) UpdateProduct(ctx context.Context, requestUpdate web.ProductUpdateRequest) web.ProductResponse {
	err := service.Validate.Struct(requestUpdate)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.GetProductById(ctx, tx, requestUpdate.Id)
	if err != nil {
		panic(exception.NewNotfoundError(err.Error()))
	}

	product.Name = requestUpdate.Name
	product.Price = requestUpdate.Price
	product.Stock = requestUpdate.Stock

	product = service.ProductRepository.UpdateProduct(ctx, tx, product)
	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) DeleteProduct(ctx context.Context, productId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.GetProductById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotfoundError(err.Error()))
	}

	service.ProductRepository.DeleteProduct(ctx, tx, product)
}
