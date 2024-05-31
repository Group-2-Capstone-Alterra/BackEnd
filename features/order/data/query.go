package data

import (
	"PetPalApp/features/order"

	"gorm.io/gorm"
)

type OrderModel struct {
	db *gorm.DB
}

func New(db *gorm.DB) order.OrderModel {
	return &OrderModel{
		db: db,
	}
}

func (om *OrderModel) CreateOrder(orderCore order.OrderCore) error {
	orderGorm := Order{
		UserID:    orderCore.UserID,
		ProductID: orderCore.ProductID,
		Quantity:  orderCore.Quantity,
		Total:     orderCore.Total,
		Status:    orderCore.Status,
	}
	tx := om.db.Create(&orderGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}