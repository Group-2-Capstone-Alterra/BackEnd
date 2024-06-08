package handler

import "PetPalApp/features/availdaydoctor"

type AvailableDayResponse struct {
	DoctorID  uint   `json:"doctor_id,omitempty"`
	Monday    string `json:"monday,omitempty"`
	Tuesday   string `json:"tuesday,omitempty"`
	Wednesday string `json:"wednesday,omitempty"`
	Thursday  string `json:"thursday,omitempty"`
	Friday    string `json:"friday,omitempty"`
}

func GormToCore(gorm availdaydoctor.Core) AvailableDayResponse {
	return AvailableDayResponse{
		Monday:    gorm.Monday,
		Tuesday:   gorm.Tuesday,
		Wednesday: gorm.Wednesday,
		Thursday:  gorm.Thursday,
		Friday:    gorm.Friday,
	}
}
