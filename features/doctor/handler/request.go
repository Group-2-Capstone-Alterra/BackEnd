package handler

import (
	"PetPalApp/features/availdaydoctor/handler"
	"PetPalApp/features/doctor"
	"log"
)

type DoctorRequest struct {
	AdminID        uint
	FullName       string                      `json:"full_name" form:"full_name"`
	Email          string                      `json:"email" form:"email"`
	Specialization string                      `json:"specialization" form:"specialization"`
	ProfilePicture string                      `json:"profile_picture" form:"profile_picture"`
	AvailableDay   handler.AvailableDayRequest `json:"available_days" form:"available_days" query:"available_days"`
}

type AddDoctorRequest struct {
	DoctorRequest
	handler.AvailableDayRequest
}

func RequestToCore(input AddDoctorRequest) doctor.Core {
	inputCore := doctor.Core{
		AdminID:        input.AdminID,
		FullName:       input.FullName,
		Specialization: input.Specialization,
		ProfilePicture: input.ProfilePicture,
		AvailableDay:   handler.RequestToCore(input.DoctorID, input.AvailableDay),
	}
	log.Println("[Handler Doctor - Request] RequestToCore ", inputCore)
	return inputCore
}
