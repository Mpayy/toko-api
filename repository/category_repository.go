package repository

import (
	"context"
	"database/sql"

	"github.com/Mpayy/toko-api/model/domain"
)

type CategoryRepository interface {
	CreateCategory(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	GetAllCategories(ctx context.Context, tx *sql.Tx) []domain.Category
	GetCategoryById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	UpdateCategory(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	DeleteCategory(ctx context.Context, tx *sql.Tx, category domain.Category)
}
