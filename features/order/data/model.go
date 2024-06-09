package data

import (
	order "PetPalApp/features/order"
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

func ToCore(o Order) order.Order {
    var paymentCore order.Payment
    if o.Payment.ID != 0 {
        paymentCore = order.Payment{
            ID:             o.Payment.ID,
            OrderID:        o.Payment.OrderID,
            PaymentMethod:  o.Payment.PaymentMethod,
			PaymentStatus:  o.Payment.PaymentStatus,
			SignatureID:    o.Payment.SignatureID,
			VANumber:       o.Payment.VANumber,
			InvoiceID: 		o.Payment.InvoiceID,	
        }
    }

    return order.Order{
        ID:             o.ID,
        UserID:         o.UserID,
        ProductID:      o.ProductID,
		ProductName:    o.ProductName,
		ProductPicture: o.ProductPicture,
		Quantity:       o.Quantity,
		Price:          o.Price,
        Status:         o.Status,
        Payment:       paymentCore,
    }
}


