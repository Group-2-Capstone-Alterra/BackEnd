package consultation

import "time"

type ConsultationCore struct {
    ID            uint
    UserID        uint
    DoctorID      uint
    Consultation  string
    Response      string
    Status        string
    CreatedAt     time.Time
}

type ConsultationModel interface {
    CreateConsultation(ConsultationCore) error
}

type ConsultationService interface {
    CreateConsultation(ConsultationCore) error
}
