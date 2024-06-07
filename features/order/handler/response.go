package handler

import (
	"PetPalApp/features/order"
	"PetPalApp/features/product"
	"PetPalApp/features/user"
)

type OrderResponse struct {
	ID 		 uint 	 `json:"id"`
	Quantity int     `json:"quantity"`
	Total    float64 `json:"total"`
	Status   string  `json:"status"`
}

func CoreToResponse(order order.OrderCore, user user.Core, product product.Core) OrderResponse {
	return OrderResponse{
		ID:       order.ID,
		Total:    order.Total,
		Status:   order.Status,
	}
}
