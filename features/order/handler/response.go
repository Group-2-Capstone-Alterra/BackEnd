package handler

import "PetPalApp/features/order"

func CoreToResponse(order order.OrderCore) OrderResponse {
	var transactions []TransactionResponse
	for _, op := range order.Transactions { 
		transactions = append(transactions, TransactionResponse{
			ID:   		 op.ID,
			UserID: 	 op.UserID,
			Amount:      op.Amount,
			Status:      op.Status,
		})
	}

	return OrderResponse{
		ID:             order.ID,
		UserID:         order.UserID,
		ProductID:      order.ProductID,
		ProductName:    order.ProductName,
		ProductPicture: order.ProductPicture,
		Quantity:       order.Quantity,
		Price: 			order.Price,
		Status:         order.Status,
		Transactions:   transactions,
	}
}


type TransactionResponse struct {
	ID     uint		`json:"id"`
	UserID uint		`json:"user_id"`
	Amount float64  `json:"amount"`
	Status string   `json:"status"`
}

type OrderResponse struct {
	ID             uint						`json:"id"`
	UserID         uint         			`json:"user_id"`
	ProductID      uint						`json:"product_id"`
	ProductName    string					`json:"product_name"`
	ProductPicture string					`json:"product_picture"`
	Quantity       uint						`json:"quantity"`
	Price          float64					`json:"price"`
	Status         string					`json:"status"`
	Transactions   []TransactionResponse	`json:"transactions"`
}