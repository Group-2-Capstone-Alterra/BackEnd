package handler

type OrderRequest struct {
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Total     float64 `json:"total"`
}
