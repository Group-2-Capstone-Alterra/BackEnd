// features/order_product/handler/handler.go
package handler

import (
	"PetPalApp/app/middlewares"
	order "PetPalApp/features/order"
	"PetPalApp/utils/responses"
	"net/http"
	"strconv"

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

    product, err := oh.OrderService.GetProductByID(newOrderReq.ProductID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Failed to get product", nil))
    }

    newOrder := order.Order{
        UserID:         uint(userID),
        ProductID:      newOrderReq.ProductID,
        ProductName:    product.ProductName,
        ProductPicture: product.ProductPicture,
        Quantity:       uint(newOrderReq.Quantity),
        Price:          product.Price * float64(newOrderReq.Quantity), 
    }

    createdOrder, err := oh.OrderService.CreateOrder(newOrder)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Failed to create order", nil))
    }

    orderResponse := CreatedResponse{
        ID: createdOrder.ID,
    }

    return c.JSON(http.StatusCreated, responses.JSONWebResponse("Order created successfully", orderResponse))
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

func (oh *OrderHandler) GetOrderByID(c echo.Context) error {
    userID, _, _ := middlewares.ExtractTokenUserId(c)
    if userID == 0 {
        return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("Unauthorized", nil))
    }

    orderIDParam := c.Param("id")
    orderID, err := strconv.ParseUint(orderIDParam, 10, 64)
    if err != nil {
        return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Invalid order ID", nil))
    }

    order, err := oh.OrderService.GetOrderByID(uint(orderID))
    if err != nil {
        return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Error retrieving order: "+err.Error(), nil))
    }

    // Ensure that the order belongs to the current user
    if order.UserID != uint(userID) {
        return c.JSON(http.StatusForbidden, responses.JSONWebResponse("You are not authorized to view this order", nil))
    }


    newResponse := OrderResponse{
        ID: order.ID,
        UserID: order.UserID,
        ProductID: order.ProductID,
        ProductName: order.ProductName,
        ProductPicture: order.ProductPicture,
        Quantity: order.Quantity,
        Price: order.Price,
        Status: order.Status,
        Payment: PaymentResponse{
            ID: order.Payment.ID,
            OrderID: order.Payment.OrderID,
            PaymentMethod: order.Payment.PaymentMethod,
            PaymentStatus: order.Payment.PaymentStatus,
            SignatureID:   order.Payment.SignatureID,
            VANumber:      order.Payment.VANumber,
            InvoiceID:     order.Payment.InvoiceID,
        },

    }

    return c.JSON(http.StatusOK, responses.JSONWebResponse("Order retrieved successfully", newResponse))
}
