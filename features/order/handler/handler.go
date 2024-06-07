// features/order_product/handler/handler.go
package handler

import (
	"PetPalApp/app/middlewares"
	order "PetPalApp/features/order"
	"PetPalApp/utils/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	OrderService order.OrderService
}

func New(os order.OrderService) *OrderHandler {
	return &OrderHandler{
		OrderService: os,
	}
}

func (oh *OrderHandler) CreateOrder(c echo.Context) error {
    userID, _, _ := middlewares.ExtractTokenUserId(c)
    if userID == 0 {
        return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("Unauthorized", nil))
    }

    var newOrderReq OrderRequest
    if err := c.Bind(&newOrderReq); err != nil {
        return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Invalid input", nil))
    }

    product, err := oh.OrderService.GetProductById(newOrderReq.ProductID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Failed to get product", nil))
    }

    newOrder := order.OrderCore{
        UserID:         uint(userID),
        ProductID:      newOrderReq.ProductID,
        ProductName:    product.ProductName,
        ProductPicture: product.ProductPicture,
        Quantity:       uint(newOrderReq.Quantity),
        Price:          product.Price * float64(newOrderReq.Quantity), 
        Status:         "Pending",
    }

    err = oh.OrderService.CreateOrder(newOrder)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Failed to create order", nil))
    }

    return c.JSON(http.StatusCreated, responses.JSONWebResponse("Order created successfully", nil))
}




func (oh *OrderHandler) GetOrdersByUserID(c echo.Context) error {
    userID, _, _ := middlewares.ExtractTokenUserId(c)
    if userID == 0 {
        return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("Unauthorized", nil))
    }

    orders, err := oh.OrderService.GetOrdersByUserID(uint(userID))
    if err != nil {
        return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Error retrieving orders: "+err.Error(), nil))
    }

    var resultResponse []OrderResponse
    for _, order := range orders {
        resultResponse = append(resultResponse, CoreToResponse(order))
    }

    return c.JSON(http.StatusOK, responses.JSONWebResponse("Orders retrieved successfully", resultResponse))
}
