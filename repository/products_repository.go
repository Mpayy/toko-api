package repository

import (
	"context"
	"database/sql"

	"github.com/Mpayy/toko-api/model/domain"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, tx *sql.Tx, products domain.Products) domain.Products
	GetAllProducts(ctx context.Context, tx *sql.Tx) []domain.Products
	GetProductById(ctx context.Context, tx *sql.Tx, productId int) (domain.Products, error)
	UpdateProduct(ctx context.Context, tx *sql.Tx, products domain.Products) domain.Products
	DeleteProduct(ctx context.Context, tx *sql.Tx, productId domain.Products)
}
