package data

import (
	"PetPalApp/features/payment"
	"time"

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
		PaymentMethod:   payment.PaymentMethod,
		PaymentStatus:   payment.PaymentStatus,
		PaymentAmount:   payment.PaymentAmount,
		TransactionTime: time.Now(),
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
		PaymentStatus:   paymentGorm.PaymentStatus,
		PaymentAmount:   paymentGorm.PaymentAmount,
		TransactionTime: paymentGorm.TransactionTime,
	}, nil
}
