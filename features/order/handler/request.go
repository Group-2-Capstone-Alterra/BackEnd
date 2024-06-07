package handler

type OrderRequest struct {
	ProductID uint `json:"product_id" form:"product_id"`
	Quantity  int  `json:"quantity" form:"quantity"`
}