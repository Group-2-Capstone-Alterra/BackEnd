package data

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	IdUser      uint
	ProductName string
	Price       float32
	Stock       string
	Description string
}
