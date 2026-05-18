package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Mpayy/toko-api/helper"
	"github.com/Mpayy/toko-api/model/domain"
)

type CategoryRepositoryImpl struct{}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) CreateCategory(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	querySql := "insert into categories (name) values (?)"
	result, err := tx.ExecContext(ctx, querySql, category.Name)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	category.Id = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) GetAllCategories(ctx context.Context, tx *sql.Tx) []domain.Category {
	querySql := "select id, name from categories"
	rows, err := tx.QueryContext(ctx, querySql)
	helper.PanicIfError(err)
	defer rows.Close()
	var categories []domain.Category
	for rows.Next() {
		var category domain.Category
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}

func (repository *CategoryRepositoryImpl) GetCategoryById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	querySql := "select id, name from categories where id = ?"
	rows, err := tx.QueryContext(ctx, querySql, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()
	var category domain.Category
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	}
	return category, errors.New("category not found")
}

func (repository *CategoryRepositoryImpl) UpdateCategory(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	querySql := "update categories set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, querySql, category.Name, category.Id)
	helper.PanicIfError(err)
	return category
}

func (repository *CategoryRepositoryImpl) DeleteCategory(ctx context.Context, tx *sql.Tx, category domain.Category) {
	querySql := "delete from categories where id = ?"
	_, err := tx.ExecContext(ctx, querySql, category.Id)
	helper.PanicIfError(err)
}
