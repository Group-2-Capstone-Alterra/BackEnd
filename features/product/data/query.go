package data

import (
	"PetPalApp/features/product"
	"PetPalApp/utils/helper"

	"gorm.io/gorm"
)

type productrQuery struct {
	db     *gorm.DB
	helper helper.HelperInterface
}

func New(db *gorm.DB, helper helper.HelperInterface) product.ProductModel {
	return &productrQuery{
		db:     db,
		helper: helper,
	}
}

const (
	quserid = "id_user = ?"
)

func (p *productrQuery) Insert(input product.Core) error {
	userGorm := CoreToGorm(input)
	tx := p.db.Create(&userGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (p *productrQuery) SelectAll(limit, offset uint, sortStr string) ([]product.Core, error) {
	var allProduct []Product
	if sortStr == "lowest distance" || sortStr == "lowest" {
		tx := p.db.Order("price asc").Limit(int(limit)).Offset(int(offset)).Find(&allProduct)
		if tx.Error != nil {
			return nil, tx.Error
		}

	} else if sortStr == "highest distance" || sortStr == "highest" {
		tx := p.db.Order("price desc").Limit(int(limit)).Offset(int(offset)).Find(&allProduct)
		if tx.Error != nil {
			return nil, tx.Error
		}
	} else {
		tx := p.db.Limit(int(limit)).Offset(int(offset)).Find(&allProduct)
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

func (p *productrQuery) SelectAllAdmin(limit, userid uint, offset uint) ([]product.Core, error) {
	var allProduct []Product
	tx := p.db.Where(quserid, userid).Find(&allProduct)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var allProductCore []product.Core
	for _, v := range allProduct {
		allProductCore = append(allProductCore, GormToCore(v))
	}
	return allProductCore, nil
}

func (p *productrQuery) SelectAllAdminByName(limit, userid uint, offset uint, name string) ([]product.Core, error) {
	var allProduct []Product
	tx := p.db.Where("product_name LIKE ?", "%"+name+"%").Where(quserid, userid).Find(&allProduct)
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

func (p *productrQuery) SelectByName(limit, offset uint, sortStr, name string) ([]product.Core, error) {
	var allProduct []Product
	if sortStr == "lowest distance" || sortStr == "lowest" {
		tx := p.db.Order("price asc").Limit(int(limit)).Offset(int(offset)).Where("product_name LIKE ?", "%"+name+"%").Find(&allProduct)
		if tx.Error != nil {
			return nil, tx.Error
		}

	} else if sortStr == "highest distance" || sortStr == "highest" {
		tx := p.db.Order("price desc").Limit(int(limit)).Offset(int(offset)).Where("product_name LIKE ?", "%"+name+"%").Find(&allProduct)
		if tx.Error != nil {
			return nil, tx.Error
		}
	} else {
		tx := p.db.Limit(int(limit)).Offset(int(offset)).Where("product_name LIKE ?", "%"+name+"%").Find(&allProduct)
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

func (p *productrQuery) SelectByIdAdmin(id uint, userid uint) (*product.Core, error) {
	var productData Product
	tx := p.db.Where(quserid, userid).First(&productData, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	projectcore := GormToCore(productData)
	return &projectcore, nil
}

func (p *productrQuery) VerIsAdmin(userid uint) (*product.Core, error) {
	var productData Product
	tx := p.db.Where(quserid, userid).Find(&productData)
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
	tx := p.db.Where(quserid, userid).Delete(&Product{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
