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

func (tm *TransactionModel) GetByUserID(userID uint) ([]transaction.TransactionCore, error) {
	var transactions []Transaction
	tx := tm.db.Where("user_id = ?", userID).Find(&transactions)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var result []transaction.TransactionCore
	for _, t := range transactions {
		result = append(result, t.ToCore())
	}
	return result, nil
}
