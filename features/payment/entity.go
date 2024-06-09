package payment

type Payment struct {
	ID            uint
	OrderID       uint
	PaymentMethod string
	PaymentStatus string
	SignatureID   string
	VANumber      string
	InvoiceID     string
}

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
}

type User struct {
	ID             uint
	FullName       string
	Email          string
	NumberPhone    string
	Address        string
	Password       string
	ProfilePicture string
	Coordinate     string
	Token          string
	Role           string
}

type PaymentModel interface {
	Create(payment Payment) (Payment, error)
	GetPaymentByID(paymentID uint) (*Payment, error)
	GetOrderByID(orderID uint) (*Order, error)
	GetUserByID(orderID uint) (*User, error)
}

type PaymentService interface {
	CreatePayment(payment Payment) (Payment, error)
	GetPaymentByID(id uint) (data *Payment, err error)
	GetOrderByID(id uint) (data *Order, err error)
	GetUserByID(id uint) (*User, error)
}