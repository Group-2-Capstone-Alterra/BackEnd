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

func (p *productrQuery) Insert(input product.Core) error {
	userGorm := CoreToGorm(input)
	tx := p.db.Create(&userGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (p *productrQuery) SelectAll(userid uint) ([]product.Core, error) {
	var allProduct []Product
	tx := p.db.Where("id_user = ?", userid).Find(&allProduct)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var allProductCore []product.Core
	for _, v := range allProduct {
		allProductCore = append(allProductCore, GormToCore(v))
	}
	return allProductCore, nil
}
