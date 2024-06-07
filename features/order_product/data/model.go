package data

import (
	"PetPalApp/features/order_product"

	"gorm.io/gorm"
)

type OrderProduct struct {
	gorm.Model
	OrderID    uint
	ProductID  uint
	Quantity   uint
	Price      float64
}

func ToCore(op OrderProduct) order_product.OrderProductCore {
	return order_product.OrderProductCore{
		ID:        op.ID,
		OrderID:   op.OrderID,
		ProductID: op.ProductID,
		Quantity:  op.Quantity,
		Price:     op.Price,
		CreatedAt: op.CreatedAt,
	}
}


