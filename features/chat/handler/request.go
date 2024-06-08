package handler

import "PetPalApp/features/chat"

type ChatRequest struct {
	Message string `json:"message" form:"message"`
}

func ReqToCore(senderID, roomchatID uint, c ChatRequest) chat.ChatCore {
	return chat.ChatCore{
		ConsultationID: uint(roomchatID),
		SenderID:       uint(senderID),
		Message:        c.Message,
	}
}
