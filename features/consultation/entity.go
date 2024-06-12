package consultation

type ConsultationCore struct {
	ID                 uint
	UserID             uint
	DoctorID           uint
	Service            string
	TransactionStatus  string
	StatusConsultation string
	ScheduledDate      string
}

type ConsultationModel interface {
	CreateConsultation(ConsultationCore) error
	GetCuntationsById(id uint) (*ConsultationCore, error)
	VerIsDoctor(userid uint, id uint) (*ConsultationCore, error)
	VerAvailConcul(currentUserId uint, id uint) (*ConsultationCore, error)
	VerUser(userID uint, doctorID uint, roomchatID uint) (*ConsultationCore, error)
	VerAdmin(userID uint, doctorID uint, roomchatID uint) (*ConsultationCore, error)
	GetConsultations(currentID uint) ([]ConsultationCore, error)
	GetConsultationsByUserID(userID uint) ([]ConsultationCore, error)
	GetConsultationsByDoctorID(doctorID uint) ([]ConsultationCore, error)
	UpdateConsultation(consultationID uint, core ConsultationCore) error
}

type ConsultationService interface {
	GetCuntationsById(id uint) (*ConsultationCore, error)
	CreateConsultation(ConsultationCore) error
	GetConsultations(currentID uint, role string) ([]ConsultationCore, error)
	GetConsultationsByUserID(userID uint) ([]ConsultationCore, error)
	GetConsultationsByDoctorID(doctorID uint) ([]ConsultationCore, error)
	UpdateConsultation(consulID uint, Core ConsultationCore) error
}
