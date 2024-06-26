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

func New(db *gorm.DB) chat.ChatModel {
	return &ChatModel{
		db: db,
	}
}

const (
	qconsulID = "consultation_id = ?"
	qsenderID = "sender_id = ?"
)

func (cm *ChatModel) CreateChat(chat chat.ChatCore) error {
	chatGorm := ToGorm(chat)
	tx := cm.db.Create(&chatGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (cm *ChatModel) GetChatsUser(currentID, roomchatID uint) ([]chat.ChatCore, error) {
	var chats []Chat
	tx := cm.db.Where(qconsulID, roomchatID).Find(&chats)
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

	tx := cm.db.Where(qconsulID, roomchatID).Find(&chats)
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
	tx := cm.db.Where(qconsulID, roomChatID).Where(qsenderID, senderID).Find(&chatData, bubbleChatID)
	if tx.Error != nil {
		return nil, fmt.Errorf("BubbleChat not match with consultation and sender id")
	}
	conculDataCore := ToCore(chatData)
	if conculDataCore.ID == 0 {
		return nil, fmt.Errorf("BubbleChat not match with consultation and sender id")
	} else {
		log.Println("BubbleChat found and match with consultation and sender id")
	}
	return &conculDataCore, nil
}

func (cm *ChatModel) Delete(roomChatID, bubbleChatID, senderID uint) error {
	tx := cm.db.Where(qconsulID, roomChatID).Where(qsenderID, senderID).Delete(&Chat{}, bubbleChatID)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
