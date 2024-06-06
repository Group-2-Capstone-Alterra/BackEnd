package handler

import (
	"PetPalApp/features/product"
)

type ProductResponse struct {
	ProductName    string  `json:"product_name" form:"product_name"`
	Price          float32 `json:"price" form:"price"`
	Stock          uint    `json:"stock" form:"stock"`
	Description    string  `json:"description" form:"description"`
	ProductPicture string  `json:"product_picture" form:"product_picture"`
}

type AllProductResponse struct {
	ID             uint    `json:"id" form:"id"`
	ProductName    string  `json:"product_name" form:"product_name"`
	Price          float32 `json:"price" form:"price"`
	ProductPicture string  `json:"product_picture" form:"product_picture"`
}

type OrderResponse struct {
	ID          uint   `json:"id"`
	ProductName string `json:"product_name"`
}

func OrderCoreToResponse(core product.Core) OrderResponse {
	return OrderResponse{
		ID:          core.ID,
		ProductName: core.ProductName,
	}
}

func GormToCore(gorm product.Core) ProductResponse {
	core := ProductResponse{
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
		ID:             gorm.ID,
		ProductName:    gorm.ProductName,
		Price:          gorm.Price,
		ProductPicture: gorm.ProductPicture,
	}
	return core
}
