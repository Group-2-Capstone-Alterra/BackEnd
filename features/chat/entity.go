package chat

type ChatCore struct {
	SenderID   uint
	ReceiverID uint
	Message    string
}

type ChatModel interface {
	CreateChat(ChatCore) error
}

type ChatService interface {
	CreateChat(ChatCore) error
}
