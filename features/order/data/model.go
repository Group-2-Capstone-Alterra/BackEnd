package data

import (
	"PetPalApp/features/order"
	orderProduct "PetPalApp/features/order_product/data"
	transaction "PetPalApp/features/transaction/data"

	"gorm.io/gorm"
)

type Order struct {
    gorm.Model
    UserID          uint     
    Total           float64 
    Status          string  
    OrderProducts   []orderProduct.OrderProduct `gorm:"foreign_key:OrderID"`
    Transactions    []transaction.Transaction   `gorm:"foreign_key:OrderID"`
}

func (o *Order) ToCore() order.OrderCore {
    return order.OrderCore{
        ID:        o.ID,
        UserID:    o.UserID,
        Total:     o.Total,
        Status:    o.Status,
        CreatedAt: o.CreatedAt,
    }
}


