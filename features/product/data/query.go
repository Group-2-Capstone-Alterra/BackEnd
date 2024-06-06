package data

import (
	"PetPalApp/features/product"
	"PetPalApp/utils/helper"
	"log"

	"gorm.io/gorm"
)

type productrQuery struct {
	db     *gorm.DB
	helper helper.HelperInterface
}

func New(db *gorm.DB, helper helper.HelperInterface) product.DataInterface {
	return &productrQuery{
		db:     db,
		helper: helper,
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

	log.Println("[Query] sortStr", sortStr)
	if sortStr == "lowest distance" || sortStr == "lowest" {
		log.Println("[query] lowest+distance")

		tx := p.db.Order("price asc").Limit(10).Offset(int(offset)).Find(&allProduct)
		if tx.Error != nil {
			return nil, tx.Error
		}

	} else if sortStr == "highest distance" || sortStr == "highest" {
		log.Println("[query] highest+distance")

		tx := p.db.Order("price desc").Limit(10).Offset(int(offset)).Find(&allProduct)
		if tx.Error != nil {
			return nil, tx.Error
		}
	} else {

		log.Println("[query] default")

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

func (p *productrQuery) SelectAllAdmin(adminid uint, offset uint) ([]product.Core, error) {
	var allProduct []Product
	log.Println("[Query - admin] SelectAllAdmin, userid: ", adminid)

	tx := p.db.Where("admin_id = ?", adminid).Find(&allProduct)
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

func (p *productrQuery) SelectByIdAdmin(id uint, adminid uint) (*product.Core, error) {
	var productData Product
	tx := p.db.Where("admin_id = ?", adminid).First(&productData, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	projectcore := GormToCore(productData)
	return &projectcore, nil
}

func (p *productrQuery) VerIsAdmin(adminid uint) (*product.Core, error) {
	var productData Product
	tx := p.db.Where("admin_id = ?", adminid).Find(&productData)
	if tx.Error != nil {
		return nil, tx.Error
	}

	projectcore := GormToCore(productData)
	if projectcore.ID == 0 {
		log.Println("[Query VerIsAdmin] not admin", projectcore.ID)
	} else {
		log.Println("[Query VerIsAdmin] is admin", projectcore.ID)
	}

	return &projectcore, nil
}

func (p *productrQuery) PutById(id uint, adminid uint, input product.Core) error {

	userGorm := CoreToGorm(input)

	tx := p.db.Model(&Product{}).Where("id = ? AND admin_id = ?", id, adminid).Updates(&userGorm)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (p *productrQuery) Delete(id uint, adminid uint) error {
	tx := p.db.Where("admin_id = ?", adminid).Delete(&Product{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
