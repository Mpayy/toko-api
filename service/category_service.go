package service

import (
	"context"

	"github.com/Mpayy/toko-api/model/web"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, requestCreate web.CategoryCreateRequest) web.CategoryResponse
	GetAllCategories(ctx context.Context) []web.CategoryResponse
	GetCategoryById(ctx context.Context, categoryId int) web.CategoryResponse
	UpdateCategory(ctx context.Context, requestUpdate web.CategoryUpdateRequest) web.CategoryResponse
	DeleteCategory(ctx context.Context, categoryId int)
}
