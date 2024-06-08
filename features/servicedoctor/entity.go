package servicedoctor

type Core struct {
	ID                  uint   `json:"id,omitempty"`
	DoctorID            uint   `json:"doctor_id,omitempty"`
	Vaccinations        string `json:"vaccinations,omitempty"`
	Operations          string `json:"operations,omitempty"`
	MCU                 string `json:"mcu,omitempty"`
	OnlineConsultations string `json:"online_consultations,omitempty"`
}
