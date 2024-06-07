package handler

import "PetPalApp/features/availdaydoctor"

type AvailableDayResponse struct {
	DoctorID  uint `json:"doctor_id,omitempty"`
	Monday    bool `json:"monday,omitempty"`
	Tuesday   bool `json:"tuesday,omitempty"`
	Wednesday bool `json:"wednesday,omitempty"`
	Thursday  bool `json:"thursday,omitempty"`
	Friday    bool `json:"friday,omitempty"`
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
