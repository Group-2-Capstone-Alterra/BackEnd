package handler

import "PetPalApp/features/order"

func CoreToResponse(order order.OrderCore) OrderResponse {
    var payment PaymentResponse
    if order.Payment.ID != 0 {
        payment = PaymentResponse{
            ID:             order.Payment.ID,
            OrderID:        order.Payment.OrderID,
            PaymentMethod:  order.Payment.PaymentMethod,
            PaymentStatus:  order.Payment.PaymentStatus,
            PaymentAmount:  order.Payment.PaymentAmount,
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
	ID              uint		`json:"id"`
	OrderID         uint		`json:"order_id"`
	PaymentMethod   string		`json:"payment_method"`
	PaymentStatus   string		`json:"payment_status"`
	PaymentAmount   float64		`json:"payment_amount"`
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