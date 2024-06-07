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
	SenderID            uint      `json:"sender_id"`
	SenderName          string    `json:"sender_name"`
	SenderProfilePict   string    `json:"sender_profilepict"`
	ReceiverID          uint      `json:"receiver_id"`
	ReceiverName        string    `json:"receiver_name"`
	ReceiverProfilePict string    `json:"receiver_profilepict"`
	Message             string    `json:"message"`
	TimeStamp           time.Time `json:"time_stamp"`
}

func AllResponseChatFromUser(gorm chat.ChatCore, user user.Core, doctor doctor.Core) ChatResponse {
	result := ChatResponse{
		ID:                  gorm.ID,
		ConsultationID:      gorm.ConsultationID,
		SenderID:            gorm.SenderID,
		SenderName:          user.FullName,
		SenderProfilePict:   user.ProfilePicture,
		ReceiverID:          doctor.AdminID,
		ReceiverName:        doctor.FullName,
		ReceiverProfilePict: doctor.ProfilePicture,
		Message:             gorm.Message,
		TimeStamp:           gorm.TimeStamp,
	}
	return result
}

func AllResponseChatFromDoctor(gorm chat.ChatCore, user user.Core, doctor doctor.Core) ChatResponse {
	result := ChatResponse{
		ID:                  gorm.ID,
		ConsultationID:      gorm.ConsultationID,
		SenderID:            doctor.AdminID,
		SenderName:          doctor.FullName,
		SenderProfilePict:   doctor.ProfilePicture,
		ReceiverID:          gorm.ReceiverID,
		ReceiverName:        user.FullName,
		ReceiverProfilePict: user.ProfilePicture,
		Message:             gorm.Message,
		TimeStamp:           gorm.TimeStamp,
	}
	return result
}
