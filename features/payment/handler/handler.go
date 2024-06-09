package handler

import (
	"PetPalApp/features/payment"
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
	paymentService payment.PaymentService
	midtrans       midtrans.Client
}

func New(ps payment.PaymentService, midtrans midtrans.Client) *PaymentHandler {
	return &PaymentHandler{
		paymentService: ps,
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

	order, err := ph.paymentService.GetOrderById(newPayment.OrderID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Failed to get product", nil))
    }

	user, err := ph.paymentService.GetUserById(order.UserID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Failed to get product", nil))
    }

	client := ph.midtrans
    req := &midtrans.ChargeReq{
        PaymentType: midtrans.SourceBankTransfer,
        BankTransfer: &midtrans.BankTransferDetail{
            Bank: midtrans.BankBca,
        },
		TransactionDetails: midtrans.TransactionDetails{
            OrderID:  order.InvoiceID,
            GrossAmt: int64(order.Price), // Jumlah transaksi
        },
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.FullName,
			LName: "",
			Phone: user.NumberPhone,
		},
        Items: &[]midtrans.ItemDetail{
            {
                ID: strconv.FormatUint(uint64(order.ProductID), 10),
                Price: int64(order.Price) / int64(order.Quantity),
                Qty: int32(order.Quantity),
                Name: order.ProductName,
            },
        },
    }
	coreGateway := midtrans.CoreGateway{
        Client: client,
    }
    // Lakukan transaksi
    resp, err := coreGateway.Charge(req)
    if err != nil {
        log.Fatalf("Transaction failed with error: %v", err)
    }
	log.Printf("Charge Request: %+v\n", req)

	var SignatureID string
	var BillingNumber string;
    // Cek hasil transaksi
    if resp.StatusCode == "201" {
		SignatureID = resp.TransactionID
		BillingNumber = resp.VANumbers[len(resp.VANumbers)-1].VANumber
        fmt.Printf("Transaction successful: %+v\n", resp)
    } else {
        fmt.Printf("Transaction failed with status: %s\n", resp.StatusMessage)
    }

	payments := payment.PaymentCore{
		PaymentMethod: newPayment.PaymentMethod,
		OrderID: newPayment.OrderID,
		InvoiceID: order.InvoiceID,
		SignatureID: SignatureID,
		BillingNumber: BillingNumber,
	}

	if SignatureID == "" {
        return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Failed to charge payment", nil))
	}

    err = ph.paymentService.CreatePayment(payments)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Failed to create payment", nil))
    }

	// Update order status
	// err = ph.paymentService.UpdateOrderStatus(newPayment.OrderID, "Paid")
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Failed to update order status", nil))
	// }

	return c.JSON(http.StatusCreated, responses.JSONWebResponse("payment created successfully", payments))
}
