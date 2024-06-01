package data

import (
	"PetPalApp/features/chat"
	"log"

	"gorm.io/gorm"
)

type ChatModel struct {
	db *gorm.DB
}

func New(db *gorm.DB) chat.DataInterface {
	return &ChatModel{
		db: db,
	}
}

func (cm *ChatModel) CreateChat(chat chat.ChatCore) error {
	chatGorm := Chat{
		ConsultationID: chat.ConsultationID,
		SenderID:       chat.SenderID,
		ReceiverID:     chat.ReceiverID,
		Message:        chat.Message,
		TimeStamp:      chat.TimeStamp,
	}
	tx := cm.db.Create(&chatGorm)
	if tx.Error != nil {
		return tx.Error
	}
	log.Println("[Query - Create Chat] Detail Chat", chatGorm)
	return nil
}

func (cm *ChatModel) GetChats(receiverID uint) ([]chat.ChatCore, error) {
	var chats []Chat
	tx := cm.db.Find(&chats)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var result []chat.ChatCore
	for _, chat := range chats {
		result = append(result, chat.ToCore())
	}

	return result, nil
}
