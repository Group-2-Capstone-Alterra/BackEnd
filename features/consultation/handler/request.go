package handler

import (
	"PetPalApp/features/consultation"
)

type ConsultationRequest struct {
	Service       string `json:"service" form:"service" validate:"required"`
	ScheduledDate string `json:"scheduled_date" form:"scheduled_date" validate:"scheduled_date"`
}

type UpdateConsultationRequest struct {
	TransactionStatus  string `json:"transaction_status" form:"transaction_status"`
	StatusConsultation string `json:"consultation_status" form:"consultation_status"`
}

func UpdateReqToCore(request UpdateConsultationRequest) consultation.ConsultationCore {
	return consultation.ConsultationCore{
		TransactionStatus:  request.TransactionStatus,
		StatusConsultation: request.StatusConsultation,
	}
}

func ReqToCore(userID, doctorID uint, c ConsultationRequest) consultation.ConsultationCore {
	return consultation.ConsultationCore{
		UserID:        uint(userID),
		DoctorID:      uint(doctorID),
		Service:       c.Service,
		ScheduledDate: c.ScheduledDate,
	}
}
