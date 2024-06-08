package data

import (
	"PetPalApp/features/availdaydoctor"

	"gorm.io/gorm"
)

type AvailableDay struct {
	gorm.Model
	DoctorID  uint
	Monday    string `gorm:"default:false"`
	Tuesday   string `gorm:"default:false"`
	Wednesday string `gorm:"default:false"`
	Thursday  string `gorm:"default:false"`
	Friday    string `gorm:"default:false"`
}

func AvailGormToCore(availGorm AvailableDay) availdaydoctor.Core {
	if availGorm.Monday == "false" {
		availGorm.Monday = ""
	}
	if availGorm.Tuesday == "false" {
		availGorm.Tuesday = ""
	}
	if availGorm.Wednesday == "false" {
		availGorm.Wednesday = ""
	}
	if availGorm.Thursday == "false" {
		availGorm.Thursday = ""
	}
	if availGorm.Friday == "false" {
		availGorm.Friday = ""
	}
	return availdaydoctor.Core{
		Monday:    availGorm.Monday,
		Tuesday:   availGorm.Tuesday,
		Wednesday: availGorm.Wednesday,
		Thursday:  availGorm.Thursday,
		Friday:    availGorm.Friday,
	}
}
