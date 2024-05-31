package service

import (
	"PetPalApp/features/chat"
)

type ChatService struct {
    chatModel chat.ChatModel
}

func New(cm chat.ChatModel) chat.ChatService {
    return &ChatService{
        chatModel: cm,
    }
}

func (cs *ChatService) CreateChat(chat chat.ChatCore) error {
    return cs.chatModel.CreateChat(chat)
}

func (cs *ChatService) GetChats(senderID, receiverID uint) ([]chat.ChatCore, error) {
    return cs.chatModel.GetChats(senderID, receiverID)
}
