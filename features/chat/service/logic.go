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

func (cs *ChatService) CreateChat(chat chat.ChatCore) error {
	log.Println("[Query - CreateChat]")
	//Verif admin or not
	doctorID, _ := cs.doctorData.SelectByAdminId(chat.SenderID)
	log.Println("[Query Chat - CreateChat] doctorID", doctorID.ID)

	isDoctor, _ := cs.consultationData.VerIsDoctor(doctorID.ID, chat.ConsultationID)
	if isDoctor.ID == 0 { //Pengirim user
		doctorID, errdoctorID := cs.consultationData.GetCuntationsById(chat.ConsultationID)
		if errdoctorID != nil {
			return fmt.Errorf("roomchat not found")
		}
		chat.ReceiverID = doctorID.DoctorID
		valConcul, _ := cs.consultationData.VerUser(chat.SenderID, chat.ReceiverID, chat.ConsultationID)
		if valConcul.ID == 0 {
			return fmt.Errorf("[Sender User] UserID and DoctorID not match at Roomchat")
		}
	} else { //Pengirim admin
		chat.ReceiverID = isDoctor.UserID //penerima adalah user
		chat.SenderID = doctorID.ID
		log.Printf("\n[Sender Admin] UserID %v and DoctorID %v Roomchat %v\n", chat.ReceiverID, chat.SenderID, chat.ConsultationID)
		valConcul, _ := cs.consultationData.VerAdmin(chat.SenderID, chat.ReceiverID, chat.ConsultationID)
		if valConcul.ID == 0 {
			return fmt.Errorf("[Sender Admin] UserID and DoctorID not match at Roomchat")
		}
	}
	chat.TimeStamp = time.Now()
	return cs.chatModel.CreateChat(chat)
}

func (cs *ChatService) GetChats(currentID, roomchatID uint) ([]chat.ChatCore, error) {

	log.Println("[Service - GetChats]")
	//get doctorID
	// doctorDetail, _ := cs.doctorData.SelectByAdminId(currentID)
	// log.Println("[Query - GetChats] doctorDetail.AvailableDay.ID", doctorDetail.ID)

	//verif is roomchat and current user is match
	verCurrent, _ := cs.consultationData.VerAvailConcul(currentID, roomchatID)
	if verCurrent.ID == 0 {
		return nil, fmt.Errorf("[Service GetChats - verCurrent] roomchat not found")
	} else {
		log.Println("[Service - GetChats] verCurrent.ID != 0")
		return cs.chatModel.GetChatsUser(currentID, roomchatID)
	}

	//Verif admin or not
	// isDoctor, _ := cs.consultationData.VerIsDoctor(currentID, roomchatID)
	// log.Println("[QueryChat - GetChats] currentID", currentID)
	// log.Println("[QueryChat - GetChats] roomchatID", roomchatID)
	// if isDoctor.ID == 0 { // if not doctor
	// 	return cs.chatModel.GetChatsUser(currentID, roomchatID)
	// } else { // if is doctor
	// 	if doctorDetail.ID == valConcul.DoctorID {
	// 		log.Println("[Query - GetChats] doctorDetail.ID == valConcul.DoctorID")
	// 		return cs.chatModel.GetChatsDoctor(currentID, roomchatID)
	// 	} else {
	// 		log.Println("[Query - GetChats] doctorDetail.ID != valConcul.DoctorID")

	// 	}
	// }
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
