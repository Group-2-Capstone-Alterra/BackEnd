package data

import (
	"PetPalApp/features/availdaydoctor"
	"log"

	"gorm.io/gorm"
)

type AvailableDay struct {
	gorm.Model
	DoctorID  uint
	Monday    bool
	Tuesday   bool
	Wednesday bool
	Thursday  bool
	Friday    bool
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
	log.Println("[Data - Availdoc] availGorm", availGorm)
	return result
}
