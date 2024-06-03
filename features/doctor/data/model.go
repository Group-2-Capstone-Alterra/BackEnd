package data

import (
	"PetPalApp/features/availdaydoctor"
	"PetPalApp/features/doctor"

	"gorm.io/gorm"
)

type Doctor struct {
	gorm.Model
	AdminID        uint
	FullName       string
	Email          string
	Specialization string
}

type AvailableDay struct {
	gorm.Model
	DoctorID  uint
	Monday    bool
	Tuesday   bool
	Wednesday bool
	Thursday  bool
	Friday    bool
}

func GormToCore(doctorGorm Doctor) doctor.Core {
	result := doctor.Core{
		ID:             doctorGorm.ID,
		AdminID:        doctorGorm.AdminID,
		FullName:       doctorGorm.FullName,
		Email:          doctorGorm.Email,
		Specialization: doctorGorm.Specialization,
	}
	return result
}

func AvailGormToCore(availGorm AvailableDay) availdaydoctor.Core {
	result := availdaydoctor.Core{
		ID:        availGorm.ID,
		DoctorID:  availGorm.DoctorID,
		Monday:    availGorm.Monday,
		Tuesday:   availGorm.Tuesday,
		Wednesday: availGorm.Wednesday,
		Thursday:  availGorm.Thursday,
		Friday:    availGorm.Friday,
	}
	return result
}
