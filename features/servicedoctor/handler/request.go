package handler

import (
	"PetPalApp/features/servicedoctor"
	"log"
)

type ServiceRequest struct {
	DoctorID     uint `json:"doctor_id" form:"doctor_id" query:"doctor_id"`
	Vaccinations bool `json:"vaccinations" form:"vaccinations" query:"vaccinations"`
	Operations   bool `json:"operations" form:"operations" query:"operations"`
	MCU          bool `json:"mcu" form:"mcu" query:"mcu"`
}

func RequestToCore(doctorID uint, input ServiceRequest) servicedoctor.Core {
	inputCore := servicedoctor.Core{
		DoctorID:     doctorID,
		Vaccinations: input.Vaccinations,
		Operations:   input.Operations,
		MCU:          input.MCU,
	}
	log.Println("[Handler Req - ServiceDoc] input", input)
	log.Println("[Handler Req - ServiceDoc] serviceCore", inputCore)
	return inputCore
}
