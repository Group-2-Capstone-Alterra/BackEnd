package data

import (
	"PetPalApp/features/consultation"
	"log"

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

func (cm *ConsultationModel) GetCuntationsById(id uint) (*consultation.ConsultationCore, error) {
	var consultationData Consultation
	tx := cm.db.First(&consultationData, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	consultationCore := ToCore(consultationData)
	return &consultationCore, nil
}

func (p *ConsultationModel) VerIsAdmin(userid uint, id uint) (*consultation.ConsultationCore, error) {
	var conculData Consultation
	tx := p.db.Where("doctor_id = ?", userid).Find(&conculData)
	if tx.Error != nil {
		return nil, tx.Error
	}

	conculDataCore := ToCore(conculData)
	if conculDataCore.ID == 0 {
		log.Println("[Query VerIsAdmin] Not Admin - ID Concul", conculDataCore.ID)
	}
	log.Println("[Query VerIsAdmin] Is Admin - ID Concul Notfound", conculDataCore.ID)
	return &conculDataCore, nil
}

func (p *ConsultationModel) VerAvailConcul(currentUserId uint, id uint) (*consultation.ConsultationCore, error) {
	var conculData Consultation
	tx := p.db.Where("doctor_id = ? OR user_id = ?", currentUserId, currentUserId).Find(&conculData, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	conculDataCore := ToCore(conculData)
	if conculDataCore.ID == 0 {
		log.Println("[Query VerIsAdmin] Not Admin - ID Concul", conculDataCore.ID)
	}
	log.Println("[Query VerIsAdmin] Is Admin - ID Concul Notfound", conculDataCore.ID)
	return &conculDataCore, nil
}
