package handler

import (
	"PetPalApp/app/middlewares"
	"PetPalApp/features/product"
	"PetPalApp/utils/responses"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productService product.ServiceInterface
}

func New(ps product.ServiceInterface) *ProductHandler {
	return &ProductHandler{
		productService: ps,
	}
}

func (ph *ProductHandler) AddProduct(c echo.Context) error {

	idToken := middlewares.ExtractTokenUserId(c) // extract id user from jwt token
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

	page := c.QueryParam("page")
	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		pageInt = 1
	}
	offset := (pageInt - 1) * 10

	sortStr := c.QueryParam("sort")

	idToken := middlewares.ExtractTokenUserId(c) // extract id user from jwt token
	log.Println("idtoken:", idToken)

	result, errResult := ph.productService.GetAll(uint(idToken), uint(offset), sortStr)
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

	idToken := middlewares.ExtractTokenUserId(c) // extract id user from jwt token
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

	idToken := middlewares.ExtractTokenUserId(c)

	updatedProduct := ProductRequest{}
	errBind := c.Bind(&updatedProduct)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error bind data: "+errBind.Error(), nil))
	}

	err := ph.productService.UpdateById(uint(idConv), uint(idToken), RequestToCore(updatedProduct))
	if err != nil {
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

	idToken := middlewares.ExtractTokenUserId(c)
	err := ph.productService.Delete(uint(idToken), uint(idConv))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error delete data", err))
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse("success delete data", err))
}
