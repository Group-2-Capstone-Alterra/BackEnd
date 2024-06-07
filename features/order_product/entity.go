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

type Product struct {
	ID		  uint
	Price     float64
}

type OrderProductModel interface {
	CreateOrderProduct(opCore OrderProductCore) error
	GetOrderProductsByOrderID(orderID uint) ([]OrderProductCore, error)
	GetPriceByProductID(productID uint) (*Product, error)
}

type OrderProductService interface {
	CreateOrderProduct(opCore OrderProductCore) error
	GetOrderProductsByOrderID(orderID uint) ([]OrderProductCore, error)
	GetProductById(id uint) (data *Product, err error)
}
