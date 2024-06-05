package data

import (
	"PetPalApp/features/transaction"

	"gorm.io/gorm"
)

type TransactionModel struct {
	db *gorm.DB
}

func New(db *gorm.DB) transaction.TransactionModel {
	return &TransactionModel{
		db: db,
	}
}

func (tm *TransactionModel) Create(transaction transaction.TransactionCore) error {
	transactionGorm := Transaction{
		UserID: transaction.UserID,
		Amount: transaction.Amount,
		Status: transaction.Status,
	}
	tx := tm.db.Create(&transactionGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

