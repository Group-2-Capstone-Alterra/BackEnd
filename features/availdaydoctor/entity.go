package availdaydoctor

type Core struct {
	ID        uint   `json:"id,omitempty"`
	DoctorID  uint   `json:"doctor_id,omitempty"`
	Monday    string `json:"monday,omitempty"`
	Tuesday   string `json:"tuesday,omitempty"`
	Wednesday string `json:"wednesday,omitempty"`
	Thursday  string `json:"thursday,omitempty"`
	Friday    string `json:"friday,omitempty"`
}
