package handler

import (
	"PetPalApp/features/consultation"
	"PetPalApp/features/doctor"
	"PetPalApp/features/user"
	"time"
)

type ConsultationResponse struct {
	ID                   uint      `json:"id"`
	UserID               uint      `json:"user_id"`
	UserFullName         string    `json:"user_fullname"`
	UserProfilePicture   string    `json:"user_profile_picture"`
	DoctorID             uint      `json:"doctor_id"`
	DoctorFullName       string    `json:"doctor_fullname"`
	DoctorProfilePicture string    `json:"doctor_profile_picture"`
	Consultation         string    `json:"consultation"`
	TransactionStatus    string    `json:"transaction_status"`
	StatusConsultation   string    `json:"status_consultation"`
	CreatedAt            time.Time `json:"created_at"`
}

func AllConsultationResponseUser(gorm consultation.ConsultationCore, user user.Core, doctor doctor.Core) ConsultationResponse {
	result := ConsultationResponse{
		ID:                   gorm.ID,
		UserID:               user.ID,
		UserFullName:         user.FullName,
		UserProfilePicture:   user.ProfilePicture,
		DoctorID:             doctor.ID,
		DoctorFullName:       doctor.FullName,
		DoctorProfilePicture: doctor.ProfilePicture,
		Consultation:         gorm.Consultation,
		TransactionStatus:    gorm.TransactionStatus,
		StatusConsultation:   gorm.StatusConsultation,
	}
	return result
}
