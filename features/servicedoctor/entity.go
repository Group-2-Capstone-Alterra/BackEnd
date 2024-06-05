package servicedoctor

type Core struct {
	ID           uint `json:"id"`
	DoctorID     uint `json:"doctor_id"`
	Vaccinations bool `json:"vaccinations"`
	Operations   bool `json:"operations"`
	MCU          bool `json:"mcu"`
}
