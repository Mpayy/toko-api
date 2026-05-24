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
	"github.com/sirupsen/logrus"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
	Logger             *logrus.Logger
}

func NewCategoryService(categoryRepository repository.CategoryRepository, db *sql.DB, validate *validator.Validate, logger *logrus.Logger) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 db,
		Validate:           validate,
		Logger:             logger,
	}
}

func (service *CategoryServiceImpl) CreateCategory(ctx context.Context, requestCreate web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(requestCreate)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: requestCreate.Name,
	}

	category = service.CategoryRepository.CreateCategory(ctx, tx, category)
	service.Logger.WithField("category", category.Id).Info("Category created")
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) GetAllCategories(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.GetAllCategories(ctx, tx)

	return helper.ToCategoryResponses(categories)
}

func (service *CategoryServiceImpl) GetCategoryById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	category, err := service.CategoryRepository.GetCategoryById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotfoundError(err.Error()))
	}
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) UpdateCategory(ctx context.Context, requestUpdate web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(requestUpdate)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.GetCategoryById(ctx, tx, requestUpdate.Id)
	if err != nil {
		panic(exception.NewNotfoundError(err.Error()))
	}

	category.Name = requestUpdate.Name

	category = service.CategoryRepository.UpdateCategory(ctx, tx, category)
	service.Logger.WithField("category", category.Id).Info("Category updated")
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) DeleteCategory(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.GetCategoryById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotfoundError(err.Error()))
	}

	service.CategoryRepository.DeleteCategory(ctx, tx, category)
	service.Logger.WithField("category", category.Id).Info("Category deleted")
}
