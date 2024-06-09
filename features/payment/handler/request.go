package handler

import "PetPalApp/features/payment"

type PaymentRequest struct {
	OrderID	      uint	  `json:"order_id" form:"order_id"`
	PaymentMethod string  `json:"payment_method" form:"payment_method"`
}

func (req *PaymentRequest) ToCore() payment.Payment {
	return payment.Payment{
		OrderID:       req.OrderID,
		PaymentMethod: req.PaymentMethod,
	}
}
