package handler

import (
	"PetPalApp/features/order"
	"PetPalApp/utils/helper"
	"PetPalApp/utils/responses"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/veritrans/go-midtrans"
)

type WebhookHandler struct {
	orderService    order.OrderService
	midtrans        midtrans.Client
}

func New(os order.OrderService, midtrans midtrans.Client) *WebhookHandler {
	return &WebhookHandler{
		orderService: os,
		midtrans: midtrans,
	}
}

func (wh *WebhookHandler) MidtransWebhook(c echo.Context) error {
	var notification MidtransNotificationRequest

	if err := c.Bind(&notification); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	orderData, err := wh.orderService.GetOrderByInvoiceID(notification.OrderId)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Failed to get product", nil))
    }

	client := wh.midtrans
	coreGateway := midtrans.CoreGateway{
        Client: client,
    }
	resp, err := coreGateway.Status(orderData.InvoiceID)
	if err != nil || !helper.ContainsString([]string{"201", "200"}, resp.StatusCode) {
		log.Printf("Check Status Transaction failed with error: %v", err)
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Failed to charge payment", nil))
	} else {
		orderAttr := order.Order{
			Status: resp.TransactionStatus,
    	}
		updateOrder, err := wh.orderService.UpdateOrder(orderData.ID, orderAttr)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Failed to update order", nil))
		}
		log.Printf("Update Transaction: %v", updateOrder)
	}
	
	log.Printf("Received notification: %+v\n", notification)

	return c.JSON(http.StatusOK, "Notification received")
}