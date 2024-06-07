package data

import (
	"PetPalApp/features/transaction"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	OrderID	   uint
	UserID     uint
	Amount     float64
	Status     string
}

func (t *Transaction) ToCore() transaction.TransactionCore {
	return transaction.TransactionCore{
		ID:        t.ID,
		UserID:    t.UserID,
		Amount:    t.Amount,
		Status:    t.Status,
		CreatedAt: t.CreatedAt,
	}
}
