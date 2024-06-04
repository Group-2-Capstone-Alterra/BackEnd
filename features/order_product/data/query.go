package data

import (
	"PetPalApp/features/order_product"

	"gorm.io/gorm"
)

type OrderProductModel struct {
	db *gorm.DB
}

func New(db *gorm.DB) order_product.OrderProductModel {
	return &OrderProductModel{
		db: db,
	}
}

func (opm *OrderProductModel) CreateOrderProduct(opCore order_product.OrderProductCore) error {
	op := OrderProduct{
		OrderID:   opCore.OrderID,
		ProductID: opCore.ProductID,
		Quantity:  opCore.Quantity,
		Price:     opCore.Price,
	}
	tx := opm.db.Create(&op)
	return tx.Error
}
