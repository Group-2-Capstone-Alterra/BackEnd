package data

import (
	"PetPalApp/features/chat"
	"fmt"
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
	log.Println("[Query - Create Chat] Chat Created")
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
	return nil
}

func (cm *ChatModel) GetChatsUser(currentID, roomchatID uint) ([]chat.ChatCore, error) {
	var chats []Chat
	tx := cm.db.Where("consultation_id = ?", roomchatID).Find(&chats)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var result []chat.ChatCore
	for _, chat := range chats {

		result = append(result, ToCore(chat))
	}

	return result, nil
}

func (cm *ChatModel) GetChatsDoctor(roomchatID uint) ([]chat.ChatCore, error) {
	var chats []Chat

	tx := cm.db.Where("consultation_id = ?", roomchatID).Find(&chats)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var result []chat.ChatCore
	for _, chat := range chats {
		result = append(result, ToCore(chat))
	}

	return result, nil
}

func (cm *ChatModel) VerAvailChat(roomChatID, bubbleChatID, senderID uint) (*chat.ChatCore, error) {
	var chatData Chat
	tx := cm.db.Where("consultation_id = ?", roomChatID).Where("sender_id = ?", senderID).Find(&chatData, bubbleChatID)
	if tx.Error != nil {
		return nil, fmt.Errorf("[Query VerAvailConcul] BubbleChat not match with consultation and sender id")
	}
	conculDataCore := ToCore(chatData)
	if conculDataCore.ID == 0 {
		return nil, fmt.Errorf("[Query VerAvailConcul] BubbleChat not match with consultation and sender id")
	} else {
		log.Println("[Query VerAvailConcul] BubbleChat found and match with consultation and sender id")
	}
	return &conculDataCore, nil
}

func (cm *ChatModel) Delete(roomChatID, bubbleChatID, senderID uint) error {
	tx := cm.db.Where("consultation_id = ?", roomChatID).Where("sender_id = ?", senderID).Delete(&Chat{}, bubbleChatID)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
