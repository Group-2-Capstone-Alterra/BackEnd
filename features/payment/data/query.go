package data

import (
	"PetPalApp/features/payment"

	"gorm.io/gorm"
)

type PaymentModel struct {
	db *gorm.DB
}

func New(db *gorm.DB) payment.PaymentModel {
	return &PaymentModel{db: db}
}

func (pm *PaymentModel) Create(payment payment.Payment) (payment.Payment, error) {
    tx := pm.db.Create(&payment)
    if tx.Error != nil {
        return payment, tx.Error
    }
    return payment, nil
}


func (pm *PaymentModel) GetPaymentByID(paymentID uint) (*payment.Payment, error) {
	var result payment.Payment
	if err := pm.db.Where("id = ?", paymentID).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil 
}

func (pm *PaymentModel) GetOrderByID(orderID uint) (*payment.Order, error) {
	var result payment.Order
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
