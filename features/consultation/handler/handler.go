package handler

import (
	"PetPalApp/app/middlewares"
	"PetPalApp/features/consultation"
	"PetPalApp/utils/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ConsultationHandler struct {
    consultationService consultation.ConsultationService
}

func New(cs consultation.ConsultationService) *ConsultationHandler {
    return &ConsultationHandler{
        consultationService: cs,
    }
}

func (ch *ConsultationHandler) CreateConsultation(c echo.Context) error {
    userID := middlewares.ExtractTokenUserId(c)
    if userID == 0 {
        return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("Unauthorized", nil))
    }

    newConsultation := ConsultationRequest{}
    if err := c.Bind(&newConsultation); err != nil {
        return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Error binding data: "+err.Error(), nil))
    }

    consultationData := consultation.ConsultationCore{
        UserID:       uint(userID),
        DoctorID:     newConsultation.DoctorID,
        Consultation: newConsultation.Consultation,
    }

    if err := ch.consultationService.CreateConsultation(consultationData); err != nil {
        return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Error creating consultation: "+err.Error(), nil))
    }

    return c.JSON(http.StatusCreated, responses.JSONWebResponse("Consultation created successfully", nil))
}

func (ch *ConsultationHandler) GetConsultationsByUserID(c echo.Context) error {
    userID := middlewares.ExtractTokenUserId(c)
    if userID == 0 {
        return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("Unauthorized", nil))
    }

    consultations, err := ch.consultationService.GetConsultationsByUserID(uint(userID))
    if err != nil {
        return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Error retrieving consultations: "+err.Error(), nil))
    }

    return c.JSON(http.StatusOK, responses.JSONWebResponse("Consultations retrieved successfully", consultations))
}