package handler

import (
	"PetPalApp/app/middlewares"
	"PetPalApp/features/doctor"
	"PetPalApp/utils/responses"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type DoctorHandler struct {
	doctorService doctor.DoctorService
	doctorData    doctor.DoctorModel
	// availDaysDoctData
}

func New(ds doctor.DoctorService, doctorData doctor.DoctorModel) *DoctorHandler {
	return &DoctorHandler{
		doctorService: ds,
		doctorData:    doctorData,
	}
}

func (dh *DoctorHandler) AddDoctor(c echo.Context) error {
	adminID, _, _ := middlewares.ExtractTokenUserId(c)
	if adminID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("Unauthorized", nil))
	}
	log.Println("[Handler Doctor - AddDoctor] adminID ", adminID)

	req := AddDoctorRequest{}
	errBind := c.Bind(&req)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Error binding doctor data: "+errBind.Error(), nil))
	}

	// Manually populate the AvailableDays field
	req.AvailableDays = make(map[string]bool)
	for key, value := range c.Request().Form {
		if strings.HasPrefix(key, "available_days[") {
			day := strings.TrimPrefix(key, "available_days[")
			day = strings.TrimSuffix(day, "]")
			req.AvailableDays[day] = value[0] == "true"
		}
	}
	// Manually populate the AvailableDays field
	req.ServiceDoctors = make(map[string]bool)
	for key, value := range c.Request().Form {
		if strings.HasPrefix(key, "services[") {
			service := strings.TrimPrefix(key, "services[")
			service = strings.TrimSuffix(service, "]")
			req.ServiceDoctors[service] = value[0] == "true"
		}
	}

	file, handler, _ := c.Request().FormFile("profile_picture")

	inputCore := AddRequestToCore(req)
	inputCore.AdminID = uint(adminID)

	_, errAdd := dh.doctorService.AddDoctor(inputCore, file, handler.Filename)
	if errAdd != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Unable to add doctor. Please contact our support team.", nil))
	}

	return c.JSON(http.StatusCreated, responses.JSONWebResponse("Doctor added successfully. Thank you.", nil))
}

func (dh *DoctorHandler) ProfileDoctor(c echo.Context) error {
	log.Println("[handler]")

	idToken, _, _ := middlewares.ExtractTokenUserId(c) // extract id user from jwt token
	log.Println("idtoken:", idToken)

	doctorDetails, _ := dh.doctorService.GetDoctorByIdAdmin(uint(idToken))
	if doctorDetails.ID == 0 {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("You don't have a doctor associated with your account. Please add a doctor to continue.", nil))
	}

	log.Println("doctorDetails.ID", doctorDetails.ID)
	availDaysDoct, errAvailDaysDoct := dh.doctorService.GetAvailDoctorByIdDoctor(doctorDetails.ID)
	if errAvailDaysDoct != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Error read data", nil))
	}

	serviceDoctor, errServiceDoctor := dh.doctorData.SelectServiceById(doctorDetails.ID)
	if errServiceDoctor != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Error read data", nil))
	}

	doctorResponse := GormToCore(*doctorDetails, *availDaysDoct, *serviceDoctor)

	return c.JSON(http.StatusOK, responses.JSONWebResponse("success get detail doctor profile", doctorResponse))
}

func (dh *DoctorHandler) UpdateProfile(c echo.Context) error {
	adminID, role, _ := middlewares.ExtractTokenUserId(c)
	if role == "user" {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("Unauthorized", nil))
	}
	log.Println("[Handler Doctor - AddDoctor] adminID ", adminID)

	req := AddDoctorRequest{}
	errBind := c.Bind(&req)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Error binding doctor data: "+errBind.Error(), nil))
	}

	// Manually populate the AvailableDays field
	log.Println("AvailableDays 1", req.AvailableDay)
	req.AvailableDays = make(map[string]bool)
	for key, value := range c.Request().Form {
		if strings.HasPrefix(key, "available_days[") {
			day := strings.TrimPrefix(key, "available_days[")
			day = strings.TrimSuffix(day, "]")
			req.AvailableDays[day] = value[0] == "true"
		}
	}
	log.Println("AvailableDays 2", req.AvailableDay)

	// Manually populate the AvailableDays field
	req.ServiceDoctors = make(map[string]bool)
	for key, value := range c.Request().Form {
		if strings.HasPrefix(key, "services[") {
			service := strings.TrimPrefix(key, "services[")
			service = strings.TrimSuffix(service, "]")
			req.ServiceDoctors[service] = value[0] == "true"
		}
	}

	file, handler, _ := c.Request().FormFile("profile_picture")

	log.Println("[Handler Doctor - AddDoctor] req ", req)
	inputCore := AddRequestToCore(req)

	// inputCore.AdminID = uint(adminID)
	log.Println("[Handler Doctor - AddDoctor] inputCore ", inputCore)

	_, errUpdate := dh.doctorService.UpdateByIdAdmin(uint(adminID), inputCore, file, handler.Filename)
	if errUpdate != nil {
		// Handle error from userService.UpdateById
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Error updating doctor's information. Please contact our support team.", nil))
	}
	// Return success response
	return c.JSON(http.StatusOK, responses.JSONWebResponse("Update successful. Doctor's data has been updated.", nil))
}

func (dh *DoctorHandler) UploadDoctorPicture(c echo.Context) error {
	adminID, role, _ := middlewares.ExtractTokenUserId(c)
	if role == "user" {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("Unauthorized", nil))
	}
	log.Println("[Handler Doctor - AddDoctor] adminID ", adminID)

	req := AddDoctorRequest{}
	errBind := c.Bind(&req)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Error binding doctor data: "+errBind.Error(), nil))
	}

	file, handler, _ := c.Request().FormFile("profile_picture")
	defer file.Close()

	log.Println("[Handler Doctor - AddDoctor] req ", req)
	inputCore := AddRequestToCore(req)

	// inputCore.AdminID = uint(adminID)
	log.Println("[Handler Doctor - AddDoctor] inputCore ", inputCore)

	_, errUpdate := dh.doctorService.UpdateByIdAdmin(uint(adminID), inputCore, file, handler.Filename)
	if errUpdate != nil {
		// Handle error from userService.UpdateById
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error update data", errUpdate))
	}
	// Return success response
	return c.JSON(http.StatusOK, responses.JSONWebResponse("success update data", nil))
}

func (dh *DoctorHandler) Delete(c echo.Context) error {
	adminID, role, _ := middlewares.ExtractTokenUserId(c)
	if role == "user" {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("Unauthorized", nil))
	}

	err := dh.doctorService.Delete(uint(adminID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Unable to delete doctor. Please contact our support team.", nil))
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse("Dcotor has been deleted. Thank you for using our service.", nil))
}
