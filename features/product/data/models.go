package data

import (
	"PetPalApp/features/product"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	IdUser         uint
	ProductName    string
	Price          float32
	Stock          uint
	Description    string
	ProductPicture string
}

func CoreToGorm(productCore product.Core) Product {
	gorm := Product{
		IdUser:         productCore.IdUser,
		ProductName:    productCore.ProductName,
		Price:          productCore.Price,
		Stock:          productCore.Stock,
		Description:    productCore.Description,
		ProductPicture: productCore.ProductPicture,
	}
	return gorm
}
