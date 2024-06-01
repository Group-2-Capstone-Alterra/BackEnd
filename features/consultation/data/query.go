package data

import (
	"PetPalApp/features/consultation"

	"gorm.io/gorm"
)

type ConsultationModel struct {
    db *gorm.DB
}

func New(db *gorm.DB) consultation.ConsultationModel {
    return &ConsultationModel{
        db: db,
    }
}

func (cm *ConsultationModel) CreateConsultation(consultationCore consultation.ConsultationCore) error {
    consultationGorm := Consultation{
        UserID:       consultationCore.UserID,
        DoctorID:     consultationCore.DoctorID,
        Consultation: consultationCore.Consultation,
        Status:       "Pending", // default status
    }
    tx := cm.db.Create(&consultationGorm)
    if tx.Error != nil {
        return tx.Error
    }
    return nil
}