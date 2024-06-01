package service

import (
	"PetPalApp/features/chat"
	"PetPalApp/features/consultation"
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

	log.Println("[Query]")
	log.Println("[query - consul id 1] ID", chat.ConsultationID)

	//Verif admin or not
	isAdmin, _ := cs.consultationData.VerIsAdmin(chat.SenderID, chat.ConsultationID)
	if isAdmin.ID == 0 {
		log.Println("[query - consul id 2] ID", chat.ConsultationID)
		doctorID, _ := cs.consultationData.GetCuntationsById(chat.ConsultationID)
		chat.ReceiverID = doctorID.DoctorID //penerima adalah admin
		log.Println("[query - sender user] ID Sender", chat.SenderID)
		log.Println("[query - receiver admin] ID Receiver", chat.ReceiverID)
	} else {
		chat.ReceiverID = isAdmin.DoctorID //penerima adalah user
		log.Println("[query - sender admin] ID Sender", chat.SenderID)
		log.Println("[query - receiver user] ID Receiver", chat.ReceiverID)
	}
	chat.TimeStamp = time.Now()
	return cs.chatModel.CreateChat(chat)
}

func (cs *ChatService) GetChats(currentID, roomchatID uint) ([]chat.ChatCore, error) {

	log.Println("[Query]")

	//Verif admin or not
	valConcul, errVal := cs.consultationData.VerAvailConcul(currentID, roomchatID)
	if valConcul.ID == 0 { // CurrentID and RoomChat not match (Roomchat not found)
		log.Println("[Query]")
		return nil, errVal
	} else { // CurrentID and RoomChat is match (Roomchat has found)
		return cs.chatModel.GetChats(roomchatID)
	}

}
