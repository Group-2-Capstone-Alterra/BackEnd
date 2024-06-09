package data

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	OrderID   		uint
	PaymentMethod   string
	SignatureID	    string
	BillingNumber   string
}

