package service

import (
	"PetPalApp/features/admin"
	"PetPalApp/features/chat"
	"PetPalApp/features/consultation"
	"PetPalApp/features/doctor"
	"PetPalApp/features/user"
	"errors"
	"time"
)

type ChatService struct {
	chatModel        chat.DataInterface
	consultationData consultation.ConsultationModel
	doctorData       doctor.DoctorModel
	userData         user.DataInterface
	adminData        admin.AdminModel
}

func New(cm chat.DataInterface, consultationData consultation.ConsultationModel, doctorData doctor.DoctorModel, userData user.DataInterface, adminData admin.AdminModel) chat.ServiceInterface {
	return &ChatService{
		chatModel:        cm,
		consultationData: consultationData,
		doctorData:       doctorData,
		userData:         userData,
		adminData:        adminData,
	}
}

const (
	notFoundRoomChat = "roomchat not found"
)

func (cs *ChatService) CreateChat(chat chat.ChatCore, role string) error {
	if chat.Message == "" {
		return errors.New("message cannot be sent empty")
	} else {
		if role == "user" { //Pengirim user
			consulData, errconsulData := cs.consultationData.GetCuntationsById(chat.ConsultationID) //get data from consultation
			if errconsulData != nil {
				return errors.New(notFoundRoomChat)
			}
			chat.ReceiverID = consulData.DoctorID                                                            // receiver is doctor
			valConcul, _ := cs.consultationData.VerUser(chat.SenderID, chat.ReceiverID, chat.ConsultationID) //valid is current id and receiver is avail in roomchat id
			if valConcul.ID == 0 {
				return errors.New("[User] UserID and DoctorID not match at Roomchat")
			}
		} else { //Pengirim admin
			consulData, errconsulData := cs.consultationData.GetCuntationsById(chat.ConsultationID) //get data from consultation
			if errconsulData != nil {
				return errors.New(notFoundRoomChat)
			}
			chat.ReceiverID = consulData.UserID //penerima adalah user
			getDoctorByAdmin, _ := cs.doctorData.SelectByAdminId(chat.SenderID)
			valConcul, _ := cs.consultationData.VerAdmin(getDoctorByAdmin.ID, chat.ReceiverID, chat.ConsultationID)
			chat.SenderID = getDoctorByAdmin.ID
			if valConcul.ID == 0 {
				return errors.New("[Admin] UserID and DoctorID not match at Roomchat")
			}
		}
		chat.TimeStamp = time.Now()
		return cs.chatModel.CreateChat(chat)
	}
}

func (cs *ChatService) GetChats(currentID uint, role string, roomchatID uint) ([]chat.ChatCore, error) {
	//ver role
	if role == "admin" { //role is admin
		doctorData, _ := cs.doctorData.SelectByAdminId(currentID)
		roomChatAvail, _ := cs.consultationData.GetCuntationsById(roomchatID)
		if roomChatAvail.ID == 0 {
			return nil, errors.New("[Admin] RoomChat Not Found")
		} else {
			receiverID := roomChatAvail.UserID
			valRoomchat, _ := cs.consultationData.VerAdmin(doctorData.ID, receiverID, roomchatID)
			if valRoomchat.ID == 0 {
				return nil, errors.New("[Admin] UserID and DoctorID not match at current Roomchat")
			} else {
				return cs.chatModel.GetChatsDoctor(roomchatID)
			}
		}
	} else { //role is user
		roomChatAvail, _ := cs.consultationData.GetCuntationsById(roomchatID)
		if roomChatAvail.ID == 0 {
			return nil, errors.New("[User] RoomChat Not Found")
		} else {
			receiverID := roomChatAvail.DoctorID
			valRoomchat, _ := cs.consultationData.VerUser(currentID, receiverID, roomchatID)
			if valRoomchat.ID == 0 {
				return nil, errors.New("[User] UserID and DoctorID not match at current Roomchat")
			} else {
				return cs.chatModel.GetChatsUser(currentID, roomchatID)
			}
		}
	}
}

func (cs *ChatService) Delete(roomChatID, bubbleChatID, senderID uint, role string) error {
	if roomChatID <= 0 {
		return errors.New("ID must be a positive integer")
	}
	if role == "admin" {
		adminData, _ := cs.doctorData.SelectByAdminId(senderID)
		senderID = adminData.ID
	}
	valConcul, _ := cs.chatModel.VerAvailChat(roomChatID, bubbleChatID, senderID)
	if valConcul == nil { // CurrentID and RoomChat not match (Roomchat not found)
		return errors.New(notFoundRoomChat)
	} else { // CurrentID and RoomChat is match (Roomchat has found)
		return cs.chatModel.Delete(roomChatID, bubbleChatID, senderID)
	}
}
