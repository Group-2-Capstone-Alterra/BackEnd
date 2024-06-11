package order

type Order struct {
	ID             uint
	UserID         uint
	ProductID      uint
	ProductName    string
	ProductPicture string
	Quantity       uint
	Price          float64
	Status         string
	InvoiceID      string
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
	SignatureID   string
	VANumber      string
	InvoiceID     string
}

type OrderModel interface {
	CreateOrder(order Order) (Order, error)
	GetOrdersByUserID(userID uint) ([]Order, error)
	GetProductByID(productID uint) (*Product, error)
	GetOrderByID(orderID uint) (*Order, error)
	UpdateOrder(orderID uint, order Order) (Order, error)
	GetOrderByInvoiceID(invoiceID string) (*Order, error)
}

type OrderService interface {
	CreateOrder(order Order) (Order, error)
	GetOrdersByUserID(userID uint) ([]Order, error)
	GetProductByID(id uint) (data *Product, err error)
	GetOrderByID(orderID uint) (*Order, error)
	UpdateOrder(orderID uint, order Order) (Order, error)
	GetOrderByInvoiceID(invoiceID string) (*Order, error)
}
