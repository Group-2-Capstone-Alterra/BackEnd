package handler

import (
	"PetPalApp/app/middlewares"
	"PetPalApp/features/consultation"
	"PetPalApp/features/doctor"
	"PetPalApp/features/user"
	"PetPalApp/utils/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ConsultationHandler struct {
	consultationService consultation.ConsultationService
	userData            user.UserModel
	doctorData          doctor.DoctorModel
}

func New(cs consultation.ConsultationService, userData user.UserModel, doctorData doctor.DoctorModel) *ConsultationHandler {
	return &ConsultationHandler{
		consultationService: cs,
		userData:            userData,
		doctorData:          doctorData,
	}
}

const (
	errRetriev = "Error retrieving consultations: "
	retrieved  = "Consultations retrieved successfully"
	errID      = "ID must be a positive integer"
)

func (ch *ConsultationHandler) CreateConsultation(c echo.Context) error {
	userID, _, _ := middlewares.ExtractTokenUserId(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("Unauthorized", nil))
	}

	doctorID := c.Param("id")
	doctorIDConv, errConv := strconv.Atoi(doctorID)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(errID, errConv))
	}

	newConsultation := ConsultationRequest{}
	if err := c.Bind(&newConsultation); err != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Error binding data: "+err.Error(), nil))
	}
	consultationData := ReqToCore(uint(userID), uint(doctorIDConv), newConsultation)

	if err := ch.consultationService.CreateConsultation(consultationData); err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Error creating consultation: "+err.Error(), nil))
	}
	return c.JSON(http.StatusCreated, responses.JSONWebResponse("Consultation created successfully", nil))
}

func (ch *ConsultationHandler) GetConsultations(c echo.Context) error {
	currentID, role, _ := middlewares.ExtractTokenUserId(c)
	if currentID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("Unauthorized", nil))
	}

	consultations, err := ch.consultationService.GetConsultations(uint(currentID), role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(errRetriev+err.Error(), nil))
	}

	var allConsultation []ConsultationResponse
	for _, v := range consultations {
		userData, _ := ch.userData.SelectById(v.UserID)
		doctorData, _ := ch.doctorData.SelectDoctorById(v.DoctorID)
		allConsultation = append(allConsultation, GormToCore(v, *userData, *doctorData))
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse(retrieved, allConsultation))
}

func (ch *ConsultationHandler) GetConsultationsByUserID(c echo.Context) error {
	doctorIDStr := c.Param("user_id")
	doctorID, err := strconv.ParseUint(doctorIDStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(errID, nil))
	}

	consultations, err := ch.consultationService.GetConsultationsByUserID(uint(doctorID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(errRetriev+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse(retrieved, consultations))
}

func (ch *ConsultationHandler) GetConsultationsByDoctorID(c echo.Context) error {
	doctorIDStr := c.Param("doctor_id")
	doctorID, err := strconv.ParseUint(doctorIDStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(errID, nil))
	}

	consultations, err := ch.consultationService.GetConsultationsByDoctorID(uint(doctorID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(errRetriev+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse(retrieved, consultations))
}

func (ch *ConsultationHandler) UpdateConsultationResponse(c echo.Context) error {
	consultationIDStr := c.Param("consultation_id")
	consultationID, err := strconv.ParseUint(consultationIDStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(errID, nil))
	}

	currentID, role, _ := middlewares.ExtractTokenUserId(c)
	if currentID == 0 && role != "admin" {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("Unauthorized", nil))
	}

	updateRequest := UpdateConsultationRequest{}
	if err := c.Bind(&updateRequest); err != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Error binding data: "+err.Error(), nil))
	}

	inputCore := UpdateReqToCore(updateRequest)
	if err := ch.consultationService.UpdateConsultation(uint(consultationID), inputCore); err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Error updating consultation response: "+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse(retrieved, nil))
}
