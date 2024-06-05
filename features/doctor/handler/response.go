package handler

import (
	"PetPalApp/features/availdaydoctor"
	"PetPalApp/features/availdaydoctor/handler"
	"PetPalApp/features/doctor"
	"log"
)

type DoctorResponse struct {
	ID             uint                         `json:"id"`
	FullName       string                       `json:"full_name"`
	Price          float32                      `json:"price"`
	About          string                       `json:"about"`
	ProfilePicture string                       `json:"profile_picture"`
	AvailableDay   handler.AvailableDayResponse `json:"available_days"`
}

type ConsulDoctorReponse struct {
	ID             uint   `json:"id,omitempty"`
	FullName       string `json:"full_name,omitempty"`
	ProfilePicture string `json:"profile_picture,omitempty"`
}

func ConsulGormToCore(gormDoctor doctor.Core) ConsulDoctorReponse {
	inputCore := ConsulDoctorReponse{
		ID:             gormDoctor.ID,
		FullName:       gormDoctor.FullName,
		ProfilePicture: gormDoctor.ProfilePicture,
	}
	return inputCore
}

type DetailDoctorResponse struct {
	DoctorResponse
	handler.AvailableDayResponse
}

func GormToCore(gormDoctor doctor.Core, gormAvaildays availdaydoctor.Core) DoctorResponse {
	inputCore := DoctorResponse{
		ID:             gormDoctor.ID,
		FullName:       gormDoctor.FullName,
		Price:          gormDoctor.Price,
		About:          gormDoctor.About,
		ProfilePicture: gormDoctor.ProfilePicture,
		AvailableDay:   handler.GormToCore(gormAvaildays),
	}
	log.Println("[Handler Doctor - Request] RequestToCore ", inputCore)
	return inputCore
}
