package data

import (
	"PetPalApp/features/product"

	"gorm.io/gorm"
)

type productrQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) product.DataInterface {
	return &productrQuery{
		db: db,
	}
}

func (u *productrQuery) Insert(input product.Core) error {
	userGorm := CoreToGorm(input)
	tx := u.db.Create(&userGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
