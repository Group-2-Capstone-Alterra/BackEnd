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
		Total:     orderCore.Total,
		Status:    orderCore.Status,
	}
	tx := om.db.Create(&orderGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (om *OrderModel) GetOrdersByUserID(userID uint) ([]order.OrderCore, error) {
    var orders []Order
    tx := om.db.Where("user_id = ?", userID).Find(&orders)
    if tx.Error != nil {
        return nil, tx.Error
    }

    var result []order.OrderCore
    for _, order := range orders {
        result = append(result, order.ToCore())
    }

    return result, nil
}