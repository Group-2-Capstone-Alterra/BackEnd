package handler

import (
	"PetPalApp/features/chat"
	"PetPalApp/features/doctor"
	"PetPalApp/features/user"
	"time"
)

type ChatResponse struct {
	ID             uint                 `json:"id"`
	ConsultationID uint                 `json:"roomchat_id"`
	Sender         SenderInformations   `json:"sender"`
	Receiver       ReceiverInformations `json:"receiver"`
	Message        string               `json:"message"`
	TimeStamp      time.Time            `json:"time_stamp"`
}

type SenderInformations struct {
	ID             uint   `json:"id"`
	Role           string `json:"role"`
	FullName       string `json:"full_name"`
	ProfilePicture string `json:"profile_picture"`
}

type ReceiverInformations struct {
	ID             uint   `json:"id"`
	Role           string `json:"role"`
	FullName       string `json:"full_name"`
	ProfilePicture string `json:"profile_picture"`
}

func AllResponseChatFromUser(gorm chat.ChatCore, user user.Core, doctor doctor.Core) ChatResponse {
	result := ChatResponse{
		ID:             gorm.ID,
		ConsultationID: gorm.ConsultationID,
		Sender: SenderInformations{
			ID:             gorm.SenderID,
			Role:           user.Role,
			FullName:       user.FullName,
			ProfilePicture: user.ProfilePicture,
		},
		Receiver: ReceiverInformations{
			ID:             doctor.AdminID,
			Role:           "doctor",
			FullName:       doctor.FullName,
			ProfilePicture: gorm.ReceiverName,
		},
		Message:   gorm.Message,
		TimeStamp: gorm.TimeStamp,
	}
	return result
}

func AllResponseChatFromDoctor(gorm chat.ChatCore, user user.Core, doctor doctor.Core) ChatResponse {
	result := ChatResponse{
		ID:             gorm.ID,
		ConsultationID: gorm.ConsultationID,
		Sender: SenderInformations{
			ID:             doctor.AdminID,
			Role:           "doctor",
			FullName:       doctor.FullName,
			ProfilePicture: doctor.ProfilePicture,
		},
		Receiver: ReceiverInformations{
			ID:             gorm.ReceiverID,
			Role:           user.Role,
			FullName:       user.FullName,
			ProfilePicture: user.ProfilePicture,
		},
		Message:   gorm.Message,
		TimeStamp: gorm.TimeStamp,
	}
	return result
}
