package data

import (
	orderProduct "PetPalApp/features/order_product/data"
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
	OrderProdcts   []orderProduct.OrderProduct `gorm:"foreign_key:ProductID"`
}

func CoreToGorm(core product.Core) Product {
	gorm := Product{
		IdUser:         core.IdUser,
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
		IdUser:         gorm.IdUser,
		ProductName:    gorm.ProductName,
		Price:          gorm.Price,
		Stock:          gorm.Stock,
		Description:    gorm.Description,
		ProductPicture: gorm.ProductPicture,
	}
	return core
}
