package data

import (
	"PetPalApp/features/order"

	"gorm.io/gorm"
)

type Order struct {
    gorm.Model
    UserID    uint   
    ProductID uint   
    Quantity  int     
    Total     float64 
    Status    string  
}

func (o *Order) ToCore() order.OrderCore {
    return order.OrderCore{
        ID:        o.ID,
        UserID:    o.UserID,
        ProductID: o.ProductID,
        Quantity:  o.Quantity,
        Total:     o.Total,
        Status:    o.Status,
        CreatedAt: o.CreatedAt,
    }
}


