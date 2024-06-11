package handler

type PaymentRequest struct {
	OrderID	      uint	  `json:"order_id" form:"order_id"`
	PaymentMethod string  `json:"payment_method" form:"payment_method"`
}