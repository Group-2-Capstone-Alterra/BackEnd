package handler

type ChatRequest struct {
	Message string `json:"message" form:"message"`
}
