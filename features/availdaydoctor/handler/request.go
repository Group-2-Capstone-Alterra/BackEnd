package handler

import (
	"PetPalApp/features/availdaydoctor"
)

type AvailableDayRequest struct {
	DoctorID  uint   `json:"doctor_id" form:"doctor_id" query:"doctor_id"`
	Monday    string `json:"monday" form:"monday" query:"monday"`
	Tuesday   string `json:"tuesday" form:"tuesday" query:"tuesday"`
	Wednesday string `json:"wednesday" form:"wednesday" query:"wednesday"`
	Thursday  string `json:"thursday" form:"thursday" query:"thursday"`
	Friday    string `json:"friday" form:"friday" query:"friday"`
}

func RequestToCore(doctorID uint, input AvailableDayRequest) availdaydoctor.Core {
	return availdaydoctor.Core{
		DoctorID:  doctorID,
		Monday:    input.Monday,
		Tuesday:   input.Tuesday,
		Wednesday: input.Wednesday,
		Thursday:  input.Thursday,
		Friday:    input.Friday,
	}
}
