package data

import (
	"PetPalApp/features/chat"
	"time"

	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	ConsultationID uint
	UserID         uint
	AdminID        uint
	Message        string
	TimeStamp      time.Time
}

func ToCore(c Chat) chat.ChatCore {
	return chat.ChatCore{
		ID:             c.ID,
		ConsultationID: c.ConsultationID,
		UserID:         c.UserID,
		AdminID:        c.AdminID,
		Message:        c.Message,
		TimeStamp:      c.TimeStamp,
	}
}
