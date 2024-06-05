package availdaydoctor

type Core struct {
	ID        uint `json:"id,omitempty"`
	DoctorID  uint `json:"doctor_id,omitempty"`
	Monday    bool `json:"monday"`
	Tuesday   bool `json:"tuesday"`
	Wednesday bool `json:"wednesday"`
	Thursday  bool `json:"thursday"`
	Friday    bool `json:"friday"`
}
