package handler

type OrderResponse struct {
	ProductID uint    `json:"product_id" form:"product_id"`
	Quantity  int     `json:"quantity" form:"quantity"`
	Total     float64 `json:"total" form:"total"`
}
