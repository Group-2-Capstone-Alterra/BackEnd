package chat

import "time"

type ChatCore struct {
	ID                  uint
	ConsultationID      uint
	SenderName          string
	ReceiverName        string
	SenderID            uint
	ReceiverID          uint
	ReceiverProfilePict string
	SenderProfilePict   string
	Message             string
	TimeStamp           time.Time
	CreatedAt           time.Time
	UpdateAt            time.Time
}

type DataInterface interface {
	CreateChat(ChatCore) error
	GetChatsUser(currentID, roomchatID uint) ([]ChatCore, error)
	GetChatsDoctor(currentID, roomchatID uint) ([]ChatCore, error)
	VerAvailChat(roomChatID, bubbleChatID, senderID uint) (*ChatCore, error)
	Delete(roomChatID, bubbleChatID, senderID uint) error
}

type ServiceInterface interface {
	CreateChat(ChatCore ChatCore, role string) error
	GetChats(currentID uint, role string, roomchatID uint) ([]ChatCore, error)
	Delete(roomChatID, bubbleChatID, senderID uint) error
}
