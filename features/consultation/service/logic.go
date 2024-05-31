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