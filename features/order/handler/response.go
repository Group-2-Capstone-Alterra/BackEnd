package handler

type PaymentResponse struct {
	ID            uint   `json:"id"`
	OrderID       uint   `json:"order_id"`
	PaymentMethod string `json:"payment_method"`
	SignatureID   string `json:"signature_id"`
	VANumber      string `json:"va_number"`
}

type OrderResponse struct {
	ID             uint            `json:"id"`
	UserID         uint            `json:"user_id"`
	ProductID      uint            `json:"product_id"`
	ProductName    string          `json:"product_name"`
	ProductPicture string          `json:"product_picture"`
	Quantity       uint            `json:"quantity"`
	Price          float64         `json:"price"`
	Status         string          `json:"status"`
	InvoiceID      string          `json:"invoice_id"`
	Payment        PaymentResponse `json:"payment"`
}