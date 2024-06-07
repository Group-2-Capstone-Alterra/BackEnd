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
	Transactions   []Transaction
}

type Product struct {
	ID             uint
	ProductName    string
	ProductPicture string
	Price          float64
}

type Transaction struct {
	ID     uint
	UserID uint
	Amount float64
	Status string
}

type OrderModel interface {
	CreateOrder(opCore OrderCore) error
	GetOrdersByUserID(userID uint) ([]OrderCore, error)
	GetPriceByProductID(productID uint) (*Product, error)
}

type OrderService interface {
	CreateOrder(opCore OrderCore) error
	GetOrdersByUserID(userID uint) ([]OrderCore, error)
	GetProductById(id uint) (data *Product, err error)
}
