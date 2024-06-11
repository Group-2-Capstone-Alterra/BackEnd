package handler

type MidtransNotificationRequest struct {
	TransactionTime   string `json:"transaction_time" form:"transaction_time"`
	TransactionStatus string `json:"transaction_status" form:"transaction_status"`
	TransactionId     string `json:"transaction_id" form:"transaction_id"`
	StatusMessage     string `json:"status_message" form:"status_message"`
	StatusCode        string `json:"status_code" form:"status_code"`
	SignatureKey      string `json:"signature_key" form:"signature_key"`
	SettlementTime    string `json:"settlement_time" form:"settlement_time"`
	PaymentType       string `json:"payment_type" form:"payment_type"`
	OrderId           string `json:"order_id" form:"order_id"`
	MerchantId        string `json:"merchant_id" form:"merchant_id"`
	GrossAmount       string `json:"gross_amount" form:"gross_amount"`
	FraudStatus       string `json:"fraud_status" form:"fraud_status"`
	Currency          string `json:"currency" form:"currency"`
}
