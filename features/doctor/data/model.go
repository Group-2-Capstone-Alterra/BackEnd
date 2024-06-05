package data

import (
	"PetPalApp/features/doctor"

	"gorm.io/gorm"
)

type Doctor struct {
	gorm.Model
	AdminID        uint
	FullName       string
	Email          string
	Specialization string
	ProfilePicture string
}

func GormToCore(doctorGorm Doctor) doctor.Core {
	result := doctor.Core{
		ID:             doctorGorm.ID,
		AdminID:        doctorGorm.AdminID,
		FullName:       doctorGorm.FullName,
		Email:          doctorGorm.Email,
		Specialization: doctorGorm.Specialization,
		ProfilePicture: doctorGorm.ProfilePicture,
	}
	return result
}
