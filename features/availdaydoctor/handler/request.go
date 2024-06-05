package handler

import (
	"PetPalApp/features/availdaydoctor"
	"log"
)

type AvailableDayRequest struct {
	DoctorID  uint `json:"doctor_id" form:"doctor_id" query:"doctor_id"`
	Monday    bool `json:"monday" form:"monday" query:"monday"`
	Tuesday   bool `json:"tuesday" form:"tuesday" query:"tuesday"`
	Wednesday bool `json:"wednesday" form:"wednesday" query:"wednesday"`
	Thursday  bool `json:"thursday" form:"thursday" query:"thursday"`
	Friday    bool `json:"friday" form:"friday" query:"friday"`
}

func RequestToCore(doctorID uint, input AvailableDayRequest) availdaydoctor.Core {
	inputCore := availdaydoctor.Core{
		DoctorID:  doctorID,
		Monday:    input.Monday,
		Tuesday:   input.Tuesday,
		Wednesday: input.Wednesday,
		Thursday:  input.Thursday,
		Friday:    input.Friday,
	}
	log.Println("[Handler Req - Availdoc] input", input)
	log.Println("[Handler Req - Availdoc] availGorm", inputCore)
	return inputCore
}
