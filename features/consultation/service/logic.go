package service

import (
	"PetPalApp/features/consultation"
)

type ConsultationService struct {
    consultationModel consultation.ConsultationModel
}

func New(cm consultation.ConsultationModel) consultation.ConsultationService {
    return &ConsultationService{
        consultationModel: cm,
    }
}

func (cs *ConsultationService) CreateConsultation(consultation consultation.ConsultationCore) error {
    return cs.consultationModel.CreateConsultation(consultation)
}

func (cs *ConsultationService) GetConsultationsByUserID(userID uint) ([]consultation.ConsultationCore, error) {
    return cs.consultationModel.GetConsultationsByUserID(userID)
}

func (cs *ConsultationService) GetConsultationsByDoctorID(doctorID uint) ([]consultation.ConsultationCore, error) {
    return cs.consultationModel.GetConsultationsByDoctorID(doctorID)
}