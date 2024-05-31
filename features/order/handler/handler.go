package handler

import (
	"PetPalApp/app/middlewares"
	"PetPalApp/features/order"
	"PetPalApp/utils/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
    orderService order.OrderService
}

func New(os order.OrderService) *OrderHandler {
    return &OrderHandler{
        orderService: os,
    }
}

func (oh *OrderHandler) CreateOrder(c echo.Context) error {
    userID := middlewares.ExtractTokenUserId(c)
    if userID == 0 {
        return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("Unauthorized", nil))
    }

    newOrder := OrderRequest{}
    if err := c.Bind(&newOrder); err != nil {
        return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Error binding data: "+err.Error(), nil))
    }

    orderData := order.OrderCore{
        UserID:    uint(userID),
        ProductID: newOrder.ProductID,
        Quantity:  newOrder.Quantity,
        Total:     newOrder.Total,
        Status:    "Pending", // default status
    }

    if err := oh.orderService.CreateOrder(orderData); err != nil {
        return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Error creating order: "+err.Error(), nil))
    }

    return c.JSON(http.StatusCreated, responses.JSONWebResponse("Order created successfully", nil))
}

func (oh *OrderHandler) GetOrdersByUserID(c echo.Context) error {
    userID := middlewares.ExtractTokenUserId(c)
    if userID == 0 {
        return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("Unauthorized", nil))
    }

    orders, err := oh.orderService.GetOrdersByUserID(uint(userID))
    if err != nil {
        return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Error retrieving orders: "+err.Error(), nil))
    }

    return c.JSON(http.StatusOK, responses.JSONWebResponse("Orders retrieved successfully", orders))
}

