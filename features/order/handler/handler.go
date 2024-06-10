// features/order_product/handler/handler.go
package handler

import (
	"PetPalApp/app/middlewares"
	order "PetPalApp/features/order"
	"PetPalApp/features/product"
	"PetPalApp/utils/responses"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	OrderService order.OrderService
	ProductData  product.ProductModel
}

func New(os order.OrderService, ProductData product.ProductModel) *OrderHandler {
	return &OrderHandler{
		OrderService: os,
		ProductData:  ProductData,
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
		Status:         "Created",
		InvoiceID:      generateInvoiceID(),
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

// func (oh *OrderHandler) GetOrdersByUserID(c echo.Context) error {
//     userID, _, _ := middlewares.ExtractTokenUserId(c)
//     if userID == 0 {
//         return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("Unauthorized", nil))
//     }

//     orders, err := oh.OrderService.GetOrdersByUserID(uint(userID))
//     if err != nil {
//         return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Error retrieving orders: "+err.Error(), nil))
//     }

//     var response []OrderResponse
//     for _, order := range orders {
//     response := append (response, OrderResponse{
//         ID: order[len(order)-1].ID,
//         UserID: order[len(order)-1].UserID,
//         ProductID: order[len(order)-1].ProductID,
//         ProductName: order[len(order)-1].ProductName,
//         ProductPicture: order[len(order)-1].ProductPicture,
//         Quantity: order[len(order)-1].Quantity,
//         Price: order[len(order)-1].Price,
//         Status: order[len(order)-1].Status,
//         Payment: PaymentResponse{
//             ID: order[len(order)-1].Payment.ID,
//             OrderID: order[len(order)-1].Payment.OrderID,
//             PaymentMethod: order[len(order)-1].Payment.PaymentMethod,
//             PaymentStatus: order[len(order)-1].Payment.PaymentStatus,
//             SignatureID:   order[len(order)-1].Payment.SignatureID,
//             VANumber:      order[len(order)-1].Payment.VANumber,
//             InvoiceID:     order[len(order)-1].InvoiceID,
//         },
//     })
//     }
//     return c.JSON(http.StatusOK, responses.JSONWebResponse("Orders retrieved successfully", response))
// }

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
		ID:             order.ID,
		UserID:         order.UserID,
		ProductID:      order.ProductID,
		ProductName:    order.ProductName,
		ProductPicture: order.ProductPicture,
		Quantity:       order.Quantity,
		Price:          order.Price,
		Status:         order.Status,
		Payment: PaymentResponse{
			ID:            order.Payment.ID,
			OrderID:       order.Payment.OrderID,
			PaymentMethod: order.Payment.PaymentMethod,
			PaymentStatus: order.Payment.PaymentStatus,
			SignatureID:   order.Payment.SignatureID,
			VANumber:      order.Payment.VANumber,
			InvoiceID:     order.Payment.InvoiceID,
		},
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse("Order retrieved successfully", newResponse))
}

func generateInvoiceID() string {
	randomNumber := rand.Intn(9000) + 1000
	currentDate := time.Now().Format("02012006")
	invoiceID := fmt.Sprintf("ORDER-%s-%d", currentDate, randomNumber)

	return invoiceID
}

func (oh *OrderHandler) GetOrdersByUserID(c echo.Context) (err error) {
	var orders []order.Order
	var product []product.Core
	userID, role, _ := middlewares.ExtractTokenUserId(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("Unauthorized", nil))
	}

	if role == "user" {
		orders, err = oh.OrderService.GetOrdersByUserID(uint(userID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Error retrieving orders: "+err.Error(), nil))
		}
	} else {
		log.Printf("\n userID: %v, role: %v", userID, role)
		product, _ = oh.ProductData.SelectAllAdmin(100, uint(userID), 0)
		log.Println("all product", product)
		for i, v := range product {
			log.Printf("\nv id %v and product %v\n", i, product)
			orders, err = oh.OrderService.GetOrdersByProductAdmin(v.ID)
			log.Println("orders", orders)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Error retrieving orders: "+err.Error(), nil))
			}
		}
	}

	var response []OrderResponse
	for _, order := range orders {
		response = append(response, OrderResponse{
			ID:             order.ID,
			UserID:         order.UserID,
			ProductID:      order.ProductID,
			ProductName:    order.ProductName,
			ProductPicture: order.ProductPicture,
			Quantity:       order.Quantity,
			Price:          order.Price,
			Status:         order.Status,
			Payment: PaymentResponse{
				ID:            order.Payment.ID,
				OrderID:       order.Payment.OrderID,
				PaymentMethod: order.Payment.PaymentMethod,
				PaymentStatus: order.Payment.PaymentStatus,
				SignatureID:   order.Payment.SignatureID,
				VANumber:      order.Payment.VANumber,
				InvoiceID:     order.InvoiceID,
			},
		})
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse("Orders retrieved successfully", response))
}
