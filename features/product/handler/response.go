package handler

import (
	"PetPalApp/features/product"
)

type ProductResponse struct {
	ID             uint    `json:"id" form:"id"`
	ProductName    string  `json:"product_name" form:"product_name"`
	Price          float32 `json:"price" form:"price"`
	Stock          uint    `json:"stock" form:"stock"`
	Description    string  `json:"description" form:"description"`
	ProductPicture string  `json:"product_picture" form:"product_picture"`
}

type AllProductResponse struct {
	ProductName    string  `json:"product_name" form:"product_name"`
	Price          float32 `json:"price" form:"price"`
	ProductPicture string  `json:"product_picture" form:"product_picture"`
}

func GormToCore(gorm product.Core) ProductResponse {
	core := ProductResponse{
		ID:             gorm.ID,
		ProductName:    gorm.ProductName,
		Price:          gorm.Price,
		Stock:          gorm.Stock,
		Description:    gorm.Description,
		ProductPicture: gorm.ProductPicture,
	}
	return core
}

func AllGormToCore(gorm product.Core) AllProductResponse {
	core := AllProductResponse{
		ProductName:    gorm.ProductName,
		Price:          gorm.Price,
		ProductPicture: gorm.ProductPicture,
	}
	return core
}
