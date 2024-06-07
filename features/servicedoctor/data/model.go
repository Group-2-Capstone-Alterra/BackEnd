package data

import (
	"PetPalApp/features/servicedoctor"
	"log"

	"gorm.io/gorm"
)

type ServiceDoctor struct {
	gorm.Model
	DoctorID            uint
	Vaccinations        bool `gorm:"default:null"`
	Operations          bool `gorm:"default:null"`
	MCU                 bool `gorm:"default:null"`
	OnlineConsultations bool `gorm:"default:null"`
}

func ServiceGormToCore(serviceGorm ServiceDoctor) servicedoctor.Core {
	result := servicedoctor.Core{
		Vaccinations:        serviceGorm.Vaccinations,
		Operations:          serviceGorm.Operations,
		MCU:                 serviceGorm.MCU,
		OnlineConsultations: serviceGorm.OnlineConsultations,
	}
	log.Println("[Data - ServiceGormToCore] serviceGorm", serviceGorm)
	return result
}
