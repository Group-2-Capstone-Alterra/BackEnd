package service

import (
	"PetPalApp/features/chat"
	"PetPalApp/features/consultation"
	"errors"
	"fmt"
	"log"
	"time"
)

type ChatService struct {
	chatModel        chat.DataInterface
	consultationData consultation.ConsultationModel
}

func New(cm chat.DataInterface, consultationData consultation.ConsultationModel) chat.ServiceInterface {
	return &ChatService{
		chatModel:        cm,
		consultationData: consultationData,
	}
}

func (cs *ChatService) CreateChat(chat chat.ChatCore) error {
	log.Println("[Query - CreateChat]")
	//Verif admin or not
	isAdmin, _ := cs.consultationData.VerIsAdmin(chat.SenderID, chat.ConsultationID)
	if isAdmin.ID == 0 { //Pengirim user
		doctorID, errdoctorID := cs.consultationData.GetCuntationsById(chat.ConsultationID)
		if errdoctorID != nil {
			return fmt.Errorf("roomchat not found")
		}
		chat.ReceiverID = doctorID.DoctorID
		valConcul, _ := cs.consultationData.VerUser(chat.SenderID, chat.ReceiverID, chat.ConsultationID)
		if valConcul.ID == 0 {
			return fmt.Errorf("UserID and DoctorID not match at Roomchat")
		}
	} else { //Pengirim admin
		chat.ReceiverID = isAdmin.UserID //penerima adalah user
		valConcul, _ := cs.consultationData.VerAdmin(chat.SenderID, chat.ReceiverID, chat.ConsultationID)
		if valConcul.ID == 0 {
			return fmt.Errorf("UserID and DoctorID not match at Roomchat 2")
		}
	}
	chat.TimeStamp = time.Now()
	return cs.chatModel.CreateChat(chat)
}

func (cs *ChatService) GetChats(currentID, roomchatID uint) ([]chat.ChatCore, error) {

	log.Println("[Query - GetChats]")
	//Verif admin or not
	valConcul, errVal := cs.consultationData.VerAvailConcul(currentID, roomchatID)
	if valConcul.ID == 0 { // CurrentID and RoomChat not match (Roomchat not found)
		return nil, errVal
	} else { // CurrentID and RoomChat is match (Roomchat has found)
		return cs.chatModel.GetChats(roomchatID)
	}
}

func (cs *ChatService) Delete(roomChatID, bubbleChatID, senderID uint) error {
	if roomChatID <= 0 {
		return errors.New("id not valid")
	}
	log.Println("[Query]")
	//Verif admin or not
	valConcul, _ := cs.chatModel.VerAvailChat(roomChatID, bubbleChatID, senderID)
	if valConcul == nil { // CurrentID and RoomChat not match (Roomchat not found)
		log.Println("[Query - Delete] valConcul NotFound")
		return errors.New("[Query - Delete] valConcul NotFound")
	} else { // CurrentID and RoomChat is match (Roomchat has found)
		log.Println("[Query - Delete] valConcul is found")
		return cs.chatModel.Delete(roomChatID, bubbleChatID, senderID)
	}
}
