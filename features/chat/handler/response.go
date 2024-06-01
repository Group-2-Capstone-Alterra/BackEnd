package handler

import (
	"PetPalApp/features/chat"
	"time"
)

type ChatResponse struct {
	ID             uint      `json:"id"`
	ConsultationID uint      `json:"roomchat_id"`
	SenderID       uint      `json:"sender_id"`
	ReceiverID     uint      `json:"receiver_id"`
	Message        string    `json:"message"`
	TimeStamp      time.Time `json:"time_stamp"`
}

func AllResponseChat(gorm chat.ChatCore) ChatResponse {
	result := ChatResponse{
		ID:             gorm.ID,
		ConsultationID: gorm.ConsultationID,
		SenderID:       gorm.SenderID,
		ReceiverID:     gorm.ReceiverID,
		Message:        gorm.Message,
		TimeStamp:      gorm.TimeStamp,
	}
	return result
}
