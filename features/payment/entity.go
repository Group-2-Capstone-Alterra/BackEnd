package payment

import "time"

type PaymentCore struct {
	ID              uint
	PaymentMethod   string
	PaymentStatus   string
	PaymentAmount   float64
	TransactionTime time.Time
}


type PaymentModel interface {
	Create(payment PaymentCore) error
}

type PaymentService interface {
	CreatePayment(payment PaymentCore) error
}