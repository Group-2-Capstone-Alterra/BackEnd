package data

import (
	payment "PetPalApp/features/payment/data"
	"PetPalApp/features/transaction"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	OrderID	   uint
	UserID     uint
	Amount     float64
	Status     string
	Payments   []payment.Payment `gorm:"foreign_key:TransactionID"`
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
