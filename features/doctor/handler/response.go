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
	Specialization string                       `json:"specialization"`
	ProfilePicture string                       `json:"profile_picture"`
	AvailableDay   handler.AvailableDayResponse `json:"available_days"`
}

type DetailDoctorResponse struct {
	DoctorResponse
	handler.AvailableDayResponse
}

func GormToCore(gormDoctor doctor.Core, gormAvaildays availdaydoctor.Core) DoctorResponse {
	inputCore := DoctorResponse{
		ID:             gormDoctor.ID,
		FullName:       gormDoctor.FullName,
		Specialization: gormDoctor.Specialization,
		ProfilePicture: gormDoctor.ProfilePicture,
		AvailableDay:   handler.GormToCore(gormAvaildays),
	}
	log.Println("[Handler Doctor - Request] RequestToCore ", inputCore)
	return inputCore
}
