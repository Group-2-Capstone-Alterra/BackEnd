package handler

type TransactionRequest struct {
	UserID uint    `json:"user_id" form:"user_id"`
	Amount float64 `json:"amount" form:"user_id"`
}