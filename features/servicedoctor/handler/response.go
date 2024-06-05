package handler

import "PetPalApp/features/servicedoctor"

type ServiceResponse struct {
	DoctorID     uint `json:"doctor_id,omitempty"`
	Vaccinations bool `json:"vaccinations"`
	Operations   bool `json:"operations" `
	MCU          bool `json:"mcu"`
}

func GormToCore(gorm servicedoctor.Core) ServiceResponse {
	return ServiceResponse{
		DoctorID:     gorm.DoctorID,
		Vaccinations: gorm.Vaccinations,
		Operations:   gorm.Operations,
		MCU:          gorm.MCU,
	}
}
