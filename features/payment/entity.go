package payment

import "PetPalApp/features/order"

type Payment struct {
	ID            uint
	OrderID       uint
	PaymentMethod string
	SignatureID   string
	VANumber      string
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
	FindOrCreatePayment(orderID uint, payment Payment) (Payment, error)
	GetPaymentByID(paymentID uint) (*Payment, error)
	GetOrderByID(orderID uint) (*order.Order, error)
	GetUserByID(orderID uint) (*User, error)
	Update(orderID uint, payment Payment) (Payment, error)
}

type PaymentService interface {
	FindOrCreatePayment(orderID uint, payment Payment) (Payment, error)
	GetPaymentByID(id uint) (data *Payment, err error)
	GetOrderByID(id uint) (data *order.Order, err error)
	GetUserByID(id uint) (*User, error)
	Update(orderID uint, payment Payment) (Payment, error)
}