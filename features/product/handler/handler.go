package handler

import (
	"PetPalApp/app/middlewares"
	"PetPalApp/features/product"
	"PetPalApp/utils/responses"
	"log"
	"net/http"
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
