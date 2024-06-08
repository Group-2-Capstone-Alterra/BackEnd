package handler

import (
	"PetPalApp/features/consultation"
	"PetPalApp/features/doctor"
	_doctorHandler "PetPalApp/features/doctor/handler"
	"PetPalApp/features/user"
	_userHandler "PetPalApp/features/user/handler"

	"time"
)

type ConsultationResponse struct {
	ID                 uint                               `json:"id,omitempty"`
	UserDetails        _userHandler.ConsulUserReponse     `json:"user_details,omitempty"`
	DoctorDetails      _doctorHandler.ConsulDoctorReponse `json:"doctor_details,omitempty"`
	Service            string                             `json:"service,omitempty"`
	TransactionStatus  string                             `json:"transaction_status,omitempty"`
	StatusConsultation string                             `json:"consultation_status,omitempty"`
	CreatedAt          time.Time                          `json:"created_at,omitempty"`
}

func GormToCore(gormConsul consultation.ConsultationCore, gormUser user.Core, gormDoctor doctor.Core) ConsultationResponse {
	result := ConsultationResponse{
		ID:                 gormConsul.ID,
		UserDetails:        _userHandler.ConsulCoreToGorm(gormUser),
		DoctorDetails:      _doctorHandler.ConsulGormToCore(gormDoctor),
		Service:            gormConsul.Service,
		TransactionStatus:  gormConsul.TransactionStatus,
		StatusConsultation: gormConsul.StatusConsultation,
		CreatedAt:          gormConsul.CreatedAt,
	}
	return result
}
