package handler

import (
	"PetPalApp/features/consultation"
	"PetPalApp/features/doctor"
	_doctorHandler "PetPalApp/features/doctor/handler"
	"PetPalApp/features/user"
	_userHandler "PetPalApp/features/user/handler"

	"time"
)

// type ConsultationResponse struct {
// 	ID                   uint      `json:"id"`
// 	UserID               uint      `json:"user_id"`
// 	UserFullName         string    `json:"user_fullname"`
// 	UserProfilePicture   string    `json:"user_profile_picture"`
// 	DoctorID             uint      `json:"doctor_id"`
// 	DoctorFullName       string    `json:"doctor_fullname"`
// 	DoctorProfilePicture string    `json:"doctor_profile_picture"`
// 	Consultation         string    `json:"consultation"`
// 	TransactionStatus    string    `json:"transaction_status"`
// 	StatusConsultation   string    `json:"status_consultation"`
// 	CreatedAt            time.Time `json:"created_at"`
// }

// func AllConsultationResponseUser(gorm consultation.ConsultationCore, user user.Core, doctor doctor.Core) ConsultationResponse {
// 	result := ConsultationResponse{
// 		ID:                   gorm.ID,
// 		UserID:               user.ID,
// 		UserFullName:         user.FullName,
// 		UserProfilePicture:   user.ProfilePicture,
// 		DoctorID:             doctor.ID,
// 		DoctorFullName:       doctor.FullName,
// 		DoctorProfilePicture: doctor.ProfilePicture,
// 		Consultation:         gorm.Consultation,
// 		TransactionStatus:    gorm.TransactionStatus,
// 		StatusConsultation:   gorm.StatusConsultation,
// 	}
// 	return result
// }

type ConsultationResponse struct {
	ID                 uint
	UserDetails        _userHandler.ConsulUserReponse
	DoctorDetails      _doctorHandler.ConsulDoctorReponse
	Consultation       string
	TransactionStatus  string
	StatusConsultation string
	CreatedAt          time.Time
}

func GormToCore(gormConsul consultation.ConsultationCore, gormUser user.Core, gormDoctor doctor.Core) ConsultationResponse {
	result := ConsultationResponse{
		ID:                 gormConsul.ID,
		UserDetails:        _userHandler.ConsulCoreToGorm(gormUser),
		DoctorDetails:      _doctorHandler.ConsulGormToCore(gormDoctor),
		Consultation:       gormConsul.Consultation,
		TransactionStatus:  gormConsul.TransactionStatus,
		StatusConsultation: gormConsul.StatusConsultation,
		CreatedAt:          gormConsul.CreatedAt,
	}
	return result
}
