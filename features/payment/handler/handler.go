package handler

import (
	newOrder "PetPalApp/features/order"
	newOrderResponse "PetPalApp/features/order/handler"
	"PetPalApp/features/payment"
	"PetPalApp/utils/helper"
	"PetPalApp/utils/responses"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/veritrans/go-midtrans"
)

type PaymentHandler struct {
	paymentService  payment.PaymentService
	orderService    newOrder.OrderService
	midtrans        midtrans.Client
}

func New(ps payment.PaymentService, os newOrder.OrderService, midtrans midtrans.Client) *PaymentHandler {
	return &PaymentHandler{
		paymentService: ps,
		orderService: os,
		midtrans: midtrans,
	}
}

func (ph *PaymentHandler) CreatePayment(c echo.Context) error {
    var newPayment PaymentRequest
    if err := c.Bind(&newPayment); err != nil {
        return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("invalid input", nil))
    }

    validate := validator.New()
    if err := validate.Struct(newPayment); err != nil {
        return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("validation failed: "+err.Error(), nil))
    }

    order, err := ph.paymentService.GetOrderByID(newPayment.OrderID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Failed to get product", nil))
    }

    user, err := ph.paymentService.GetUserByID(order.UserID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Failed to get product", nil))
    }

    payments := payment.Payment{
        PaymentMethod: newPayment.PaymentMethod,
        OrderID:       newPayment.OrderID,
    }

    createdPayment, err := ph.paymentService.FindOrCreatePayment(newPayment.OrderID, payments)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Failed to create payment", nil))
    }

    client := ph.midtrans
    req := &midtrans.ChargeReq{
        PaymentType: midtrans.SourceBankTransfer,
        BankTransfer: &midtrans.BankTransferDetail{
            Bank: midtrans.BankBca,
        },
        TransactionDetails: midtrans.TransactionDetails{
            OrderID:  order.InvoiceID,
            GrossAmt: int64(order.Price), 
        },
        CustomerDetail: &midtrans.CustDetail{
            Email: user.Email,
            FName: user.FullName,
            LName: "",
            Phone: user.NumberPhone,
        },
        Items: &[]midtrans.ItemDetail{
            {
                ID:    strconv.FormatUint(uint64(order.ProductID), 10),
                Price: int64(order.Price) / int64(order.Quantity),
                Qty:   int32(order.Quantity),
                Name:  order.ProductName,
            },
        },
    }
    coreGateway := midtrans.CoreGateway{
        Client: client,
    }
   
    resp, err := coreGateway.Charge(req)
    log.Printf("Charge Request: %+v\n", resp)
    if err != nil {
        log.Printf("Transaction failed with error: %v", err)
    }

    var SignatureID string
    var VANumber string
    var OrderStatus string
    
    if resp.StatusCode == "201" {
        SignatureID = resp.TransactionID
        VANumber = resp.VANumbers[len(resp.VANumbers)-1].VANumber
        OrderStatus = resp.TransactionStatus
        fmt.Printf("Transaction successful: %+v\n", resp)
    } else {
        fmt.Printf("Transaction failed with status: %s\n", resp.StatusMessage)
    }

    if SignatureID == "" {
        resp, err = coreGateway.Status(order.InvoiceID)
        log.Printf("Check Status Request: %+v\n", resp)
        if err != nil || !helper.ContainsString([]string{"201", "200"}, resp.StatusCode) {
            log.Printf("Check Status Transaction failed with error: %v", err)
            return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Failed to charge payment", nil))
        }

        SignatureID = resp.TransactionID
        VANumber = resp.VANumbers[len(resp.VANumbers)-1].VANumber
        OrderStatus = resp.TransactionStatus
    }

    payments = payment.Payment{
        SignatureID:   SignatureID,
        VANumber:      VANumber,
    }

    createdPayment, err = ph.paymentService.Update(newPayment.OrderID, payments)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Failed to update payment", nil))
    }

    orderAttr := newOrder.Order{
        ID: newPayment.OrderID,
        Status: OrderStatus,
    }

	log.Printf("orderAttr: %+v\n", orderAttr)
    updateOrder, err := ph.orderService.UpdateOrder(newPayment.OrderID, orderAttr)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Failed to update order", nil))
    }
	log.Printf("updateOrder: %+v\n", updateOrder)


    paymentresponse := newOrderResponse.OrderResponse{
        ID: order.ID,
        UserID: order.UserID,
        ProductID: order.ProductID,
        ProductName: order.ProductName,
        ProductPicture: order.ProductPicture,
        Quantity: order.Quantity,
        Price: order.Price,
        Status: updateOrder.Status,
        InvoiceID:     order.InvoiceID,
        Payment: newOrderResponse.PaymentResponse{
            ID:            createdPayment.ID,
            PaymentMethod: createdPayment.PaymentMethod,
            OrderID:       createdPayment.OrderID,
            SignatureID:   createdPayment.SignatureID,
            VANumber:      createdPayment.VANumber,
        },

    }

    return c.JSON(http.StatusCreated, responses.JSONWebResponse("payment created successfully", paymentresponse))
}

