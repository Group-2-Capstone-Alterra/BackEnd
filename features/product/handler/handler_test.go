package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"PetPalApp/features/product"
	"PetPalApp/mocks"
)

func TestProductHandlerAddProduct(t *testing.T) {
	mockProductService := &mocks.ProductService{}
	mockHelper := &mocks.HelperInterface{}

	productHandler := New(mockProductService, mockHelper)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/products", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	idToken := uint(1)
	mockProductService.On("Create", idToken, mock.Anything, mock.Anything, mock.Anything).Return("product_id", nil)

	productRequest := ProductRequest{
		ProductName:    "Product 1",
		Price:          10000,
		Stock:          10,
		Description:    "This is product 1",
		ProductPicture: "product1.jpg",
	}

	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(productRequest)
	if err != nil {
		t.Fatalf("failed to encode product request: %v", err)
	}
	req.Body = ioutil.NopCloser(buf)
	req.ContentLength = int64(buf.Len())

	err = productHandler.AddProduct(c)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestProductHandlerGetAllProduct(t *testing.T) {
	mockProductService := &mocks.ProductService{}
	mockHelper := &mocks.HelperInterface{}

	productHandler := New(mockProductService, mockHelper)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/products?page=1&limit=10&sort=asc", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	idToken := uint(1)
	mockProductService.On("GetAll", idToken, uint(10), "asc", uint(0), "").Return([]product.Core{{ID: 1, ProductName: "Product 1"}}, nil)

	err := productHandler.GetAllProduct(c)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestProductHandlerGetProductById(t *testing.T) {
	mockProductService := &mocks.ProductService{}
	mockHelper := &mocks.HelperInterface{}

	productHandler := New(mockProductService, mockHelper)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/products/1", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	idToken := uint(1)
	mockProductService.On("GetProductById", uint(1), idToken).Return(&product.Core{ID: 1, ProductName: "Product 1"}, nil)

	err := productHandler.GetProductById(c)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestProductHandlerUpdateProductById(t *testing.T) {
	mockProductService := &mocks.ProductService{}
	mockHelper := &mocks.HelperInterface{}

	productHandler := New(mockProductService, mockHelper)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/products/1", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	idToken := uint(1)
	mockProductService.On("UpdateById", uint(1), idToken, mock.Anything, mock.Anything, mock.Anything).Return("product_id", nil)

	productRequest := ProductRequest{
		ProductName:    "Product 1 Updated",
		Price:          15000,
		Stock:          20,
		Description:    "This is product 1 updated",
		ProductPicture: "product1_updated.jpg",
	}
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(productRequest)
	assert.NoError(t, err)

	err = productHandler.UpdateProductById(c)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestProductHandlerDeleteProduct(t *testing.T) {
	mockProductService := &mocks.ProductService{}
	mockHelper := &mocks.HelperInterface{}

	productHandler := New(mockProductService, mockHelper)

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/products/1", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	idToken := uint(1)
	mockProductService.On("Delete", uint(1), idToken).Return(nil)

	err := productHandler.Delete(c)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
}
