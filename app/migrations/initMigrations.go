package migrations

import (
	_dataAdmin "PetPalApp/features/admin/data"
	_dataChat "PetPalApp/features/chat/data"
	_dataDoctor "PetPalApp/features/doctor/data"
	_dataProduct "PetPalApp/features/product/data"
	_dataUser "PetPalApp/features/user/data"

	"gorm.io/gorm"
)

func InitMigrations(db *gorm.DB) {
	db.AutoMigrate(&_dataUser.User{})
	db.AutoMigrate(&_dataAdmin.Admin{})
	db.AutoMigrate(&_dataProduct.Product{})
	db.AutoMigrate(&_dataDoctor.Doctor{})
	db.AutoMigrate(&_dataChat.Chat{})
}
