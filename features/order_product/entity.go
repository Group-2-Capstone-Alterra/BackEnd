package order_product

import "time"

type OrderProductCore struct {
	ID        uint
	OrderID   uint
	ProductID uint
	Quantity  uint
	Price     float64
	CreatedAt time.Time
}

type OrderProductModel interface {
	CreateOrderProduct(opCore OrderProductCore) error
}

type OrderProductService interface {
	CreateOrderProduct(opCore OrderProductCore) error
}
