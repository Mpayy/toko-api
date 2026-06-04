package test

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/Mpayy/toko-api/app"
	"github.com/Mpayy/toko-api/controller"
	"github.com/Mpayy/toko-api/helper"
	"github.com/Mpayy/toko-api/middleware"
	"github.com/Mpayy/toko-api/repository"
	"github.com/Mpayy/toko-api/service"
	"github.com/stretchr/testify/assert"
)

func setupTestDb() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/toko-api-test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)

	return db
}

func setupTestRouter(db *sql.DB) http.Handler {

	config := app.NewViper()
	validate := app.NewValidator()
	logger := app.NewLogger()
	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db, validate, logger)
	productsController := controller.NewProductsController(productService)

	router := app.NewRouter(productsController, nil, logger)

	return middleware.NewMiddleware(config, router, logger)
}

func truncateProduct(db *sql.DB) {
	db.Exec("TRUNCATE TABLE products")

}

func TestCreateCategorySuccess(t *testing.T) {
	db := setupTestDb()
	truncateProduct(db)
	router := setupTestRouter(db)

	requestBody := strings.NewReader(`{"name" : "product test", "price" : 10000, "stock" : 10}`)
	request, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/api/products", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "RAHASIA")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusCreated, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	_ = json.Unmarshal(body, &responseBody)
	assert.Equal(t, 200, responseBody["code"])
	assert.Equal(t, "success", responseBody["message"])
}
