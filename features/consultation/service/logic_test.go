package service_test

import (
	"PetPalApp/features/consultation"
	"PetPalApp/features/consultation/service"
	"PetPalApp/features/doctor"
	"PetPalApp/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateConsultation(t *testing.T) {
	mockConsultationModel := new(mocks.ConsultationModel)
	mockDoctorModel := new(mocks.DoctorModel)
	mockAdminModel := new(mocks.AdminModel)
	consultationService := service.New(mockConsultationModel, mockDoctorModel, mockAdminModel)

	input := consultation.ConsultationCore{
		DoctorID: 1,
	}

	doctorInfo := &doctor.Core{
		ID: 1,
	}

	mockDoctorModel.On("SelectDoctorById", input.DoctorID).Return(doctorInfo, nil)
	mockConsultationModel.On("CreateConsultation", mock.AnythingOfType("consultation.ConsultationCore")).Return(nil)

	err := consultationService.CreateConsultation(input)

	assert.Nil(t, err)
	mockDoctorModel.AssertExpectations(t)
	mockConsultationModel.AssertExpectations(t)
}

func TestCreateConsultationDoctorNotFound(t *testing.T) {
	mockConsultationModel := new(mocks.ConsultationModel)
	mockDoctorModel := new(mocks.DoctorModel)
	mockAdminModel := new(mocks.AdminModel)
	consultationService := service.New(mockConsultationModel, mockDoctorModel, mockAdminModel)

	input := consultation.ConsultationCore{
		DoctorID: 1,
	}

	doctorInfo := &doctor.Core{
		ID: 0,
	}

	mockDoctorModel.On("SelectDoctorById", input.DoctorID).Return(doctorInfo, nil)

	err := consultationService.CreateConsultation(input)

	assert.NotNil(t, err)
	assert.Equal(t, "Doctor with that ID was not found in any clinic.", err.Error())
	mockDoctorModel.AssertExpectations(t)
	mockConsultationModel.AssertNotCalled(t, "CreateConsultation")
}

func TestGetConsultations_User(t *testing.T) {
	mockConsultationModel := new(mocks.ConsultationModel)
	mockDoctorModel := new(mocks.DoctorModel)
	mockAdminModel := new(mocks.AdminModel)
	consultationService := service.New(mockConsultationModel, mockDoctorModel, mockAdminModel)

	currentID := uint(1)
	role := "user"
	expectedConsultations := []consultation.ConsultationCore{
		{DoctorID: 1, UserID: 1},
	}

	mockConsultationModel.On("GetConsultationsByUserID", currentID).Return(expectedConsultations, nil)

	result, err := consultationService.GetConsultations(currentID, role)

	assert.Nil(t, err)
	assert.Equal(t, expectedConsultations, result)
	mockConsultationModel.AssertExpectations(t)
}

func TestGetConsultations_Admin(t *testing.T) {
	mockConsultationModel := new(mocks.ConsultationModel)
	mockDoctorModel := new(mocks.DoctorModel)
	mockAdminModel := new(mocks.AdminModel)
	consultationService := service.New(mockConsultationModel, mockDoctorModel, mockAdminModel)

	currentID := uint(1)
	role := "admin"
	doctorID := &doctor.Core{ID: 1}
	expectedConsultations := []consultation.ConsultationCore{
		{DoctorID: 1, UserID: 1},
	}

	mockDoctorModel.On("SelectByAdminId", currentID).Return(doctorID, nil)
	mockConsultationModel.On("GetConsultationsByDoctorID", doctorID.ID).Return(expectedConsultations, nil)

	result, err := consultationService.GetConsultations(currentID, role)

	assert.Nil(t, err)
	assert.Equal(t, expectedConsultations, result)
	mockDoctorModel.AssertExpectations(t)
	mockConsultationModel.AssertExpectations(t)
}

func TestGetConsultationsByUserID(t *testing.T) {
	mockConsultationModel := new(mocks.ConsultationModel)
	mockDoctorModel := new(mocks.DoctorModel)
	mockAdminModel := new(mocks.AdminModel)
	consultationService := service.New(mockConsultationModel, mockDoctorModel, mockAdminModel)

	userID := uint(1)
	expectedConsultations := []consultation.ConsultationCore{
		{DoctorID: 1, UserID: 1},
	}

	mockConsultationModel.On("GetConsultationsByUserID", userID).Return(expectedConsultations, nil)

	result, err := consultationService.GetConsultationsByUserID(userID)

	assert.Nil(t, err)
	assert.Equal(t, expectedConsultations, result)
	mockConsultationModel.AssertExpectations(t)
}

func TestGetConsultationsByDoctorID(t *testing.T) {
	mockConsultationModel := new(mocks.ConsultationModel)
	mockDoctorModel := new(mocks.DoctorModel)
	mockAdminModel := new(mocks.AdminModel)
	consultationService := service.New(mockConsultationModel, mockDoctorModel, mockAdminModel)

	doctorID := uint(1)
	expectedConsultations := []consultation.ConsultationCore{
		{DoctorID: 1, UserID: 1},
	}

	mockConsultationModel.On("GetConsultationsByDoctorID", doctorID).Return(expectedConsultations, nil)

	result, err := consultationService.GetConsultationsByDoctorID(doctorID)

	assert.Nil(t, err)
	assert.Equal(t, expectedConsultations, result)
	mockConsultationModel.AssertExpectations(t)
}

func TestUpdateConsultation(t *testing.T) {
	mockConsultationModel := new(mocks.ConsultationModel)
	mockDoctorModel := new(mocks.DoctorModel)
	mockAdminModel := new(mocks.AdminModel)
	consultationService := service.New(mockConsultationModel, mockDoctorModel, mockAdminModel)

	consulID := uint(1)
	input := consultation.ConsultationCore{
		DoctorID: 1,
		UserID:   1,
	}

	mockConsultationModel.On("UpdateConsultation", consulID, input).Return(nil)

	err := consultationService.UpdateConsultation(consulID, input)

	assert.Nil(t, err)
	mockConsultationModel.AssertExpectations(t)
}
