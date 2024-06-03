package availdaydoctor

type Core struct {
	ID        uint `json:"id"`
	DoctorID  uint `json:"doctor_id"`
	Monday    bool `json:"monday"`
	Tuesday   bool `json:"tuesday"`
	Wednesday bool `json:"wednesday"`
	Thursday  bool `json:"thursday"`
	Friday    bool `json:"friday"`
}
