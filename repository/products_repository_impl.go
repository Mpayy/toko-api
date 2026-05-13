package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Mpayy/toko-api/helper"
	"github.com/Mpayy/toko-api/model/domain"
)

type ProductRepositoryImpl struct{}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) CreateProduct(ctx context.Context, tx *sql.Tx, products domain.Products) domain.Products {
	querySql := "insert into products (name, price, stock) values (?,?,?)"
	result, err := tx.ExecContext(ctx, querySql, products.Name, products.Price, products.Stock)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	products.Id = int(id)
	return products
}

func (repository *ProductRepositoryImpl) GetAllProducts(ctx context.Context, tx *sql.Tx) []domain.Products {
	querySql := "select id, name, price, stock from products"
	rows, err := tx.QueryContext(ctx, querySql)
	helper.PanicIfError(err)
	defer rows.Close()
	var products []domain.Products
	for rows.Next() {
		var product domain.Products
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Stock)
		helper.PanicIfError(err)
		products = append(products, product)
	}
	return products
}

func (repository *ProductRepositoryImpl) GetProductById(ctx context.Context, tx *sql.Tx, productId int) (domain.Products, error) {
	querySql := "select id, name, price, stock from products where id = ?"
	rows, err := tx.QueryContext(ctx, querySql, productId)
	helper.PanicIfError(err)
	defer rows.Close()
	var product domain.Products
	if rows.Next() {
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Stock)
		helper.PanicIfError(err)
		return product, nil
	}
	return product, errors.New("product not found")
}

func (repository *ProductRepositoryImpl) UpdateProduct(ctx context.Context, tx *sql.Tx, products domain.Products) domain.Products {
	querySql := "update products set name = ?, price = ?, stock = ? where id = ?"
	_, err := tx.ExecContext(ctx, querySql, products.Name, products.Price, products.Stock, products.Id)
	helper.PanicIfError(err)
	return products
}

func (repository *ProductRepositoryImpl) DeleteProduct(ctx context.Context, tx *sql.Tx, product domain.Products) {
	querySql := "delete from products where id = ?"
	_, err := tx.ExecContext(ctx, querySql, product.Id)
	helper.PanicIfError(err)
}
