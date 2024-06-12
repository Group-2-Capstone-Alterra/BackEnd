package service

import (
	"PetPalApp/features/admin"
	"PetPalApp/features/consultation"
	"PetPalApp/features/doctor"
	"errors"
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
	doctorAvailCheck, _ := cs.dataDoctor.SelectDoctorById(consultation.DoctorID)
	if doctorAvailCheck.ID == 0 {
		return errors.New("Doctor with that ID was not found in any clinic.")
	} else {
		return cs.consultationModel.CreateConsultation(consultation)
	}
}

func (cs *ConsultationService) GetConsultations(currentID uint, role string) ([]consultation.ConsultationCore, error) {
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

func (cs *ConsultationService) GetCuntationsById(id uint) (*consultation.ConsultationCore, error) {
	return cs.consultationModel.GetCuntationsById(id)
}

func (cs *ConsultationService) GetConsultationsByUserID(userID uint) ([]consultation.ConsultationCore, error) {
	return cs.consultationModel.GetConsultationsByUserID(userID)
}

func (cs *ConsultationService) GetConsultationsByDoctorID(doctorID uint) ([]consultation.ConsultationCore, error) {
	return cs.consultationModel.GetConsultationsByDoctorID(doctorID)
}

func (cs *ConsultationService) UpdateConsultation(consulID uint, Core consultation.ConsultationCore) error {
	return cs.consultationModel.UpdateConsultation(consulID, Core)
}
