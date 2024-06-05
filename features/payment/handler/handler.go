package handler

import (
	"PetPalApp/features/payment"
	"PetPalApp/utils/responses"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type PaymentHandler struct {
	paymentService payment.PaymentService
}

func New(ps payment.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		paymentService: ps,
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

	paymentCore := newPayment.ToCore()
	if err := ph.paymentService.CreatePayment(paymentCore); err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("failed to create payment", nil))
	}

	paymentResponse := FromCore(paymentCore)
	return c.JSON(http.StatusCreated, responses.JSONWebResponse("payment created successfully", paymentResponse))
}
