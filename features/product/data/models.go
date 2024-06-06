package data

import (
	"PetPalApp/features/product"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	AdminID        uint
	ProductName    string
	Price          float32
	Stock          uint
	Description    string
	ProductPicture string
}

func CoreToGorm(core product.Core) Product {
	gorm := Product{
		AdminID:        core.AdminID,
		ProductName:    core.ProductName,
		Price:          core.Price,
		Stock:          core.Stock,
		Description:    core.Description,
		ProductPicture: core.ProductPicture,
	}
	return gorm
}

func GormToCore(gorm Product) product.Core {
	core := product.Core{
		ID:             gorm.ID,
		AdminID:        gorm.AdminID,
		ProductName:    gorm.ProductName,
		Price:          gorm.Price,
		Stock:          gorm.Stock,
		Description:    gorm.Description,
		ProductPicture: gorm.ProductPicture,
	}
	return core
}
