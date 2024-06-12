package service_test

import (
	"PetPalApp/features/chat"
	"PetPalApp/features/chat/service"
	"PetPalApp/features/consultation"
	"PetPalApp/features/doctor"
	"PetPalApp/mocks"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateChatUser(t *testing.T) {
	mockChatModel := new(mocks.ChatModel)
	mockConsultationModel := new(mocks.ConsultationModel)
	mockDoctorModel := new(mocks.DoctorModel)
	mockUserModel := new(mocks.UserModel)
	mockAdminModel := new(mocks.AdminModel)
	mockConsultationService := new(mocks.ConsultationService)

	chatService := service.New(mockChatModel, mockConsultationModel, mockDoctorModel, mockUserModel, mockAdminModel, mockConsultationService)

	inputChatNonEmpty := chat.ChatCore{
		SenderID:       1,
		ConsultationID: 1,
		Message:        "Hello, Doctor!",
	}

	inputChatEmpty := chat.ChatCore{
		SenderID:       1,
		ConsultationID: 1,
		Message:        "",
	}

	inputConsulNotFound := chat.ChatCore{
		SenderID:       1,
		ConsultationID: 3,
		Message:        "Hello, Doctor!",
	}

	inputConsulNotMatch := chat.ChatCore{
		SenderID:       2,
		ConsultationID: 1,
		Message:        "Hello, Doctor!",
	}

	consultationData := &consultation.ConsultationCore{
		ID:       1,
		DoctorID: 2,
		UserID:   1,
	}

	//normal
	mockConsultationService.On("GetCuntationsById", inputChatNonEmpty.ConsultationID).Return(consultationData, nil)
	mockConsultationModel.On("VerUser", inputChatNonEmpty.SenderID, consultationData.DoctorID, inputChatNonEmpty.ConsultationID).Return(consultationData, nil)
	mockChatModel.On("CreateChat", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		chat := args.Get(0).(chat.ChatCore)
		chat.TimeStamp = time.Now() // Ensure timestamp is set in the test
	})
	errNonEmptyChat := chatService.CreateChat(inputChatNonEmpty, "user")
	assert.Nil(t, errNonEmptyChat)

	//message not found
	mockConsultationService.On("GetCuntationsById", inputChatEmpty.ConsultationID).Return(consultationData, nil)
	mockConsultationModel.On("VerUser", inputChatEmpty.SenderID, consultationData.DoctorID, inputChatEmpty.ConsultationID).Return(consultationData, nil)
	mockChatModel.On("CreateChat", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		chat := args.Get(0).(chat.ChatCore)
		chat.TimeStamp = time.Now() // Ensure timestamp is set in the test
	})
	errEmptyChat := chatService.CreateChat(inputChatEmpty, "user")
	assert.NotNil(t, errEmptyChat)
	assert.Equal(t, "message cannot be sent empty", errEmptyChat.Error())

	//roomchat not found
	mockConsultationService.On("GetCuntationsById", inputConsulNotFound.ConsultationID).Return(nil, errors.New("roomchat not found"))
	errConsulNotfound := chatService.CreateChat(inputConsulNotFound, "user")
	assert.NotNil(t, errConsulNotfound)
	assert.Equal(t, "roomchat not found", errConsulNotfound.Error())

	//UserID and DoctorID not match at Roomchat
	mockConsultationService.On("GetCuntationsById", inputConsulNotMatch.ConsultationID).Return(consultationData, nil)
	mockConsultationModel.On("VerUser", inputConsulNotMatch.SenderID, consultationData.DoctorID, inputConsulNotMatch.ConsultationID).Return(nil, errors.New("UserID and DoctorID not match at Roomchat"))
	mockChatModel.On("CreateChat", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		chat := args.Get(0).(chat.ChatCore)
		chat.TimeStamp = time.Now() // Ensure timestamp is set in the test
	})
	errVerUser := chatService.CreateChat(inputConsulNotMatch, "user")
	assert.NotNil(t, errVerUser)
	assert.Equal(t, "UserID and DoctorID not match at Roomchat", errVerUser.Error())

	mockConsultationModel.AssertExpectations(t)
	mockChatModel.AssertExpectations(t)
}

func TestCreateChatAdmin(t *testing.T) {
	mockChatModel := new(mocks.ChatModel)
	mockConsultationModel := new(mocks.ConsultationModel)
	mockDoctorModel := new(mocks.DoctorModel)
	mockUserModel := new(mocks.UserModel)
	mockAdminModel := new(mocks.AdminModel)
	mockConsultationService := new(mocks.ConsultationService)

	chatService := service.New(mockChatModel, mockConsultationModel, mockDoctorModel, mockUserModel, mockAdminModel, mockConsultationService)

	inputChat := chat.ChatCore{
		SenderID:       1,
		ConsultationID: 1,
		Message:        "Hello, User!",
	}

	inputChatConsultationModel := chat.ChatCore{
		SenderID:       1,
		ConsultationID: 3,
		Message:        "Hello, User!",
	}

	consultationData := &consultation.ConsultationCore{
		ID:       1,
		DoctorID: 2,
		UserID:   1,
	}

	doctorData := &doctor.Core{
		ID: 2,
	}

	//normal
	mockConsultationService.On("GetCuntationsById", inputChat.ConsultationID).Return(consultationData, nil)
	mockDoctorModel.On("SelectByAdminId", inputChat.SenderID).Return(doctorData, nil)
	mockConsultationModel.On("VerAdmin", doctorData.ID, consultationData.UserID, inputChat.ConsultationID).Return(consultationData, nil)
	mockChatModel.On("CreateChat", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		chat := args.Get(0).(chat.ChatCore)
		chat.TimeStamp = time.Now() // Ensure timestamp is set in the test
	})

	err := chatService.CreateChat(inputChat, "admin")
	assert.Nil(t, err)

	//roomchat not found
	mockConsultationService.On("GetCuntationsById", inputChatConsultationModel.ConsultationID).Return(nil, errors.New("roomchat not found"))
	errConsulNotfound := chatService.CreateChat(inputChatConsultationModel, "admin")
	assert.NotNil(t, errConsulNotfound)
	assert.Equal(t, "roomchat not found", errConsulNotfound.Error())

	mockConsultationService.AssertExpectations(t)
	mockChatModel.AssertExpectations(t)
	mockConsultationModel.AssertExpectations(t) // Add this line
}

func TestGetChatsUserRole(t *testing.T) {
	mockChatModel := new(mocks.ChatModel)
	mockConsultationModel := new(mocks.ConsultationModel)
	mockDoctorModel := new(mocks.DoctorModel)
	mockUserModel := new(mocks.UserModel)
	mockAdminModel := new(mocks.AdminModel)
	mockConsultationService := new(mocks.ConsultationService)

	chatService := service.New(mockChatModel, mockConsultationModel, mockDoctorModel, mockUserModel, mockAdminModel, mockConsultationService)

	currentID := uint(1)
	role := "user"
	roomchatID := uint(1)
	expectedChats := []chat.ChatCore{
		{ID: 1, SenderID: 1, ReceiverID: 2, Message: "Hello, Doctor!"},
		{ID: 2, SenderID: 2, ReceiverID: 1, Message: "Hello, User!"},
	}

	currentIDWrong := uint(1)
	roleWrong := "user"
	roomchatIDWrong := uint(1)

	consultationData := &consultation.ConsultationCore{
		ID:       1,
		DoctorID: 2,
		UserID:   1,
	}

	//normal
	mockConsultationModel.On("GetCuntationsById", roomchatID).Return(consultationData, nil)
	mockConsultationModel.On("VerUser", currentID, consultationData.DoctorID, roomchatID).Return(consultationData, nil)
	mockChatModel.On("GetChatsUser", currentID, roomchatID).Return(expectedChats, nil)
	result, err := chatService.GetChats(currentID, role, roomchatID)
	assert.Nil(t, err)
	assert.Equal(t, expectedChats, result)

	mockConsultationModel.AssertExpectations(t)
	mockChatModel.AssertExpectations(t)

	// Reset mock expectations
	mockConsultationModel.ExpectedCalls = nil
	mockChatModel.ExpectedCalls = nil

	//normal
	mockConsultationModel.On("GetCuntationsById", roomchatIDWrong).Return(nil, errors.New("Roomchat Not Found"))
	_, errCreate := chatService.GetChats(currentIDWrong, roleWrong, roomchatIDWrong)
	assert.NotNil(t, errCreate)
	assert.Equal(t, "RoomChat Not Found", errCreate.Error())

	mockConsultationModel.AssertExpectations(t)
	mockChatModel.AssertExpectations(t)
}

func TestGetChatsAdminRole(t *testing.T) {
	mockChatModel := new(mocks.ChatModel)
	mockConsultationModel := new(mocks.ConsultationModel)
	mockDoctorModel := new(mocks.DoctorModel)
	mockUserModel := new(mocks.UserModel)
	mockAdminModel := new(mocks.AdminModel)
	mockConsultationService := new(mocks.ConsultationService)

	chatService := service.New(mockChatModel, mockConsultationModel, mockDoctorModel, mockUserModel, mockAdminModel, mockConsultationService)

	currentID := uint(1)
	role := "admin"
	roomchatID := uint(1)
	expectedChats := []chat.ChatCore{
		{ID: 1, SenderID: 1, ReceiverID: 2, Message: "Hello, Doctor!"},
		{ID: 2, SenderID: 2, ReceiverID: 1, Message: "Hello, User!"},
	}

	currentIDWrong := uint(1)
	roleWrong := "admin"
	roomchatIDWrong := uint(1)

	doctorData := &doctor.Core{
		ID: 2,
	}

	consultationData := &consultation.ConsultationCore{
		ID:       1,
		DoctorID: 2,
		UserID:   1,
	}

	mockDoctorModel.On("SelectByAdminId", currentID).Return(doctorData, nil)
	mockConsultationModel.On("GetCuntationsById", roomchatID).Return(consultationData, nil)
	mockConsultationModel.On("VerAdmin", doctorData.ID, consultationData.UserID, roomchatID).Return(consultationData, nil)
	mockChatModel.On("GetChatsDoctor", roomchatID).Return(expectedChats, nil)

	result, err := chatService.GetChats(currentID, role, roomchatID)

	assert.Nil(t, err)
	assert.Equal(t, expectedChats, result)
	mockConsultationModel.AssertExpectations(t)
	mockChatModel.AssertExpectations(t)

	// Reset mock expectations
	mockConsultationModel.ExpectedCalls = nil
	mockChatModel.ExpectedCalls = nil

	mockConsultationModel.On("GetCuntationsById", roomchatIDWrong).Return(nil, errors.New("Roomchat Not Found"))
	_, errCreate := chatService.GetChats(currentIDWrong, roleWrong, roomchatIDWrong)
	assert.NotNil(t, errCreate)
	assert.Equal(t, "RoomChat Not Found", errCreate.Error())

	mockConsultationModel.AssertExpectations(t)
	mockChatModel.AssertExpectations(t)
}

func TestDeleteChat(t *testing.T) {
	mockChatModel := new(mocks.ChatModel)
	mockConsultationModel := new(mocks.ConsultationModel)
	mockDoctorModel := new(mocks.DoctorModel)
	mockUserModel := new(mocks.UserModel)
	mockAdminModel := new(mocks.AdminModel)
	mockConsultationService := new(mocks.ConsultationService)

	chatService := service.New(mockChatModel, mockConsultationModel, mockDoctorModel, mockUserModel, mockAdminModel, mockConsultationService)

	roomChatID := uint(1)
	bubbleChatID := uint(1)
	senderID := uint(1)
	role := "user"

	chatData := &chat.ChatCore{
		ID: 1,
	}

	mockChatModel.On("VerAvailChat", roomChatID, bubbleChatID, senderID).Return(chatData, nil)
	mockChatModel.On("Delete", roomChatID, bubbleChatID, senderID).Return(nil)

	err := chatService.Delete(roomChatID, bubbleChatID, senderID, role)

	assert.Nil(t, err)
	mockChatModel.AssertExpectations(t)
}

func TestDeleteChat_Admin(t *testing.T) {
	mockChatModel := new(mocks.ChatModel)
	mockConsultationModel := new(mocks.ConsultationModel)
	mockDoctorModel := new(mocks.DoctorModel)
	mockUserModel := new(mocks.UserModel)
	mockAdminModel := new(mocks.AdminModel)
	mockConsultationService := new(mocks.ConsultationService)

	chatService := service.New(mockChatModel, mockConsultationModel, mockDoctorModel, mockUserModel, mockAdminModel, mockConsultationService)

	roomChatID := uint(1)
	roomChatIDWrong := uint(123)
	bubbleChatID := uint(1)
	senderID := uint(1)
	role := "admin"

	doctorData := &doctor.Core{
		ID: 2,
	}

	chatData := &chat.ChatCore{
		ID: 1,
	}

	// Test case: roomChatID <= 0
	errInvalidID := chatService.Delete(0, bubbleChatID, senderID, role)
	assert.NotNil(t, errInvalidID)
	assert.Equal(t, "ID must be a positive integer", errInvalidID.Error())

	// Test case: valConcul == nil
	mockDoctorModel.On("SelectByAdminId", senderID).Return(doctorData, nil)
	mockChatModel.On("VerAvailChat", roomChatID, bubbleChatID, doctorData.ID).Return(nil, nil)
	errNotFoundRoom := chatService.Delete(roomChatID, bubbleChatID, senderID, role)
	assert.NotNil(t, errNotFoundRoom)
	assert.Contains(t, errNotFoundRoom.Error(), "roomchat not found")

	// Test case: roomChatIDWrong
	mockDoctorModel.On("SelectByAdminId", senderID).Return(doctorData, nil)
	mockChatModel.On("VerAvailChat", roomChatIDWrong, bubbleChatID, doctorData.ID).Return(chatData, nil)
	mockChatModel.On("Delete", roomChatIDWrong, bubbleChatID, doctorData.ID).Return(errors.New("roomchat not found"))
	errNotFoundRoom = chatService.Delete(roomChatIDWrong, bubbleChatID, senderID, role)
	assert.NotNil(t, errNotFoundRoom)
	assert.Contains(t, errNotFoundRoom.Error(), "roomchat not found")

	mockChatModel.AssertExpectations(t)
	mockDoctorModel.AssertExpectations(t)
}
