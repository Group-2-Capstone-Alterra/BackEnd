package data

import (
	"PetPalApp/features/chat"
	"time"

	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	ConsultationID uint
	SenderID       uint
	ReceiverID     uint
	Message        string
	TimeStamp      time.Time
}

func ToCore(c Chat) chat.ChatCore {
	return chat.ChatCore{
		ID:             c.ID,
		ConsultationID: c.ConsultationID,
		SenderID:       c.SenderID,
		ReceiverID:     c.ReceiverID,
		Message:        c.Message,
		TimeStamp:      c.TimeStamp,
	}
}
