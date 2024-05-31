package data

import (
	"PetPalApp/features/chat"

	"gorm.io/gorm"
)

type ChatModel struct {
    db *gorm.DB
}

func New(db *gorm.DB) chat.ChatModel {
    return &ChatModel{
        db: db,
    }
}

func (cm *ChatModel) CreateChat(chat chat.ChatCore) error {
    chatGorm := Chat{
        SenderID:   chat.SenderID,
        ReceiverID: chat.ReceiverID,
        Message:    chat.Message,
    }
    tx := cm.db.Create(&chatGorm)
    if tx.Error != nil {
        return tx.Error
    }
    return nil
}
