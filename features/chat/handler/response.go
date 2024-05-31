package handler

type ChatResponse struct {
	ReceiverID uint   `json:"receiver_id"`
	Message    string `json:"message"`
}
