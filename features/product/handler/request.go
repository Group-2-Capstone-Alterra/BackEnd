package handler

import (
	"PetPalApp/features/product"

	"gorm.io/gorm"
)

type ProductRequest struct {
	gorm.Model
	AdminID        uint
	ProductName    string  `json:"product_name" form:"product_name"`
	Price          float32 `json:"price" form:"price"`
	Stock          uint    `json:"stock" form:"stock"`
	Description    string  `json:"description" form:"description"`
	ProductPicture string  `json:"product_picture" form:"product_picture"`
}

func RequestToCore(input ProductRequest) product.Core {
	inputCore := product.Core{
		AdminID:        input.AdminID,
		ProductName:    input.ProductName,
		Price:          input.Price,
		Stock:          input.Stock,
		Description:    input.Description,
		ProductPicture: input.ProductPicture,
	}
	return inputCore
}
