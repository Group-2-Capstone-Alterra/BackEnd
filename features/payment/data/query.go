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

func (pm *PaymentModel) Create(payment payment.PaymentCore) error {
	paymentGorm := Payment{
		OrderID:         payment.OrderID,
		PaymentMethod:   payment.PaymentMethod,
		SignatureID: 	 payment.SignatureID,
		BillingNumber:   payment.BillingNumber,
	}
	tx := pm.db.Create(&paymentGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (pm *PaymentModel) GetPaymentByID(id uint) (payment.PaymentCore, error) {
	var paymentGorm Payment
	tx := pm.db.First(&paymentGorm, id)
	if tx.Error != nil {
		return payment.PaymentCore{}, tx.Error
	}

	return payment.PaymentCore{
		ID:              paymentGorm.ID,
		PaymentMethod:   paymentGorm.PaymentMethod,
		SignatureID: 	 paymentGorm.SignatureID,
		BillingNumber:   paymentGorm.BillingNumber,
	}, nil
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
