package service

import (
	"context"

	"github.com/Mpayy/toko-api/model/web"
)

type ProductService interface {
	CreateProduct(ctx context.Context, requestCreate web.ProductCreateRequest) web.ProductResponse
	GetAllProducts(ctx context.Context) []web.ProductResponse
	GetProductById(ctx context.Context, productId int) web.ProductResponse
	UpdateProduct(ctx context.Context, requestUpdate web.ProductUpdateRequest) web.ProductResponse
	DeleteProduct(ctx context.Context, productId int)
}
