package handler

type OrderCreateRequest struct {
	ProductID uint `json:"product_id" form:"product_id"`
	Quantity  int  `json:"quantity" form:"quantity"`
}

type OrderUpdateStatusRequest struct {
	ID uint `json:"id" form:"id"`
}