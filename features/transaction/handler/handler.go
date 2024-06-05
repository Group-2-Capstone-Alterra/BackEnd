package handler

import (
	"PetPalApp/features/transaction"
	"PetPalApp/utils/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	transactionService transaction.TransactionService
}

func New(ts transaction.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		transactionService: ts,
	}
}

func (th *TransactionHandler) CreateTransaction(c echo.Context) error {
	newTransaction := TransactionRequest{}
	errBind := c.Bind(&newTransaction)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error binding data: "+errBind.Error(), nil))
	}

	transaction := transaction.TransactionCore{
		UserID: newTransaction.UserID,
		Amount: newTransaction.Amount,
		Status: "Pending",
	}

	errCreate := th.transactionService.CreateTransaction(transaction)
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error creating transaction: "+errCreate.Error(), nil))
	}

	return c.JSON(http.StatusCreated, responses.JSONWebResponse("transaction created successfully", nil))
}
