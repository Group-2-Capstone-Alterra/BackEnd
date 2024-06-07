package service

import (
	"PetPalApp/features/admin"
	"PetPalApp/features/chat"
	"PetPalApp/features/consultation"
	"PetPalApp/features/doctor"
	"PetPalApp/features/user"
	"errors"
	"fmt"
	"log"
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

func (cs *ChatService) CreateChat(chat chat.ChatCore, role string) error {
	if role == "user" { //Pengirim user
		consulData, errconsulData := cs.consultationData.GetCuntationsById(chat.ConsultationID) //get data from consultation
		if errconsulData != nil {
			return fmt.Errorf("roomchat not found")
		}
		chat.ReceiverID = consulData.DoctorID                                                            // receiver is doctor
		valConcul, _ := cs.consultationData.VerUser(chat.SenderID, chat.ReceiverID, chat.ConsultationID) //valid is current id and receiver is avail in roomchat id
		if valConcul.ID == 0 {
			return fmt.Errorf("[Sender User] UserID and DoctorID not match at Roomchat")
		}
	} else { //Pengirim admin
		consulData, errconsulData := cs.consultationData.GetCuntationsById(chat.ConsultationID) //get data from consultation
		if errconsulData != nil {
			return fmt.Errorf("roomchat not found")
		}
		chat.ReceiverID = consulData.UserID //penerima adalah user
		getDoctorByAdmin, _ := cs.doctorData.SelectByAdminId(chat.SenderID)
		valConcul, _ := cs.consultationData.VerAdmin(getDoctorByAdmin.ID, chat.ReceiverID, chat.ConsultationID)
		if valConcul.ID == 0 {
			return fmt.Errorf("[Sender Admin] UserID and DoctorID not match at Roomchat")
		}
	}
	chat.TimeStamp = time.Now()
	return cs.chatModel.CreateChat(chat)
}

func (cs *ChatService) GetChats(currentID uint, role string, roomchatID uint) ([]chat.ChatCore, error) {
	//ver role
	if role == "admin" { //role is admin
		roomChatAvail, _ := cs.consultationData.GetCuntationsById(roomchatID)
		if roomChatAvail.ID == 0 {
			return nil, fmt.Errorf("[Role Admin] RoomChat Not Found")
		} else {
			receiverID := roomChatAvail.UserID
			getDoctorByAdmin, _ := cs.doctorData.SelectByAdminId(currentID)
			valRoomchat, _ := cs.consultationData.VerAdmin(getDoctorByAdmin.ID, receiverID, roomchatID)
			if valRoomchat.ID == 0 {
				return nil, fmt.Errorf("[Sender Admin] UserID and DoctorID not match at current Roomchat")
			} else {
				return cs.chatModel.GetChatsUser(getDoctorByAdmin.ID, roomchatID)
			}
		}
	} else { //role is user
		roomChatAvail, _ := cs.consultationData.GetCuntationsById(roomchatID)
		if roomChatAvail.ID == 0 {
			return nil, fmt.Errorf("[Role User] RoomChat Not Found")
		} else {
			receiverID := roomChatAvail.DoctorID
			valRoomchat, _ := cs.consultationData.VerUser(currentID, receiverID, roomchatID)
			if valRoomchat.ID == 0 {
				return nil, fmt.Errorf("[Sender User] UserID and DoctorID not match at current Roomchat")
			} else {
				log.Println("[Service - GetChats] Data has been found")
				return cs.chatModel.GetChatsUser(currentID, roomchatID)
			}
		}
	}
}

func (cs *ChatService) Delete(roomChatID, bubbleChatID, senderID uint) error {
	if roomChatID <= 0 {
		return errors.New("ID must be a positive integer")
	}
	valConcul, _ := cs.chatModel.VerAvailChat(roomChatID, bubbleChatID, senderID)
	if valConcul == nil { // CurrentID and RoomChat not match (Roomchat not found)
		return errors.New("Chat not found")
	} else { // CurrentID and RoomChat is match (Roomchat has found)
		return cs.chatModel.Delete(roomChatID, bubbleChatID, senderID)
	}
}
