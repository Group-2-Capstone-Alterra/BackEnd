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
	ID                 uint                               `json:"id"`
	UserDetails        _userHandler.ConsulUserReponse     `json:"user_details"`
	DoctorDetails      _doctorHandler.ConsulDoctorReponse `json:"doctor_details"`
	Service            string                             `json:"service"`
	TransactionStatus  string                             `json:"transaction_status"`
	StatusConsultation string                             `json:"consultation_status"`
	CreatedAt          time.Time                          `json:"created_at"`
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
