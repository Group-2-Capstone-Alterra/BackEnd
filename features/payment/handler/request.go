package handler

import "PetPalApp/features/payment"

type PaymentRequest struct {
	PaymentMethod string  `json:"payment_method" form:"payment_method" validate:"required"`
	PaymentAmount float64 `json:"payment_amount" form:"payment_method" validate:"required"`
}

func (req *PaymentRequest) ToCore() payment.PaymentCore {
	return payment.PaymentCore{
		PaymentMethod: req.PaymentMethod,
		PaymentAmount: req.PaymentAmount,
	}
}
