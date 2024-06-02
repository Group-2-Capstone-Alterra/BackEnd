package chat

import "time"

type ChatCore struct {
	ID             uint
	ConsultationID uint
	SenderID       uint
	ReceiverID     uint
	Message        string
	TimeStamp      time.Time
	CreatedAt      time.Time
	UpdateAt       time.Time
}

type DataInterface interface {
	CreateChat(ChatCore) error
	GetChats(roomchatID uint) ([]ChatCore, error)
	VerAvailChat(roomChatID, bubbleChatID, senderID uint) (*ChatCore, error)
	Delete(roomChatID, bubbleChatID, senderID uint) error
}

type ServiceInterface interface {
	CreateChat(ChatCore ChatCore) error
	GetChats(currentID, roomchatID uint) ([]ChatCore, error)
	Delete(roomChatID, bubbleChatID, senderID uint) error
}
