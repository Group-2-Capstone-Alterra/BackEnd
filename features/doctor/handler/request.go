package handler

import (
	"PetPalApp/features/availdaydoctor"
	_availHandler "PetPalApp/features/availdaydoctor/handler"
	"PetPalApp/features/doctor"
	"PetPalApp/features/servicedoctor"
	_serviceHandler "PetPalApp/features/servicedoctor/handler"
)

type DoctorRequest struct {
	AdminID        uint
	FullName       string  `json:"full_name" form:"full_name"`
	ProfilePicture string  `json:"profile_picture" form:"profile_picture"`
	About          string  `json:"about" form:"about"`
	Price          float32 `json:"price" form:"price"`
	AvailableDay   _availHandler.AvailableDayRequest
	ServiceDoctor  _serviceHandler.ServiceRequest
}

type AddDoctorRequest struct {
	DoctorRequest
	_availHandler.AvailableDayRequest
	_serviceHandler.ServiceRequest
}

func handlerToCoreAvailableDay(req _availHandler.AvailableDayRequest) availdaydoctor.Core {
	return availdaydoctor.Core{
		Monday:    req.Monday,
		Tuesday:   req.Tuesday,
		Wednesday: req.Wednesday,
		Thursday:  req.Thursday,
		Friday:    req.Friday,
	}
}

func handlerToCoreService(req _serviceHandler.ServiceRequest) servicedoctor.Core {
	return servicedoctor.Core{
		Vaccinations:        req.Vaccinations,
		Operations:          req.Operations,
		MCU:                 req.MCU,
		OnlineConsultations: req.OnlineConsultations,
	}
}

func AddRequestToCore(req AddDoctorRequest) doctor.Core {
	inputCore := doctor.Core{
		AdminID:        req.AdminID,
		FullName:       req.FullName,
		About:          req.About,
		Price:          req.Price,
		ProfilePicture: req.ProfilePicture,
		AvailableDay:   handlerToCoreAvailableDay(req.AvailableDay),
		ServiceDoctor:  handlerToCoreService(req.ServiceDoctor),
	}
	return inputCore
}
