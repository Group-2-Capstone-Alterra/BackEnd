package handler

import "PetPalApp/features/payment"

type PaymentRequest struct {
	OrderID	      uint	  `json:"order_id" form:"order_id"`
	PaymentMethod string  `json:"payment_method" form:"payment_method" validate:"required"`
}

func (req *PaymentRequest) ToCore() payment.PaymentCore {
	return payment.PaymentCore{
		OrderID:       req.OrderID,
		PaymentMethod: req.PaymentMethod,
	}
}
