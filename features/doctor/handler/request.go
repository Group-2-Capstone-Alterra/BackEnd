package handler

import (
	"PetPalApp/features/availdaydoctor"
	"PetPalApp/features/availdaydoctor/handler"
	_availHandler "PetPalApp/features/availdaydoctor/handler"
	"PetPalApp/features/doctor"
	"PetPalApp/features/servicedoctor"
	_serviceHandler "PetPalApp/features/servicedoctor/handler"
	"net/url"
	"strings"
)

type DoctorRequest struct {
	AdminID        uint
	FullName       string                            `json:"full_name" form:"full_name"`
	ProfilePicture string                            `json:"profile_picture" form:"profile_picture"`
	About          string                            `json:"about" form:"about"`
	Price          float32                           `json:"price" form:"price"`
	AvailableDay   _availHandler.AvailableDayRequest `json:"available_days" form:"available_days" query:"available_days"`
	ServiceDoctor  _serviceHandler.ServiceRequest    `json:"services" form:"services" query:"services"`
}

type AddDoctorRequest struct {
	DoctorRequest
	AvailableDays  map[string]bool `form:"available_days"`
	ServiceDoctors map[string]bool `form:"services"`
}

func (d *DoctorRequest) UnmarshalForm(form url.Values) error {
	// ...
	availableDays := make(map[string]bool)
	for key, value := range form {
		if strings.HasPrefix(key, "available_days[") {
			day := strings.TrimPrefix(key, "available_days[")
			day = strings.TrimSuffix(day, "]")
			availableDays[day] = value[0] == "true"
		}
	}

	d.AvailableDay = handler.AvailableDayRequest{
		Monday:    availableDays["monday"],
		Tuesday:   availableDays["tuesday"],
		Wednesday: availableDays["wednesday"],
		Thursday:  availableDays["thursday"],
		Friday:    availableDays["friday"],
	}

	servicesDoctors := make(map[string]bool)
	for key, value := range form {
		if strings.HasPrefix(key, "services[") {
			service := strings.TrimPrefix(key, "services[")
			service = strings.TrimSuffix(service, "]")
			servicesDoctors[service] = value[0] == "true"
		}
	}

	d.ServiceDoctor = _serviceHandler.ServiceRequest{
		Vaccinations: servicesDoctors["vaccinations"],
		Operations:   servicesDoctors["operations"],
		MCU:          servicesDoctors["mcu"],
	}
	return nil
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
		Vaccinations: req.Vaccinations,
		Operations:   req.Operations,
		MCU:          req.MCU,
	}
}

func AddRequestToCore(req AddDoctorRequest) doctor.Core {
	inputCore := doctor.Core{
		AdminID:        req.AdminID,
		FullName:       req.FullName,
		About:          req.About,
		Price:          req.Price,
		ProfilePicture: req.ProfilePicture,
		AvailableDay: handlerToCoreAvailableDay(_availHandler.AvailableDayRequest{
			Monday:    req.AvailableDays["monday"],
			Tuesday:   req.AvailableDays["tuesday"],
			Wednesday: req.AvailableDays["wednesday"],
			Thursday:  req.AvailableDays["thursday"],
			Friday:    req.AvailableDays["friday"],
		}),
		ServiceDoctor: handlerToCoreService(_serviceHandler.ServiceRequest{
			Vaccinations: req.ServiceDoctors["vaccinations"],
			Operations:   req.ServiceDoctors["operations"],
			MCU:          req.ServiceDoctors["mcu"],
		}),
	}
	return inputCore
}
