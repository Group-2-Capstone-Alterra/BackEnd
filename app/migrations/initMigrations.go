package migrations

import (
	_dataAdmin "PetPalApp/features/admin/data"
	_dataProduct "PetPalApp/features/product/data"
	_dataUser "PetPalApp/features/user/data"

	"gorm.io/gorm"
)

func InitMigrations(db *gorm.DB) {
	db.AutoMigrate(&_dataUser.User{})
	db.AutoMigrate(&_dataAdmin.Admin{})
	db.AutoMigrate(&_dataProduct.Product{})
}
