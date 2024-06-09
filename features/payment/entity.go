package payment

type PaymentCore struct {
	ID            uint
	OrderID       uint
	PaymentMethod string
	SignatureID   string
	BillingNumber string
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
	Create(payment PaymentCore) error
	GetOrderByID(orderID uint) (*Order, error)
	GetUserByID(orderID uint) (*User, error)
}

type PaymentService interface {
	CreatePayment(payment PaymentCore) error
	GetOrderById(id uint) (data *Order, err error)
	GetUserById(id uint) (*User, error)
}