package chat

import "time"

type ChatCore struct {
	ID                  uint
	ConsultationID      uint
	UserName            string
	AdminName           string
	UserID              uint
	AdminID             uint
	AdminProfilePict    string
	UserProfilePict     string
	Message             string
	TimeStamp           time.Time
	CreatedAt           time.Time
	UpdateAt            time.Time
}

type DataInterface interface {
	CreateChat(ChatCore) error
	GetChatsUser(currentID, roomchatID uint) ([]ChatCore, error)
	GetChatsDoctor(currentID, roomchatID uint) ([]ChatCore, error)
	VerAvailChat(roomChatID, bubbleChatID, userID uint) (*ChatCore, error)
	Delete(roomChatID, bubbleChatID, userID uint) error
}

type ServiceInterface interface {
	CreateChat(ChatCore ChatCore, role string) error
	GetChats(currentID uint, role string, roomchatID uint) ([]ChatCore, error)
	Delete(roomChatID, bubbleChatID, userID uint) error
}
