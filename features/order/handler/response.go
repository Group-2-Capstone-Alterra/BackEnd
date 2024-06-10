package handler

type PaymentResponse struct {
	ID            uint   `json:"id"`
	OrderID       uint   `json:"order_id"`
	PaymentMethod string `json:"payment_method"`
	PaymentStatus string `json:"payment_status"`
	SignatureID   string `json:"signature_id"`
	VANumber      string `json:"va_number"`
	InvoiceID     string `json:"invoice_id"`
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
	Payment        PaymentResponse `json:"payment"`
}

type CreatedResponse struct {
	ID uint `json:"id"`
}
