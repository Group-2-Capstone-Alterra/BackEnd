package service_test

import (
	"PetPalApp/features/chat"
	"PetPalApp/features/chat/service"
	"PetPalApp/features/consultation"
	"PetPalApp/features/doctor"
	"PetPalApp/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateChat_User(t *testing.T) {
	mockChatModel := new(mocks.ChatModel)
	mockConsultationModel := new(mocks.ConsultationModel)
	mockDoctorModel := new(mocks.DoctorModel)
	mockUserModel := new(mocks.UserModel)
	mockAdminModel := new(mocks.AdminModel)

	chatService := service.New(mockChatModel, mockConsultationModel, mockDoctorModel, mockUserModel, mockAdminModel)

	inputChat := chat.ChatCore{
		SenderID:       1,
		ConsultationID: 1,
		Message:        "Hello, Doctor!",
	}

	consultationData := &consultation.ConsultationCore{
		ID:       1,
		DoctorID: 2,
		UserID:   1,
	}

	mockConsultationModel.On("GetCuntationsById", inputChat.ConsultationID).Return(consultationData, nil)
	mockConsultationModel.On("VerUser", inputChat.SenderID, consultationData.DoctorID, inputChat.ConsultationID).Return(consultationData, nil)
	mockChatModel.On("CreateChat", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		chat := args.Get(0).(chat.ChatCore)
		chat.TimeStamp = time.Now() // Ensure timestamp is set in the test
	})

	err := chatService.CreateChat(inputChat, "user")

	assert.Nil(t, err)
	mockConsultationModel.AssertExpectations(t)
	mockChatModel.AssertExpectations(t)
}

func TestCreateChat_Admin(t *testing.T) {
	mockChatModel := new(mocks.ChatModel)
	mockConsultationModel := new(mocks.ConsultationModel)
	mockDoctorModel := new(mocks.DoctorModel)
	mockUserModel := new(mocks.UserModel)
	mockAdminModel := new(mocks.AdminModel)

	chatService := service.New(mockChatModel, mockConsultationModel, mockDoctorModel, mockUserModel, mockAdminModel)

	inputChat := chat.ChatCore{
		SenderID:       1,
		ConsultationID: 1,
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

	mockConsultationModel.On("GetCuntationsById", inputChat.ConsultationID).Return(consultationData, nil)
	mockDoctorModel.On("SelectByAdminId", inputChat.SenderID).Return(doctorData, nil)
	mockConsultationModel.On("VerAdmin", doctorData.ID, consultationData.UserID, inputChat.ConsultationID).Return(consultationData, nil)
	mockChatModel.On("CreateChat", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		chat := args.Get(0).(chat.ChatCore)
		chat.TimeStamp = time.Now() // Ensure timestamp is set in the test
	})

	err := chatService.CreateChat(inputChat, "admin")

	assert.Nil(t, err)
	mockConsultationModel.AssertExpectations(t)
	mockChatModel.AssertExpectations(t)
}

func TestGetChats_UserRole(t *testing.T) {
	mockChatModel := new(mocks.ChatModel)
	mockConsultationModel := new(mocks.ConsultationModel)
	mockDoctorModel := new(mocks.DoctorModel)
	mockUserModel := new(mocks.UserModel)
	mockAdminModel := new(mocks.AdminModel)

	chatService := service.New(mockChatModel, mockConsultationModel, mockDoctorModel, mockUserModel, mockAdminModel)

	currentID := uint(1)
	role := "user"
	roomchatID := uint(1)
	expectedChats := []chat.ChatCore{
		{ID: 1, SenderID: 1, ReceiverID: 2, Message: "Hello, Doctor!"},
		{ID: 2, SenderID: 2, ReceiverID: 1, Message: "Hello, User!"},
	}

	consultationData := &consultation.ConsultationCore{
		ID:       1,
		DoctorID: 2,
		UserID:   1,
	}

	mockConsultationModel.On("GetCuntationsById", roomchatID).Return(consultationData, nil)
	mockConsultationModel.On("VerUser", currentID, consultationData.DoctorID, roomchatID).Return(consultationData, nil)
	mockChatModel.On("GetChatsUser", currentID, roomchatID).Return(expectedChats, nil)

	result, err := chatService.GetChats(currentID, role, roomchatID)

	assert.Nil(t, err)
	assert.Equal(t, expectedChats, result)
	mockConsultationModel.AssertExpectations(t)
	mockChatModel.AssertExpectations(t)
}

func TestGetChats_AdminRole(t *testing.T) {
	mockChatModel := new(mocks.ChatModel)
	mockConsultationModel := new(mocks.ConsultationModel)
	mockDoctorModel := new(mocks.DoctorModel)
	mockUserModel := new(mocks.UserModel)
	mockAdminModel := new(mocks.AdminModel)

	chatService := service.New(mockChatModel, mockConsultationModel, mockDoctorModel, mockUserModel, mockAdminModel)

	currentID := uint(1)
	role := "admin"
	roomchatID := uint(1)
	expectedChats := []chat.ChatCore{
		{ID: 1, SenderID: 1, ReceiverID: 2, Message: "Hello, Doctor!"},
		{ID: 2, SenderID: 2, ReceiverID: 1, Message: "Hello, User!"},
	}

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
}

func TestDeleteChat(t *testing.T) {
	mockChatModel := new(mocks.ChatModel)
	mockConsultationModel := new(mocks.ConsultationModel)
	mockDoctorModel := new(mocks.DoctorModel)
	mockUserModel := new(mocks.UserModel)
	mockAdminModel := new(mocks.AdminModel)

	chatService := service.New(mockChatModel, mockConsultationModel, mockDoctorModel, mockUserModel, mockAdminModel)

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

	chatService := service.New(mockChatModel, mockConsultationModel, mockDoctorModel, mockUserModel, mockAdminModel)

	roomChatID := uint(1)
	bubbleChatID := uint(1)
	senderID := uint(1)
	role := "admin"

	doctorData := &doctor.Core{
		ID: 2,
	}

	chatData := &chat.ChatCore{
		ID: 1,
	}

	mockDoctorModel.On("SelectByAdminId", senderID).Return(doctorData, nil)
	mockChatModel.On("VerAvailChat", roomChatID, bubbleChatID, doctorData.ID).Return(chatData, nil)
	mockChatModel.On("Delete", roomChatID, bubbleChatID, doctorData.ID).Return(nil)

	err := chatService.Delete(roomChatID, bubbleChatID, senderID, role)

	assert.Nil(t, err)
	mockChatModel.AssertExpectations(t)
	mockDoctorModel.AssertExpectations(t)
}
