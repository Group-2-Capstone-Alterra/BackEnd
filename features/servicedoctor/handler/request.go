package handler

import (
	"PetPalApp/features/servicedoctor"
)

type ServiceRequest struct {
	DoctorID            uint   `json:"doctor_id" form:"doctor_id" query:"doctor_id"`
	Vaccinations        string `json:"vaccinations" form:"vaccinations" query:"vaccinations"`
	Operations          string `json:"operations" form:"operations" query:"operations"`
	MCU                 string `json:"mcu" form:"mcu" query:"mcu"`
	OnlineConsultations string `json:"online_consultations" form:"online_consultations" query:"online_consultations"`
}

func RequestToCore(doctorID uint, input ServiceRequest) servicedoctor.Core {
	inputCore := servicedoctor.Core{
		DoctorID:            doctorID,
		Vaccinations:        input.Vaccinations,
		Operations:          input.Operations,
		MCU:                 input.MCU,
		OnlineConsultations: input.OnlineConsultations,
	}
	return inputCore
}
