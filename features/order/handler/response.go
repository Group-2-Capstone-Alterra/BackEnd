package handler

import "PetPalApp/features/order"

func CoreToResponse(order order.Order) OrderResponse {
    var payment PaymentResponse
    if order.Payment.ID != 0 {
        payment = PaymentResponse{
            ID:             order.Payment.ID,
            OrderID:        order.Payment.OrderID,
            PaymentMethod:  order.Payment.PaymentMethod,
            PaymentStatus:  order.Payment.PaymentStatus,
			SignatureID:    order.Payment.SignatureID,
			VANumber:       order.Payment.VANumber,
			InvoiceID: 		order.Payment.InvoiceID,
        }
    }

    return OrderResponse{
        ID:             order.ID,
        UserID:         order.UserID,
        ProductID:      order.ProductID,
		ProductName:    order.ProductName,
		ProductPicture: order.ProductPicture,
		Quantity:       order.Quantity,
		Price:          order.Price,
        Status:         order.Status,
        Payment:        payment,
    }
}


type PaymentResponse struct {
	ID            uint		`json:"id"`
	OrderID       uint		`json:"order_id"`
	PaymentMethod string	`json:"payment_method"`
	PaymentStatus string	`json:"payment_status"`
	SignatureID   string	`json:"signature_id"`
	VANumber      string	`json:"va_number"`
	InvoiceID     string	`json:"invoice_id"`
}

type OrderResponse struct {
	ID             uint						`json:"id"`
	UserID         uint         			`json:"user_id"`
	ProductID      uint						`json:"product_id"`
	ProductName    string					`json:"product_name"`
	ProductPicture string					`json:"product_picture"`
	Quantity       uint						`json:"quantity"`
	Price          float64					`json:"price"`
	Status         string					`json:"status"`
	Payment   	   PaymentResponse			`json:"payment"`
}

type CreatedResponse struct {
	ID             uint						`json:"id"`
}