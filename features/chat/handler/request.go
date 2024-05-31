package handler

type ChatRequest struct {
	ReceiverID uint   `json:"receiver_id" form:"receiver_id"`
	Message    string `json:"message" form:"message"`
}
