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

const (
	qdoctorID = "doctor_id = ?"
	quserID   = "user_id = ?"
)

func (cm *ConsultationModel) CreateConsultation(consultationCore consultation.ConsultationCore) error {
	consultationGorm := Consultation{
		UserID:   consultationCore.UserID,
		DoctorID: consultationCore.DoctorID,
		Service:  consultationCore.Service,
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

func (p *ConsultationModel) VerIsDoctor(userid uint, id uint) (*consultation.ConsultationCore, error) {
	var conculData Consultation
	tx := p.db.Where(qdoctorID, userid).Find(&conculData)
	if tx.Error != nil {
		return nil, tx.Error
	}

	conculDataCore := ToCore(conculData)
	return &conculDataCore, nil
}

func (p *ConsultationModel) VerAvailConcul(currentUserId uint, id uint) (*consultation.ConsultationCore, error) {
	var conculData Consultation
	tx := p.db.Where("doctor_id = ? OR user_id = ?", currentUserId, currentUserId).Find(&conculData, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	conculDataCore := ToCore(conculData)
	return &conculDataCore, nil
}

func (p *ConsultationModel) VerUser(userID uint, doctorID uint, roomchatID uint) (*consultation.ConsultationCore, error) {
	var conculData Consultation
	tx := p.db.Where(quserID, userID).Where(qdoctorID, doctorID).Find(&conculData, roomchatID)

	if tx.Error != nil {
		return nil, tx.Error
	}

	conculDataCore := ToCore(conculData)
	return &conculDataCore, nil
}

func (p *ConsultationModel) VerAdmin(doctorID uint, userID uint, roomchatID uint) (*consultation.ConsultationCore, error) {
	var conculData Consultation
	tx := p.db.Where(qdoctorID, doctorID).Where(quserID, userID).Find(&conculData, roomchatID)

	if tx.Error != nil {
		return nil, tx.Error
	}

	conculDataCore := ToCore(conculData)
	return &conculDataCore, nil
}

func (cm *ConsultationModel) GetConsultations(currentID uint) ([]consultation.ConsultationCore, error) {
	var consultations []Consultation
	tx := cm.db.Where("user_id = ? OR doctor_id = ?", currentID, currentID).Find(&consultations)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var result []consultation.ConsultationCore
	for _, consultation := range consultations {
		result = append(result, ToCore(consultation))
	}

	return result, nil
}

func (cm *ConsultationModel) GetConsultationsByUserID(userID uint) ([]consultation.ConsultationCore, error) {
	var consultations []Consultation
	tx := cm.db.Where(quserID, userID).Find(&consultations)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var result []consultation.ConsultationCore
	for _, consultation := range consultations {
		result = append(result, ToCore(consultation))
	}

	return result, nil
}

func (cm *ConsultationModel) GetConsultationsByDoctorID(doctorID uint) ([]consultation.ConsultationCore, error) {
	var consultations []Consultation
	tx := cm.db.Where(qdoctorID, doctorID).Find(&consultations)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var result []consultation.ConsultationCore
	for _, consultation := range consultations {
		result = append(result, ToCore(consultation))
	}

	return result, nil
}

func (cm *ConsultationModel) UpdateConsultation(consultationID uint, Core consultation.ConsultationCore) error {
	tx := cm.db.Model(&Consultation{}).Where("id = ?", consultationID).Update("status_consultation", Core.StatusConsultation)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
