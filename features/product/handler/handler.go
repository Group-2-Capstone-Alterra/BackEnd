package handler

import (
	"PetPalApp/app/middlewares"
	"PetPalApp/features/product"
	"PetPalApp/utils/helper"
	"PetPalApp/utils/responses"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

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
	log.Println("idtoken:", idToken)

	newProduct := ProductRequest{}
	errBind := c.Bind(&newProduct)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("err bind"+errBind.Error(), nil))
	}

	file, handler, err := c.Request().FormFile("product_picture")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Unable to upload photo: " + err.Error(),
		})
	}
	defer file.Close()

	inputCore := RequestToCore(newProduct)
	_, errInsert := ph.productService.Create(uint(idToken), inputCore, file, handler.Filename)
	if errInsert != nil {
		if strings.Contains(errInsert.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error add data", errInsert))
		}
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error add data", errInsert))
	}
	return c.JSON(http.StatusCreated, responses.JSONWebResponse("success add data", nil))
}

func (ph *ProductHandler) GetAllProduct(c echo.Context) error {
	log.Println("[handler]")
	page := c.QueryParam("page")
	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		pageInt = 1
	}
	offset := (pageInt - 1) * 10

	sortStr := c.QueryParam("sort")

	idToken, role, _ := middlewares.ExtractTokenUserId(c) // extract id user from jwt token
	log.Println("idtoken:", idToken)

	result, errResult := ph.productService.GetAll(uint(idToken), role, uint(offset), sortStr)
	if errResult != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error read data", nil))
	}

	var allProduct []AllProductResponse
	for _, v := range result {
		allProduct = append(allProduct, AllGormToCore(v))
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse("success read data", allProduct))
}

func (ph *ProductHandler) GetProductById(c echo.Context) error {

	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error get id", idConv))
	}

	idToken, _, _ := middlewares.ExtractTokenUserId(c) // extract id user from jwt token
	log.Println("idtoken:", idToken)

	productData, errProductData := ph.productService.GetProductById(uint(idConv), uint(idToken))
	if errProductData != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error read data", nil))
	}

	productResponse := GormToCore(*productData)
	return c.JSON(http.StatusOK, responses.JSONWebResponse("success get detail product", productResponse))
}

func (ph *ProductHandler) UpdateProductById(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error get user id", idConv))
	}

	idToken, _, _ := middlewares.ExtractTokenUserId(c)

	updatedProduct := ProductRequest{}
	errBind := c.Bind(&updatedProduct)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error bind data: "+errBind.Error(), nil))
	}

	var file multipart.File
	var handler *multipart.FileHeader
	var err error

	file, handler, err = c.Request().FormFile("product_picture")
	if err != nil {
		if err != http.ErrMissingFile {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Unable to upload photo: " + err.Error(),
			})
		}
		// Handle the case where no file was uploaded
		file = nil
		handler = nil
	} else {
		defer file.Close()
	}

	inputCore := RequestToCore(updatedProduct)

	var filename string
	if handler != nil {
		filename = handler.Filename
	}

	_, errUpdate := ph.productService.UpdateById(uint(idConv), uint(idToken), inputCore, file, filename)
	if errUpdate != nil {
		// Handle error from userService.UpdateById
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error update data", err))
	}
	// Return success response
	return c.JSON(http.StatusOK, responses.JSONWebResponse("success update data", err))
}

func (ph *ProductHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error get user id", idConv))
	}

	idToken, _, _ := middlewares.ExtractTokenUserId(c)
	err := ph.productService.Delete(uint(idConv), uint(idToken))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error delete data", err))
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse("success delete data", err))
}
