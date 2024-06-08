package handler

import "PetPalApp/features/consultation"

type ConsultationRequest struct {
	Service string `json:"service" form:"service" validate:"required"`
}

type UpdateConsultationRequest struct {
	TransactionStatus  string `json:"transaction_status" form:"transaction_status"`
	StatusConsultation string `json:"consultation_status" form:"consultation_status"`
}

func ReqToCore(request UpdateConsultationRequest) consultation.ConsultationCore {
	return consultation.ConsultationCore{
		TransactionStatus:  request.TransactionStatus,
		StatusConsultation: request.StatusConsultation,
	}
}
