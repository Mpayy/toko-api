package helper

import (
	"github.com/Mpayy/toko-api/model/domain"
	"github.com/Mpayy/toko-api/model/web"
)

func ToProductResponse(product domain.Products) web.ProductResponse {
	return web.ProductResponse{
		Id:    product.Id,
		Name:  product.Name,
		Price: product.Price,
		Stock: product.Stock,
	}
}

func ToProductResponses(products []domain.Products) []web.ProductResponse {
	var productsResponses []web.ProductResponse
	for _, product := range products {
		productsResponses = append(productsResponses, ToProductResponse(product))
	}
	return productsResponses
}

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}
