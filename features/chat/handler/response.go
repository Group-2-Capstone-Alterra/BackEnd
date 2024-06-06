package handler

import (
	"PetPalApp/features/chat"
	"PetPalApp/features/doctor"
	"PetPalApp/features/user"
	"time"
)

type ChatResponse struct {
	ID                  uint      `json:"id"`
	ConsultationID      uint      `json:"roomchat_id"`
	UserID           	uint      `json:"user_id"`
	UserName            string    `json:"user_name"`
	UserProfilePict     string    `json:"user_profilepict"`
	AdminID             uint      `json:"admin_id"`
	AdminName           string    `json:"admin_name"`
	AdminProfilePict    string    `json:"admin_profilepict"`
	Message             string    `json:"message"`
	TimeStamp           time.Time `json:"time_stamp"`
}

func AllResponseChatFromUser(gorm chat.ChatCore, user user.Core, doctor doctor.Core) ChatResponse {
	result := ChatResponse{
		ID:                  gorm.ID,
		ConsultationID:      gorm.ConsultationID,
		UserID:              gorm.UserID,
		UserName:            user.FullName,
		UserProfilePict:     user.ProfilePicture,
		AdminID:             gorm.AdminID,
		AdminName:           doctor.FullName,
		AdminProfilePict:    doctor.ProfilePicture,
		Message:             gorm.Message,
		TimeStamp:           gorm.TimeStamp,
	}
	return result
}

func AllResponseChatFromDoctor(gorm chat.ChatCore, user user.Core, doctor doctor.Core) ChatResponse {
	result := ChatResponse{
		ID:                  gorm.ID,
		ConsultationID:      gorm.ConsultationID,
		AdminID:             gorm.AdminID,
		AdminName:           doctor.FullName,
		AdminProfilePict:    doctor.ProfilePicture,
		UserID:              gorm.UserID,
		UserName:            user.FullName,
		UserProfilePict:     user.ProfilePicture,
		Message:             gorm.Message,
		TimeStamp:           gorm.TimeStamp,
	}
	return result
}
