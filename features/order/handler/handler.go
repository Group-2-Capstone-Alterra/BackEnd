// features/order_product/handler/handler.go
package handler

import (
	"PetPalApp/app/middlewares"
	"PetPalApp/features/order"
	"PetPalApp/utils/helper"
	"PetPalApp/utils/responses"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/veritrans/go-midtrans"
)

type OrderHandler struct {
	OrderService order.OrderService
	midtrans        midtrans.Client
}

func New(os order.OrderService, midtrans midtrans.Client) *OrderHandler {
	return &OrderHandler{
		OrderService: os,
		midtrans: midtrans,
	}
}

func (oh *OrderHandler) CreateOrder(c echo.Context) error {
    userID, _, _ := middlewares.ExtractTokenUserId(c)
    if userID == 0 {
        return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("Unauthorized", nil))
    }

    var newOrderReq OrderCreateRequest
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
        Status:         "created",
        InvoiceID:      generateInvoiceID(),
    }

    order, err := oh.OrderService.CreateOrder(newOrder)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Failed to create order", nil))
    }

    orderResponse := OrderResponse{
        ID: order.ID,
        UserID: order.UserID,
        ProductID: order.ProductID,
        ProductName: order.ProductName,
        ProductPicture: order.ProductPicture,
        Quantity: order.Quantity,
        Price: order.Price,
        Status: order.Status,
        InvoiceID: order.InvoiceID,
    }
    return c.JSON(http.StatusCreated, responses.JSONWebResponse("Order created successfully", orderResponse))
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

    if order.UserID != uint(userID) {
        return c.JSON(http.StatusForbidden, responses.JSONWebResponse("You are not authorized to view this order", nil))
    }


    orderResponse := OrderResponse{
        ID: order.ID,
        UserID: order.UserID,
        ProductID: order.ProductID,
        ProductName: order.ProductName,
        ProductPicture: order.ProductPicture,
        Quantity: order.Quantity,
        Price: order.Price,
        Status: order.Status,
        InvoiceID:     order.InvoiceID,
        Payment: PaymentResponse{
            ID: order.Payment.ID,
            OrderID: order.Payment.OrderID,
            PaymentMethod: order.Payment.PaymentMethod,
            SignatureID:   order.Payment.SignatureID,
            VANumber:      order.Payment.VANumber,
        },
    }

    return c.JSON(http.StatusOK, responses.JSONWebResponse("Order retrieved successfully", orderResponse))
}

func generateInvoiceID() string {
	randomNumber := rand.Intn(9000) + 1000
	currentDate := time.Now().Format("02012006")
	invoiceID := fmt.Sprintf("ORDER-%s-%d", currentDate, randomNumber)

	return invoiceID
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

    var response []OrderResponse
    for _, order := range orders {
        response = append(response, OrderResponse{
            ID: order.ID,
            UserID: order.UserID,
            ProductID: order.ProductID,
            ProductName: order.ProductName,
            ProductPicture: order.ProductPicture,
            Quantity: order.Quantity,
            Price: order.Price,
            Status: order.Status,
            InvoiceID: order.InvoiceID,
            Payment: PaymentResponse{
                ID: order.Payment.ID,
                OrderID: order.Payment.OrderID,
                PaymentMethod: order.Payment.PaymentMethod,
                SignatureID:   order.Payment.SignatureID,
                VANumber:      order.Payment.VANumber,
            },
        })
    }

    return c.JSON(http.StatusOK, responses.JSONWebResponse("Orders retrieved successfully", response))
}


func (oh *OrderHandler) UpdateStatus(c echo.Context) error {
    var orderParam OrderUpdateStatusRequest
    if err := c.Bind(&orderParam); err != nil {
        return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Invalid input", nil))
    }

    orderData, err := oh.OrderService.GetOrderByID(orderParam.ID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Error retrieving order: "+err.Error(), nil))
    }

    client := oh.midtrans
    coreGateway := midtrans.CoreGateway{
        Client: client,
    }
    resp, err := coreGateway.Status(orderData.InvoiceID)
    log.Printf("Check Status Request: %+v\n", resp)

    OrderStatus := orderData.Status
    if helper.ContainsString([]string{"201", "200"}, resp.StatusCode) {
        OrderStatus = resp.TransactionStatus
        orderAttr := order.Order{
            Status: OrderStatus,
        }
        orderData, err := oh.OrderService.UpdateOrder(orderParam.ID, orderAttr)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Failed to update payment", nil))
        }
        log.Printf("updateOrder Request: %+v\n", orderData)
    }

    orderResponse := OrderResponse{
        ID: orderData.ID,
        UserID: orderData.UserID,
        ProductID: orderData.ProductID,
        ProductName: orderData.ProductName,
        ProductPicture: orderData.ProductPicture,
        Quantity: orderData.Quantity,
        Price: orderData.Price,
        Status: OrderStatus,
        InvoiceID: orderData.InvoiceID,
        Payment: PaymentResponse{
            ID: orderData.Payment.ID,
            OrderID: orderData.Payment.OrderID,
            PaymentMethod: orderData.Payment.PaymentMethod,
            SignatureID:   orderData.Payment.SignatureID,
            VANumber:      orderData.Payment.VANumber,
        },
    }

    return c.JSON(http.StatusOK, responses.JSONWebResponse("Orders retrieved successfully", orderResponse))
}