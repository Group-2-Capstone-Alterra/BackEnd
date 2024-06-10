package data

import (
	payment "PetPalApp/features/payment/data"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID          uint
	ProductID  		uint
	ProductName 	string
	ProductPicture 	string
	Quantity   		uint
	Price      		float64
	Status			string
    InvoiceID       string
	Payment         payment.Payment   `gorm:"foreign_key:OrderID"`
}


