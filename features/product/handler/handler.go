package handler

import (
	"PetPalApp/app/middlewares"
	"PetPalApp/features/product"
	"PetPalApp/utils/helper"
	"PetPalApp/utils/responses"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productService product.ServiceInterface
	helper         helper.HelperInterface
}

func New(ps product.ServiceInterface, helper helper.HelperInterface) *ProductHandler {
	return &ProductHandler{
		productService: ps,
		helper:         helper,
	}
}

func (ph *ProductHandler) AddProduct(c echo.Context) error {

	idToken, _, _ := middlewares.ExtractTokenUserId(c) // extract id user from jwt token
	newProduct := ProductRequest{}
	errBind := c.Bind(&newProduct)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Error bind data"+errBind.Error(), nil))
	}
	file, handler, err := c.Request().FormFile("product_picture")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Photo is required. Please upload a valid photo.",
		})
	}
	defer file.Close()

	inputCore := RequestToCore(newProduct)
	_, errInsert := ph.productService.Create(uint(idToken), inputCore, file, handler.Filename)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Failed to add product. Please ensure all fields are filled in correctly.", nil))
	}
	return c.JSON(http.StatusCreated, responses.JSONWebResponse("Add product successful!", nil))
}

func (ph *ProductHandler) GetAllProduct(c echo.Context) error {
	page := c.QueryParam("page")
	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		pageInt = 1
	}
	offset := (pageInt - 1) * 1
	sortStr := c.QueryParam("sort")
	limitProduct := c.QueryParam("limit")
	limit, errlimit := strconv.Atoi(limitProduct)
	if errlimit != nil || pageInt < 1 {
		pageInt = 1
	}

	idToken, role, _ := middlewares.ExtractTokenUserId(c) // extract id user from jwt token

	result, errResult := ph.productService.GetAll(uint(idToken), uint(limit), role, uint(offset), sortStr)
	if errResult != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Failed to retrieve products", nil))
	}

	var allProduct []AllProductResponse
	for _, v := range result {
		allProduct = append(allProduct, AllGormToCore(v))
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse("Products retrieved successfully", allProduct))
}

func (ph *ProductHandler) GetProductById(c echo.Context) error {

	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("ID must be a positive integer", idConv))
	}

	idToken, _, _ := middlewares.ExtractTokenUserId(c) // extract id user from jwt token

	productData, errProductData := ph.productService.GetProductById(uint(idConv), uint(idToken))
	if errProductData != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Failed to retrieve products", nil))
	}

	productResponse := GormToCore(*productData)
	return c.JSON(http.StatusOK, responses.JSONWebResponse("Products retrieved successfully", productResponse))
}

func (ph *ProductHandler) UpdateProductById(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("ID must be a positive integer", idConv))
	}

	idToken, _, _ := middlewares.ExtractTokenUserId(c)

	updatedProduct := ProductRequest{}
	errBind := c.Bind(&updatedProduct)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Error bind data: "+errBind.Error(), nil))
	}

	var file multipart.File
	var handler *multipart.FileHeader

	file, handler, _ = c.Request().FormFile("product_picture")
	if file == nil && handler == nil {
		file = nil
		handler = nil
	}
	inputCore := RequestToCore(updatedProduct)

	var filename string
	if handler != nil {
		filename = handler.Filename
	}

	_, errUpdate := ph.productService.UpdateById(uint(idConv), uint(idToken), inputCore, file, filename)
	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Failed to update product", errUpdate))
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse("Product updated successfully", nil))
}

func (ph *ProductHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("ID must be a positive integer", idConv))
	}

	idToken, _, _ := middlewares.ExtractTokenUserId(c)
	err := ph.productService.Delete(uint(idConv), uint(idToken))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Failed to delete product", err))
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse("Product deleted successfully", nil))
}
