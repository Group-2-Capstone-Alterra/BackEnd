package service

import (
	"PetPalApp/features/user"
	"PetPalApp/mocks"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserServiceCreate(t *testing.T) {
	mockUserData := &mocks.UserModel{}
	mockHashService := &mocks.HashInterface{}
	mockHelperService := &mocks.HelperInterface{}
	userService := New(mockUserData, mockHashService, mockHelperService)

	input := user.Core{
		FullName: "John Doe",
		Email:    "johndoe@example.com",
		Password: "hashed_password",
	}

	mockUserData.On("Insert", input).Return(nil)
	mockHashService.On("HashPassword", input.Password).Return("hashed_password", nil)

	err := userService.Create(input)
	assert.Nil(t, err)

	mockUserData.AssertExpectations(t)
	mockHashService.AssertExpectations(t)
}

func TestUserServiceGetProfile(t *testing.T) {
	mockUserData := &mocks.UserModel{}
	mockHashService := &mocks.HashInterface{}
	mockHelperService := &mocks.HelperInterface{}
	// mockJwt := &mocks.JwtInterface{}
	userService := New(mockUserData, mockHashService, mockHelperService)

	id := uint(1)

	mockUserData.On("SelectById", id).Return(&user.Core{
		ID:       1,
		FullName: "John Doe",
		Email:    "johndoe@example.com",
	}, nil)

	data, err := userService.GetProfile(id)
	assert.NotNil(t, data)
	assert.Nil(t, err)

	mockUserData.AssertExpectations(t)
}

func TestUserServiceUpdateById(t *testing.T) {
	mockUserData := &mocks.UserModel{}
	mockHashService := &mocks.HashInterface{}
	mockHelperService := &mocks.HelperInterface{}
	// mockJwt := &mocks.JwtInterface{}
	userService := New(mockUserData, mockHashService, mockHelperService)

	id := uint(1)
	input := user.Core{
		FullName:       "Jane Doe",
		Email:          "janedoe@example.com",
		Password:       "hashed_password", // Use the hashed password here
		ProfilePicture: "profile_picture.jpg",
	}
	file := &bytes.Buffer{}
	handlerFilename := "profile_picture.jpg"

	mockUserData.On("PutById", id, input).Return(nil)
	mockHashService.On("HashPassword", input.Password).Return("hashed_password", nil)
	mockHelperService.On("UploadProfilePicture", file, mock.Anything).Return("profile_picture.jpg", nil)
	photoFileName, err := userService.UpdateById(id, input, file, handlerFilename)
	assert.Equal(t, "profile_picture.jpg", photoFileName)
	assert.Nil(t, err)

	mockUserData.AssertExpectations(t)
	mockHashService.AssertExpectations(t)
	mockHelperService.AssertExpectations(t)
}

func TestUserServiceDelete(t *testing.T) {
	mockUserData := &mocks.UserModel{}
	mockHashService := &mocks.HashInterface{}
	mockHelperService := &mocks.HelperInterface{}
	// mockJwt := &mocks.JwtInterface{}
	userService := New(mockUserData, mockHashService, mockHelperService)

	id := uint(1)

	mockUserData.On("Delete", id).Return(nil)

	err := userService.Delete(id)
	assert.Nil(t, err)

	mockUserData.AssertExpectations(t)
}
