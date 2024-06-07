package service

import (
	"PetPalApp/features/admin"
	"PetPalApp/features/consultation"
	"PetPalApp/features/doctor"
	"log"
)

type ConsultationService struct {
	consultationModel consultation.ConsultationModel
	dataDoctor        doctor.DoctorModel
	dataAdmin         admin.AdminModel
}

func New(cm consultation.ConsultationModel, dataDoctor doctor.DoctorModel, dataAdmin admin.AdminModel) consultation.ConsultationService {
	return &ConsultationService{
		consultationModel: cm,
		dataDoctor:        dataDoctor,
		dataAdmin:         dataAdmin,
	}
}

func (cs *ConsultationService) CreateConsultation(consultation consultation.ConsultationCore) error {
	return cs.consultationModel.CreateConsultation(consultation)
}

func (cs *ConsultationService) GetConsultations(currentID uint, role string) ([]consultation.ConsultationCore, error) {
	log.Println("[Service - GetConsultations] Role : ", role)
	//check doctor or not
	if role == "user" {
		log.Println("[Service - GetConsultations] Role User")
		return cs.consultationModel.GetConsultationsByUserID(currentID)
	} else {
		log.Println("[Service - GetConsultations] Role Admin")
		doctorID, _ := cs.dataDoctor.SelectByAdminId(currentID)
		return cs.consultationModel.GetConsultationsByDoctorID(doctorID.ID)
	}
}

func (cs *ConsultationService) GetConsultationsByUserID(userID uint) ([]consultation.ConsultationCore, error) {
	return cs.consultationModel.GetConsultationsByUserID(userID)
}

func (cs *ConsultationService) GetConsultationsByDoctorID(doctorID uint) ([]consultation.ConsultationCore, error) {
	return cs.consultationModel.GetConsultationsByDoctorID(doctorID)
}

func (cs *ConsultationService) UpdateConsultationResponse(consultationID uint, response string) error {
	return cs.consultationModel.UpdateConsultationResponse(consultationID, response)
}
