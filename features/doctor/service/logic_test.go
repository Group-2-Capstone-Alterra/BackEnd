package service_test

import (
	"PetPalApp/features/availdaydoctor"
	"PetPalApp/features/doctor"
	"PetPalApp/features/doctor/service"
	"PetPalApp/mocks"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddDoctor(t *testing.T) {
	mockDoctorModel := new(mocks.DoctorModel)
	mockHelper := new(mocks.HelperInterface)
	doctorService := service.New(mockDoctorModel, mockHelper)

	input := doctor.Core{
		FullName:       "Dr. John Doe",
		Price:          500,
		AdminID:        1,
		ProfilePicture: "",
	}
	file := strings.NewReader("file content")
	handlerFilename := "doctor_profile.jpg"
	timestamp := time.Now().Unix()
	fileName := fmt.Sprintf("%d_%s", timestamp, handlerFilename)
	expectedURL := "https://example.com/doctor_profile.jpg"

	mockHelper.On("UploadDoctorPicture", file, fileName).Return(expectedURL, nil)

	mockDoctorModel.On("SelectByAdminId", input.AdminID).Return(&doctor.Core{}, nil)
	mockDoctorModel.On("AddDoctor", mock.AnythingOfType("doctor.Core")).Return(nil)

	result, err := doctorService.AddDoctor(input, file, handlerFilename)

	assert.Nil(t, err)
	assert.Equal(t, expectedURL, result)

	mockHelper.AssertExpectations(t)
	mockDoctorModel.AssertExpectations(t)
}

func TestGetDoctorByIdAdmin(t *testing.T) {
	mockDoctorModel := new(mocks.DoctorModel)
	mockHelper := new(mocks.HelperInterface)
	doctorService := service.New(mockDoctorModel, mockHelper)

	adminID := uint(1)
	expectedDoctor := &doctor.Core{
		FullName: "Dr. John Doe",
		Price:    500,
		AdminID:  1,
	}

	mockDoctorModel.On("SelectByAdminId", adminID).Return(expectedDoctor, nil)

	result, err := doctorService.GetDoctorByIdAdmin(adminID)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedDoctor, result)

	mockDoctorModel.AssertExpectations(t)
}

func TestGetAvailDoctorByIdDoctor(t *testing.T) {
	mockDoctorModel := new(mocks.DoctorModel)
	mockHelper := new(mocks.HelperInterface)
	doctorService := service.New(mockDoctorModel, mockHelper)

	doctorID := uint(1)
	expectedAvailDay := &availdaydoctor.Core{
		ID: 1,
	}

	mockDoctorModel.On("SelectAvailDayById", doctorID).Return(expectedAvailDay, nil)

	result, err := doctorService.GetAvailDoctorByIdDoctor(doctorID)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedAvailDay, result)

	mockDoctorModel.AssertExpectations(t)
}

func TestUpdateByIdAdmin(t *testing.T) {
	mockDoctorModel := new(mocks.DoctorModel)
	mockHelper := new(mocks.HelperInterface)
	doctorService := service.New(mockDoctorModel, mockHelper)

	adminID := uint(1)
	input := doctor.Core{
		FullName: "Dr. John Doe Updated",
		Price:    600,
	}
	file := strings.NewReader("updated file content")
	handlerFilename := "updated_file.jpg"
	timestamp := time.Now().Unix()
	fileName := fmt.Sprintf("%d_%s", timestamp, handlerFilename)
	expectedURL := "https://example.com/updated_file.jpg"

	mockHelper.On("UploadDoctorPicture", file, fileName).Return(expectedURL, nil)

	mockDoctorModel.On("PutByIdAdmin", adminID, mock.MatchedBy(func(input doctor.Core) bool {
		return input.ProfilePicture == expectedURL
	})).Return(nil)

	result, err := doctorService.UpdateByIdAdmin(adminID, input, file, handlerFilename)

	assert.Nil(t, err)
	assert.Equal(t, expectedURL, result)

	mockHelper.AssertExpectations(t)
	mockDoctorModel.AssertExpectations(t)
}

func TestDeleteDoctor(t *testing.T) {
	mockDoctorModel := new(mocks.DoctorModel)
	mockHelper := new(mocks.HelperInterface)
	doctorService := service.New(mockDoctorModel, mockHelper)

	adminID := uint(1)

	mockDoctorModel.On("Delete", adminID).Return(nil)

	err := doctorService.Delete(adminID)

	assert.Nil(t, err)

	mockDoctorModel.AssertExpectations(t)
}
