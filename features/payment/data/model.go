package data

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	OrderID   				uint
	PaymentMethod   		string
	PaymentStatus 			string
	SignatureID	    		string
	VANumber    			string
	InvoiceID     			string
}

