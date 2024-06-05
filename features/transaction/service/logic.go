package service

import (
	"PetPalApp/features/transaction"
	"errors"
)

type TransactionService struct {
	TransactionModel transaction.TransactionModel
}

func New(tm transaction.TransactionModel) transaction.TransactionService {
	return &TransactionService{
		TransactionModel: tm,
	}
}

func (ts *TransactionService) CreateTransaction(transaction transaction.TransactionCore) error {
	if transaction.UserID == 0 || transaction.Amount <= 0 {
		return errors.New("invalid transaction data")
	}
	return ts.TransactionModel.Create(transaction)
}

func (ts *TransactionService) GetTransactionsByUserID(userID uint) ([]transaction.TransactionCore, error) {
	return ts.TransactionModel.GetByUserID(userID)
}
