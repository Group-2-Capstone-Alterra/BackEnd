package data

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	TransactionID   uint
	PaymentMethod   string
	PaymentStatus   string
	PaymentAmount   float64
}

