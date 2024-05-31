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

func (p *productrQuery) SelectAll(offset uint, sortStr string) ([]product.Core, error) {
	var allProduct []Product
	if sortStr == "lowest" {
		tx := p.db.Order("price asc").Limit(10).Offset(int(offset)).Find(&allProduct)
		if tx.Error != nil {
			return nil, tx.Error
		}
	} else if sortStr == "highest" {
		tx := p.db.Order("price desc").Limit(10).Offset(int(offset)).Find(&allProduct)
		if tx.Error != nil {
			return nil, tx.Error
		}
	} else {
		tx := p.db.Limit(10).Offset(int(offset)).Find(&allProduct)
		if tx.Error != nil {
			return nil, tx.Error
		}
	}
	var allProductCore []product.Core
	for _, v := range allProduct {
		allProductCore = append(allProductCore, GormToCore(v))
	}
	return allProductCore, nil
}

func (p *productrQuery) SelectAllAdmin(userid uint, offset uint) ([]product.Core, error) {
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

func (p *productrQuery) SelectById(id uint) (*product.Core, error) {
	var productData Product
	tx := p.db.First(&productData, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	projectcore := GormToCore(productData)
	return &projectcore, nil
}

func (p *productrQuery) SelectByIdAdmin(id uint, userid uint) (*product.Core, error) {
	var productData Product
	tx := p.db.Where("id_user = ?", userid).First(&productData, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	projectcore := GormToCore(productData)
	return &projectcore, nil
}

func (p *productrQuery) PutById(id uint, userid uint, input product.Core) error {

	userGorm := CoreToGorm(input)

	tx := p.db.Model(&Product{}).Where("id = ? AND id_user = ?", id, userid).Updates(&userGorm)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (p *productrQuery) Delete(id uint, userid uint) error {
	tx := p.db.Where("id_user = ?", userid).Delete(&Product{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
