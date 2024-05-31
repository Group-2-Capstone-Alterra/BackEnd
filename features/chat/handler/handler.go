package handler

import (
	"PetPalApp/app/middlewares"
	"PetPalApp/features/chat"
	"PetPalApp/utils/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ChatHandler struct {
    chatService chat.ChatService
}

func New(cs chat.ChatService) *ChatHandler {
    return &ChatHandler{
        chatService: cs,
    }
}

func (ch *ChatHandler) CreateChat(c echo.Context) error {
    senderID := middlewares.ExtractTokenUserId(c)
    if senderID == 0 {
        return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("Unauthorized", nil))
    }

    newChat := ChatRequest{}
    if err := c.Bind(&newChat); err != nil {
        return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Error binding data: "+err.Error(), nil))
    }

    chatData := chat.ChatCore{
        SenderID:   uint(senderID),
        ReceiverID: newChat.ReceiverID,
        Message:    newChat.Message,
    }

    if err := ch.chatService.CreateChat(chatData); err != nil {
        return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Error creating chat: "+err.Error(), nil))
    }

    return c.JSON(http.StatusCreated, responses.JSONWebResponse("Chat created successfully", nil))
}

func (ch *ChatHandler) GetChats(c echo.Context) error {
    senderID := middlewares.ExtractTokenUserId(c)
    if senderID == 0 {
        return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("Unauthorized", nil))
    }

    receiverID, err := strconv.ParseUint(c.QueryParam("receiver_id"), 10, 32)
    if err != nil {
        return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Invalid receiver ID", nil))
    }

    chats, err := ch.chatService.GetChats(uint(senderID), uint(receiverID))
    if err != nil {
        return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Error retrieving chats: "+err.Error(), nil))
    }

    return c.JSON(http.StatusOK, responses.JSONWebResponse("Chats retrieved successfully", chats))
}