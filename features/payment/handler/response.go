package handler

type PaymentResponse struct {
	ID              uint      `json:"id"`
	OrderID			uint      `json:"order_id"`
	PaymentMethod   string    `json:"payment_method"`
	BillingNumber   string    `json:"billing_number"`
	InvoiceID       string    `json:"invoice_id"`
}
