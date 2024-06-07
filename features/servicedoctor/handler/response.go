package handler

import "PetPalApp/features/servicedoctor"

type ServiceResponse struct {
	DoctorID            uint `json:"doctor_id,omitempty"`
	Vaccinations        bool `json:"vaccinations,omitempty"`
	Operations          bool `json:"operations,omitempty" `
	MCU                 bool `json:"mcu,omitempty"`
	OnlineConsultations bool `json:"online_consultations,omitempty"`
}

func GormToCore(gorm servicedoctor.Core) ServiceResponse {
	return ServiceResponse{
		DoctorID:            gorm.DoctorID,
		Vaccinations:        gorm.Vaccinations,
		Operations:          gorm.Operations,
		MCU:                 gorm.MCU,
		OnlineConsultations: gorm.OnlineConsultations,
	}
}
