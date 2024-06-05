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
		// Status:       "Pending", // default status

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
	tx := p.db.Where("doctor_id = ?", userid).Find(&conculData)
	if tx.Error != nil {
		return nil, tx.Error
	}

	conculDataCore := ToCore(conculData)
	if conculDataCore.ID == 0 {
		log.Println("[Query VerIsAdmin] Not Admin")
	} else {
		log.Println("[Query VerIsAdmin] Is Admin")
	}
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
		log.Println("[Query VerAvailConcul] Roomchat Not Found")
	} else {
		log.Println("[Query VerAvailConcul] Roomchat Has Found")
	}
	return &conculDataCore, nil
}

func (p *ConsultationModel) VerUser(userID uint, doctorID uint, roomchatID uint) (*consultation.ConsultationCore, error) {
	var conculData Consultation
	tx := p.db.Where("user_id = ?", userID).Where("doctor_id = ?", doctorID).Find(&conculData, roomchatID)

	if tx.Error != nil {
		return nil, tx.Error
	}

	conculDataCore := ToCore(conculData)
	if conculDataCore.ID == 0 {
		log.Printf("\n[Query VerUser] User not found at RoomchatID\n")
	} else {
		log.Printf("\n[Query VerUser] User has found at RoomchatID\n")

	}
	return &conculDataCore, nil
}

func (p *ConsultationModel) VerAdmin(doctorID uint, userID uint, roomchatID uint) (*consultation.ConsultationCore, error) {
	var conculData Consultation
	tx := p.db.Where("doctor_id = ?", doctorID).Where("user_id = ?", userID).Find(&conculData, roomchatID)

	if tx.Error != nil {
		return nil, tx.Error
	}

	conculDataCore := ToCore(conculData)
	if conculDataCore.ID == 0 {
		log.Printf("\n[Query VerAdmin] Doctor not found at RoomchatID\n")
	} else {
		log.Printf("\n[Query VerAdmin] Doctor has found at RoomchatID\n")

	}
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
	tx := cm.db.Where("user_id = ?", userID).Find(&consultations)
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
	tx := cm.db.Where("doctor_id = ?", doctorID).Find(&consultations)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var result []consultation.ConsultationCore
	for _, consultation := range consultations {
		result = append(result, ToCore(consultation))
	}

	return result, nil
}

func (cm *ConsultationModel) UpdateConsultationResponse(consultationID uint, response string) error {
	tx := cm.db.Model(&Consultation{}).Where("id = ?", consultationID).Update("response", response)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
