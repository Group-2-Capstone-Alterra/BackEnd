package order

type OrderCore struct {
	ID             uint
	UserID         uint
	ProductID      uint
	ProductName    string
	ProductPicture string
	Quantity       uint
	Price          float64
	Status         string
	Payment        Payment
}

type Product struct {
	ID             uint
	ProductName    string
	ProductPicture string
	Price          float64
}

type Payment struct {
	ID            uint
	OrderID       uint
	PaymentMethod string
	PaymentStatus string
	PaymentAmount float64
}

type OrderModel interface {
	CreateOrder(opCore OrderCore) error
	GetOrdersByUserID(userID uint) ([]OrderCore, error)
	GetProductById(productID uint) (*Product, error)
}

type OrderService interface {
	CreateOrder(opCore OrderCore) error
	GetOrdersByUserID(userID uint) ([]OrderCore, error)
	GetProductById(id uint) (data *Product, err error)
}
