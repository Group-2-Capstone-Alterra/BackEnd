package handler

import (
	"PetPalApp/features/order"
	"PetPalApp/features/product"
	"PetPalApp/features/user"

	_productHandler "PetPalApp/features/product/handler"
	_userHandler "PetPalApp/features/user/handler"
)

type OrderResponse struct {
	ID uint `json:"id"`
	// UserID    uint    `json:"user_id"`
	// ProductID uint    `json:"product_id"`
	User     _userHandler.OrderResponse
	Product  _productHandler.OrderResponse
	Quantity int     `json:"quantity"`
	Total    float64 `json:"total"`
	Status   string  `json:"status"`
}

func CoreToResponse(order order.OrderCore, user user.Core, product product.Core) OrderResponse {
	return OrderResponse{
		ID:       order.ID,
		User:     _userHandler.OrderCoreToResponse(user),
		Product:  _productHandler.OrderCoreToResponse(product),
		Quantity: order.Quantity,
		Total:    order.Total,
		Status:   order.Status,
	}
}
