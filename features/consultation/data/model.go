package data

import (
	chat "PetPalApp/features/chat/data"
	"PetPalApp/features/consultation"

	"gorm.io/gorm"
)

type Consultation struct {
	gorm.Model
	UserID             uint
	DoctorID           uint
	Consultation       string
	TransactionStatus  string 		`gorm:"default:'Pending'"`
	StatusConsultation string 		`gorm:"default:'New Consultation'"`
	Chats			   []chat.Chat  `gorm:"foreign_key:ConsultationID"`
}

const (
	Pending = "Pending"
	Paid    = "Paid"
	Failed  = "Failed"
)

const (
	NewConsultation = "New Consultation"
	InProgress      = "In Progress"
	Finished        = "Finished"
)

func ToCore(c Consultation) consultation.ConsultationCore {
	return consultation.ConsultationCore{
		ID:                 c.ID,
		UserID:             c.UserID,
		DoctorID:           c.DoctorID,
		Consultation:       c.Consultation,
		TransactionStatus:  c.TransactionStatus,
		StatusConsultation: c.StatusConsultation,
		CreatedAt:          c.CreatedAt,
	}
}
