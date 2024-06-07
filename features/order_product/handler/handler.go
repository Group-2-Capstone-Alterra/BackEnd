// features/order_product/handler/handler.go
package handler

import (
	"PetPalApp/features/order_product"
	"PetPalApp/utils/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrderProductHandler struct {
	OrderProductService order_product.OrderProductService
}

func New(ops order_product.OrderProductService) *OrderProductHandler {
	return &OrderProductHandler{
		OrderProductService: ops,
	}
}

func (oph *OrderProductHandler) CreateOrderProduct(c echo.Context) error {
    var newOrderProductReq OrderProductRequest
    if err := c.Bind(&newOrderProductReq); err != nil {
        return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Invalid input", nil))
    }

    product, err := oph.OrderProductService.GetProductById(newOrderProductReq.ProductID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Failed to get product", nil))
    }

    newOrderProduct := order_product.OrderProductCore{
        OrderID:   newOrderProductReq.OrderID,
        ProductID: newOrderProductReq.ProductID,
        Quantity:  uint(newOrderProductReq.Quantity),
        Price:     product.Price * float64(newOrderProductReq.Quantity), 
    }

    err = oph.OrderProductService.CreateOrderProduct(newOrderProduct)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Failed to create order product", nil))
    }

    return c.JSON(http.StatusCreated, responses.JSONWebResponse("Order product created successfully", nil))
}




func (oph *OrderProductHandler) GetOrderProductsByOrderID(c echo.Context) error {
	orderID, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Invalid order ID", nil))
	}

	orderProducts, err := oph.OrderProductService.GetOrderProductsByOrderID(uint(orderID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Failed to get order products", nil))
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse("Order products retrieved successfully", orderProducts))
}
