package servicedoctor

type Core struct {
	ID                  uint `json:"id,omitempty"`
	DoctorID            uint `json:"doctor_id,omitempty"`
	Vaccinations        bool `json:"vaccinations,omitempty"`
	Operations          bool `json:"operations,omitempty"`
	MCU                 bool `json:"mcu,omitempty"`
	OnlineConsultations bool `json:"online_consultations,omitempty"`
}
