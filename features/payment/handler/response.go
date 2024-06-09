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
