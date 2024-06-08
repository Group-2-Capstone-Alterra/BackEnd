package data

import (
	"PetPalApp/features/servicedoctor"

	"gorm.io/gorm"
)

type ServiceDoctor struct {
	gorm.Model
	DoctorID            uint
	Vaccinations        string `gorm:"default:false"`
	Operations          string `gorm:"default:false"`
	MCU                 string `gorm:"default:false"`
	OnlineConsultations string `gorm:"default:false"`
}

func ServiceGormToCore(serviceGorm ServiceDoctor) servicedoctor.Core {
	if serviceGorm.Vaccinations == "false" {
		serviceGorm.Vaccinations = ""
	}
	if serviceGorm.Operations == "false" {
		serviceGorm.Operations = ""
	}
	if serviceGorm.MCU == "false" {
		serviceGorm.MCU = ""
	}
	if serviceGorm.OnlineConsultations == "false" {
		serviceGorm.OnlineConsultations = ""
	}

	result := servicedoctor.Core{
		Vaccinations:        serviceGorm.Vaccinations,
		Operations:          serviceGorm.Operations,
		MCU:                 serviceGorm.MCU,
		OnlineConsultations: serviceGorm.OnlineConsultations,
	}
	return result
}
