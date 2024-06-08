package handler

import "PetPalApp/features/servicedoctor"

type ServiceResponse struct {
	DoctorID            uint   `json:"doctor_id,omitempty"`
	Vaccinations        string `json:"vaccinations,omitempty"`
	Operations          string `json:"operations,omitempty" `
	MCU                 string `json:"mcu,omitempty"`
	OnlineConsultations string `json:"online_consultations,omitempty"`
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
