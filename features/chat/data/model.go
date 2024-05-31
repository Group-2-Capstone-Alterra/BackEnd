package data

import (
	"PetPalApp/features/chat"

	"gorm.io/gorm"
)

type Chat struct {
    gorm.Model
    SenderID   uint   
    ReceiverID uint   
    Message    string 
}

func (c *Chat) ToCore() chat.ChatCore {
    return chat.ChatCore{
        SenderID:   c.SenderID,
        ReceiverID: c.ReceiverID,
        Message:    c.Message,
    }
}