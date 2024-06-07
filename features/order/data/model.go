package data

import (
	order "PetPalApp/features/order"
	transaction "PetPalApp/features/transaction/data"

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
	Transactions    []transaction.Transaction   `gorm:"foreign_key:OrderID"`
}


func ToCore(o Order) order.OrderCore {
    var transactionCore []order.Transaction
    for _, op := range o.Transactions {
        transaction := order.Transaction{
            ID:        op.ID,
            UserID:    op.UserID,
            Amount:    op.Amount,
            Status:    op.Status,
        }
        transactionCore = append(transactionCore, transaction)
    }

    return order.OrderCore{
        ID:             o.ID,
        UserID:         o.UserID,
        ProductID:      o.ProductID,
		ProductName:    o.ProductName,
		ProductPicture: o.ProductPicture,
		Quantity:       o.Quantity,
		Price:          o.Price,
        Status:         o.Status,
        Transactions:   transactionCore,
    }
}


