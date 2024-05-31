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

func (cm *ChatModel) GetChats(senderID, receiverID uint) ([]chat.ChatCore, error) {
    var chats []Chat
    tx := cm.db.Where("sender_id = ? AND receiver_id = ?", senderID, receiverID).
        Or("sender_id = ? AND receiver_id = ?", receiverID, senderID).Find(&chats)
    if tx.Error != nil {
        return nil, tx.Error
    }

    var result []chat.ChatCore
    for _, chat := range chats {
        result = append(result, chat.ToCore())
    }

    return result, nil
}
