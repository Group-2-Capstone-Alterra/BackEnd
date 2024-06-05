package transaction

import (
	"time"
)

type TransactionCore struct {
	ID        uint
	UserID    uint
	Amount    float64
	Status    string
	CreatedAt time.Time
}

type TransactionModel interface {
	Create(transaction TransactionCore) error
	GetByUserID(userID uint) ([]TransactionCore, error)
}

type TransactionService interface {
	CreateTransaction(transaction TransactionCore) error
	GetTransactionsByUserID(userID uint) ([]TransactionCore, error)
}