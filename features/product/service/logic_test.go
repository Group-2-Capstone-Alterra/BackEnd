package service

import (
	"PetPalApp/features/product"
	"PetPalApp/mocks"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateProduct(t *testing.T) {
	mockProductData := new(mocks.ProductModel)
	mockHelper := new(mocks.HelperInterface)
	mockAdminData := new(mocks.AdminModel)

	productService := New(mockProductData, mockHelper, mockAdminData)

	id := uint(1)
	input := product.Core{
		ProductName: "Test Product",
		Price:       10.99,
		Stock:       10,
		Description: "This is a test product",
	}
	file := strings.NewReader("test file content")
	handlerFilename := "test_file.jpg"

	mockHelper.On("UploadProductPicture", file, mock.MatchedBy(func(filename string) bool {
		return regexp.MustCompile(`^\d+_test_file\.jpg$`).MatchString(filename)
	})).Return("test_file.jpg", nil)
	mockProductData.On("Insert", input).Return(nil)

	result, err := productService.Create(id, input, file, handlerFilename)

	assert.Nil(t, err)
	assert.Equal(t, "test_file.jpg", result)

	mockHelper.AssertExpectations(t)
	mockProductData.AssertExpectations(t)
}

func TestGetAllProducts(t *testing.T) {
	mockProductData := new(mocks.ProductModel)
	mockHelper := new(mocks.HelperInterface)
	mockAdminData := new(mocks.AdminModel)

	productService := New(mockProductData, mockHelper, mockAdminData)

	userid := uint(1)
	limit := uint(10)
	role := "user"
	offset := uint(0)
	sortStr := "latest"

	mockProductData.On("SelectAll", limit, offset, sortStr).Return([]product.Core{
		{
			ProductName: "Product 1",
			Price:       10.99,
			Stock:       10,
			Description: "This is product 1",
		},
		{
			ProductName: "Product 2",
			Price:       20.99,
			Stock:       20,
			Description: "This is product 2",
		},
	}, nil)

	result, err := productService.GetAll(userid, limit, role, offset, sortStr)

	assert.Nil(t, err)
	assert.Len(t, result, 2)

	mockProductData.AssertExpectations(t)
}

func TestGetProductById(t *testing.T) {
	mockProductData := new(mocks.ProductModel)
	mockHelper := new(mocks.HelperInterface)
	mockAdminData := new(mocks.AdminModel)

	productService := New(mockProductData, mockHelper, mockAdminData)

	id := uint(1)
	userid := uint(1)

	mockProductData.On("SelectById", id, userid).Return(&product.Core{
		ProductName: "Product 1",
		Price:       10.99,
		Stock:       10,
		Description: "This is product 1",
	}, nil)

	result, err := productService.GetProductById(id, userid)

	assert.Nil(t, err)
	assert.NotNil(t, result)

	mockProductData.AssertExpectations(t)
}

func TestUpdateProduct(t *testing.T) {
	mockProductData := new(mocks.ProductModel)
	mockHelper := new(mocks.HelperInterface)
	mockAdminData := new(mocks.AdminModel)

	productService := New(mockProductData, mockHelper, mockAdminData)

	id := uint(1)
	userid := uint(1)
	input := product.Core{
		ProductName: "Updated Product",
		Price:       20.99,
		Stock:       20,
		Description: "This is an updated product",
	}
	file := strings.NewReader("updated file content")
	handlerFilename := "updated_file.jpg"

	mockHelper.On("UploadProductPicture", file, handlerFilename).Return("updated_file.jpg", nil)
	mockProductData.On("PutById", id, userid, input).Return(nil)

	result, err := productService.UpdateById(id, userid, input, file, handlerFilename)

	assert.Nil(t, err)
	assert.Equal(t, "updated_file.jpg", result)

	mockHelper.AssertExpectations(t)
	mockProductData.AssertExpectations(t)
}

func TestDeleteProduct(t *testing.T) {
	mockProductData := new(mocks.ProductModel)
	mockHelper := new(mocks.HelperInterface)
	mockAdminData := new(mocks.AdminModel)

	productService := New(mockProductData, mockHelper, mockAdminData)

	id := uint(1)
	userid := uint(1)

	mockProductData.On("Delete", id, userid).Return(nil)

	err := productService.Delete(id, userid)

	assert.Nil(t, err)

	mockProductData.AssertExpectations(t)
}
