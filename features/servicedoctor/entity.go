package servicedoctor

type Core struct {
	ID           uint `json:"id,omitempty"`
	DoctorID     uint `json:"doctor_id,omitempty"`
	Vaccinations bool `json:"vaccinations"`
	Operations   bool `json:"operations"`
	MCU          bool `json:"mcu"`
}
