package migrations

import (
	_dataAdmin "PetPalApp/features/admin/data"
	_dataAvailday "PetPalApp/features/availdaydoctor/data"
	_dataChat "PetPalApp/features/chat/data"
	_dataConsultation "PetPalApp/features/consultation/data"
	_dataDoctor "PetPalApp/features/doctor/data"
	_dataOrder "PetPalApp/features/order/data"
	_dataPayment "PetPalApp/features/payment/data"
	_dataProduct "PetPalApp/features/product/data"
	_dataServiceDoctor "PetPalApp/features/servicedoctor/data"
	_dataUser "PetPalApp/features/user/data"

	"gorm.io/gorm"
)

func InitMigrations(db *gorm.DB) {
	db.AutoMigrate(&_dataUser.User{})
	db.AutoMigrate(&_dataAdmin.Admin{})
	db.AutoMigrate(&_dataProduct.Product{})
	db.AutoMigrate(&_dataDoctor.Doctor{})
	db.AutoMigrate(&_dataAvailday.AvailableDay{})
	db.AutoMigrate(&_dataChat.Chat{})
	db.AutoMigrate(&_dataOrder.Order{})
	db.AutoMigrate(&_dataConsultation.Consultation{})
	db.AutoMigrate(&_dataPayment.Payment{})
	db.AutoMigrate(&_dataServiceDoctor.ServiceDoctor{})
}
