package handler

import (
	"PetPalApp/app/middlewares"
	"PetPalApp/features/admin"
	"PetPalApp/features/chat"
	"PetPalApp/features/consultation"
	"PetPalApp/features/doctor"
	"PetPalApp/features/user"
	"PetPalApp/utils/responses"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ChatHandler struct {
	chatService      chat.ServiceInterface
	consultationData consultation.ConsultationModel
	userData         user.DataInterface
	doctorData       doctor.DoctorModel
	adminData        admin.AdminModel
}

func New(cs chat.ServiceInterface, consultationData consultation.ConsultationModel, userData user.DataInterface, doctorData doctor.DoctorModel, adminData admin.AdminModel) *ChatHandler {
	return &ChatHandler{
		chatService:      cs,
		consultationData: consultationData,
		userData:         userData,
		doctorData:       doctorData,
		adminData:        adminData,
	}
}

func (ch *ChatHandler) CreateChat(c echo.Context) error {
	log.Println("[Handler]")
	roomchatID := c.Param("id")
	roomchatIDConv, errConv := strconv.Atoi(roomchatID)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("ID must be a positive integer", errConv))
	}

	senderID, role, _ := middlewares.ExtractTokenUserId(c)
	if senderID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("Unauthorized", nil))
	}

	newChat := ChatRequest{}
	if err := c.Bind(&newChat); err != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Error binding data: "+err.Error(), nil))
	}

	chatData := chat.ChatCore{
		ConsultationID: uint(roomchatIDConv),
		SenderID:       uint(senderID),
		Message:        newChat.Message,
	}

	if err := ch.chatService.CreateChat(chatData, role); err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Error creating chat: "+err.Error(), nil))
	}

	return c.JSON(http.StatusCreated, responses.JSONWebResponse("Chat created successfully", nil))
}

func (ch *ChatHandler) GetChats(c echo.Context) error {
	log.Println("[Handler]")
	roomchatID := c.Param("id")
	roomchatIDConv, errConv := strconv.Atoi(roomchatID)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("ID must be a positive integer", errConv))
	}
	log.Println("[Handler] Roomchat ID", roomchatIDConv)

	currentID, role, _ := middlewares.ExtractTokenUserId(c)
	if currentID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("Unauthorized", nil))
	}
	log.Println("[Handler] Current User ID", currentID)

	chats, err := ch.chatService.GetChats(uint(currentID), role, uint(roomchatIDConv))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Error retrieving chats: "+err.Error(), nil))
	}

	var allChat []ChatResponse
	for _, v := range chats {
		log.Println("[Handler] for _, v")
		consulData, _ := ch.consultationData.GetCuntationsById(uint(roomchatIDConv))
		// log.Println("[Handler - Chats] consulData", consulData)

		// log.Println("[Handler - Chats] doctorData", doctorData)
		if v.SenderID == consulData.DoctorID { //if sender doctor
			log.Println("[Handler] if sender doctor")
			doctorData, _ := ch.doctorData.SelectDoctorById(v.SenderID)
			userData, _ := ch.userData.SelectById(v.ReceiverID)
			allChat = append(allChat, AllResponseChatFromDoctor(v, *userData, *doctorData))
		} else {
			log.Println("[Handler] if sender user")
			userData, _ := ch.userData.SelectById(v.SenderID)
			doctorData, _ := ch.doctorData.SelectDoctorById(v.ReceiverID)
			allChat = append(allChat, AllResponseChatFromUser(v, *userData, *doctorData))
		}
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse("Chats retrieved successfully", allChat))
}

func (ch *ChatHandler) Delete(c echo.Context) error {
	roomChatID := c.Param("id")
	roomChatIDConv, errRoomChatIDConv := strconv.Atoi(roomChatID)
	if errRoomChatIDConv != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("ID must be a positive integer", roomChatIDConv))
	}
	bubbleChat := c.QueryParam("bubble")
	bubbleChatInt, errBubleChatInt := strconv.Atoi(bubbleChat)
	if errBubleChatInt != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("ID bubble chat must be a positive integer", errBubleChatInt))
	}
	currentID, role, _ := middlewares.ExtractTokenUserId(c)

	err := ch.chatService.Delete(uint(roomChatIDConv), uint(bubbleChatInt), uint(currentID), role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Failed to delete chat, please ensure the room chat ID and bubble chat are correct.", nil))
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse("Chat deleted successfully", nil))
}
