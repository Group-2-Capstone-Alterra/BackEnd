package data

import (
	"PetPalApp/features/order"
	"fmt"
	"log"

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

func (om *OrderModel) CreateOrder(order order.Order) (order.Order, error) {
    tx := om.db.Create(&order)
    if tx.Error != nil {
        return order, tx.Error
    }
    return order, nil
}

func (om *OrderModel) UpdateOrder(orderID uint, updatedOrder order.Order) (order.Order, error) {
	log.Printf("orderID: %+v\n", orderID)

	var order order.Order
    if err := om.db.Where("id = ?", orderID).First(&order).Error; err != nil {
        return order, fmt.Errorf("failed to find payment with OrderID %d: %v", orderID, err)
    }

    order.Status = updatedOrder.Status

    if err := om.db.Save(&order).Error; err != nil {
        return order, fmt.Errorf("failed to update payment: %v", err)
    }

    return order, nil
}

func (om *OrderModel) GetOrdersByUserID(userID uint) ([]order.Order, error) {
    var result []order.Order
    if err := om.db.Preload("Payment").Where("user_id = ?", userID).Find(&result).Error; err != nil {
        return nil, err
    }

    return result, nil
}

func (om *OrderModel) GetProductByID(productID uint) (*order.Product, error) {
	var result order.Product
	if err := om.db.Where("id = ?", productID).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil 
}

func (om *OrderModel) GetOrderByID(orderID uint) (*order.Order, error) {
	var result order.Order
	if err := om.db.Preload("Payment").Where("id = ?", orderID).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil 
}

func (om *OrderModel) GetOrderByInvoiceID(invoiceID string) (*order.Order, error) {
	var result order.Order
	if err := om.db.Where("invoice_id = ?", invoiceID).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil 
}
