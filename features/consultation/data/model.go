package data

import (
	"PetPalApp/features/consultation"

	"gorm.io/gorm"
)

type Consultation struct {
	gorm.Model
	UserID             uint
	DoctorID           uint
	Service            string
	TransactionStatus  string `gorm:"default:'Pending'"`
	StatusConsultation string `gorm:"default:'New Consultation'"`
}

func ToCore(c Consultation) consultation.ConsultationCore {
	return consultation.ConsultationCore{
		ID:                 c.ID,
		UserID:             c.UserID,
		DoctorID:           c.DoctorID,
		Service:            c.Service,
		TransactionStatus:  c.TransactionStatus,
		StatusConsultation: c.StatusConsultation,
		CreatedAt:          c.CreatedAt,
	}
}
