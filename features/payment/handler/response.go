package handler

import (
	"PetPalApp/features/payment"
	"time"
)

type PaymentResponse struct {
	ID              uint      `json:"id"`
	PaymentMethod   string    `json:"payment_method"`
	PaymentStatus   string    `json:"payment_status"`
	PaymentAmount   float64   `json:"payment_amount"`
	TransactionTime time.Time `json:"transaction_time"`
}

func FromCore(core payment.PaymentCore) PaymentResponse {
	return PaymentResponse{
		ID:              core.ID,
		PaymentMethod:   core.PaymentMethod,
		PaymentStatus:   core.PaymentStatus,
		PaymentAmount:   core.PaymentAmount,
		TransactionTime: core.TransactionTime,
	}
}
