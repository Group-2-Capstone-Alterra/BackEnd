package handler

type OrderProductRequest struct {
	OrderID   uint    `json:"order_id" form:"order_id"`
	ProductID uint    `json:"product_id" form:"product_id"`
	Quantity  int     `json:"quantity" form:"quantity"`
	Price     float64 `json:"price" form:"price"`
}