package data

import (
	"time"
)

type Payment struct {
	ID              uint `gorm:"primaryKey"`
	PaymentMethod   string
	PaymentStatus   string
	PaymentAmount   float64
	TransactionTime time.Time
}

