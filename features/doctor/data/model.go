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
	ProfilePicture string `gorm:"default:'https://air-bnb.s3.ap-southeast-2.amazonaws.com/profilepicture/default.jpg'"`
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

func CoreToGorm(input doctor.Core) Doctor {
	return Doctor{
		AdminID:        input.AdminID,
		FullName:       input.FullName,
		ProfilePicture: input.ProfilePicture,
		About:          input.About,
		Price:          input.Price,
	}
}
