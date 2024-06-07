package availdaydoctor

type Core struct {
	ID        uint `json:"id,omitempty"`
	DoctorID  uint `json:"doctor_id,omitempty"`
	Monday    bool `json:"monday,omitempty"`
	Tuesday   bool `json:"tuesday,omitempty"`
	Wednesday bool `json:"wednesday,omitempty"`
	Thursday  bool `json:"thursday,omitempty"`
	Friday    bool `json:"friday,omitempty"`
}
