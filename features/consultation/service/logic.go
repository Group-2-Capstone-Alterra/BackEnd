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

func (cs *ConsultationService) GetConsultations(currentID uint) ([]consultation.ConsultationCore, error) {
	//check doctor or not
	isDoctor, _ := cs.dataDoctor.SelectByAdminId(currentID)
	// isAdmin, _ := cs.dataAdmin.AdminById(currentID)
	if isDoctor.ID != 0 { //is doctor
		log.Println("[Service - GetConsultations] Doctor")
		return cs.consultationModel.GetConsultationsByDoctorID(isDoctor.ID)
	} else { //not doctor
		log.Println("[Service - GetConsultations] User")
		return cs.consultationModel.GetConsultationsByUserID(currentID)
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
