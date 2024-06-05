package data

import (
	"PetPalApp/features/availdaydoctor"
	"log"

	"gorm.io/gorm"
)

type AvailableDay struct {
	gorm.Model
	DoctorID  uint
	Monday    bool `gorm:"default:null"`
	Tuesday   bool `gorm:"default:null"`
	Wednesday bool `gorm:"default:null"`
	Thursday  bool `gorm:"default:null"`
	Friday    bool `gorm:"default:null"`
}

func AvailGormToCore(availGorm AvailableDay) availdaydoctor.Core {
	result := availdaydoctor.Core{

		Monday:    availGorm.Monday,
		Tuesday:   availGorm.Tuesday,
		Wednesday: availGorm.Wednesday,
		Thursday:  availGorm.Thursday,
		Friday:    availGorm.Friday,
	}
	log.Println("[Data - Availdoc] availGorm", availGorm)
	return result
}
