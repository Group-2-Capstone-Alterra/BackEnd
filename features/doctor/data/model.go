package data

import (
	"PetPalApp/features/doctor"

	"gorm.io/gorm"
)

type Doctor struct {
	gorm.Model
	AdminID        uint
	FullName       string
	About          string
	Price          float32
	ProfilePicture string
}

func GormToCore(doctorGorm Doctor) doctor.Core {
	result := doctor.Core{
		ID:             doctorGorm.ID,
		AdminID:        doctorGorm.AdminID,
		FullName:       doctorGorm.FullName,
		About:          doctorGorm.About,
		Price:          doctorGorm.Price,
		ProfilePicture: doctorGorm.ProfilePicture,
	}
	return result
}
