package consultation

import "time"

type ConsultationCore struct {
	ID                 uint
	UserID             uint
	DoctorID           uint
	Consultation       string
	TransactionStatus  string
	StatusConsultation string
	CreatedAt          time.Time
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
	UpdateConsultationResponse(consultationID uint, response string) error
}

type ConsultationService interface {
	CreateConsultation(ConsultationCore) error
	GetConsultations(currentID uint) ([]ConsultationCore, error)
	GetConsultationsByUserID(userID uint) ([]ConsultationCore, error)
	GetConsultationsByDoctorID(doctorID uint) ([]ConsultationCore, error)
	UpdateConsultationResponse(consultationID uint, response string) error
}
