package data

import (
	"PetPalApp/features/order"
	"PetPalApp/features/payment"
	"fmt"

	"gorm.io/gorm"
)

type PaymentModel struct {
	db *gorm.DB
}

func New(db *gorm.DB) payment.PaymentModel {
	return &PaymentModel{db: db}
}

func (pm *PaymentModel) FindOrCreatePayment(orderID uint, payment payment.Payment) (payment.Payment, error) {
    if err := pm.db.Where("order_id = ?", orderID).First(&payment).Error; err != nil {
		tx := pm.db.Create(&payment)
		if tx.Error != nil {
			return payment, tx.Error
		}
    }

    return payment, nil
}

func (pm *PaymentModel) Update(orderID uint, updatedPayment payment.Payment) (payment.Payment, error) {
    var payment payment.Payment
    if err := pm.db.Where("order_id = ?", orderID).First(&payment).Error; err != nil {
        return payment, fmt.Errorf("failed to find payment with OrderID %d: %v", orderID, err)
    }

    payment.SignatureID = updatedPayment.SignatureID
    payment.VANumber = updatedPayment.VANumber

    if err := pm.db.Save(&payment).Error; err != nil {
        return payment, fmt.Errorf("failed to update payment: %v", err)
    }

    return payment, nil
}

func (pm *PaymentModel) GetOrderByID(orderID uint) (*order.Order, error) {
	var result order.Order
	if err := pm.db.Where("id = ?", orderID).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil 
}

func (pm *PaymentModel) GetUserByID(userID uint) (*payment.User, error) {
	var result payment.User
	if err := pm.db.Where("id = ?", userID).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil 
}
