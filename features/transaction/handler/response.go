package handler

type TransactionResponse struct {
	ID        uint    `json:"id"`
	UserID    uint    `json:"user_id"`
	Amount    float64 `json:"amount"`
	Status    string  `json:"status"`
	CreatedAt string  `json:"created_at"`
}