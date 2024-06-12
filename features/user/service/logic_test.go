package service

import (
	"PetPalApp/features/user"
	"PetPalApp/mocks"
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setup() (*userService, *mocks.UserModel, *mocks.HashInterface, *mocks.HelperInterface) {
	mockUserData := &mocks.UserModel{}
	mockHashService := &mocks.HashInterface{}
	mockHelperService := &mocks.HelperInterface{}
	userService := New(mockUserData, mockHashService, mockHelperService).(*userService) // Type assertion here
	return userService, mockUserData, mockHashService, mockHelperService
}

func TestUserServiceCreate(t *testing.T) {
	userService, mockUserData, mockHashService, _ := setup()

	input := user.Core{
		FullName: "John Doe",
		Email:    "johndoe@example.com",
		Password: "plain_password",
	}

	t.Run("success", func(t *testing.T) {
		mockUserData.On("Insert", mock.Anything).Return(nil)
		mockHashService.On("HashPassword", input.Password).Return("hashed_password", nil)

		err := userService.Create(input)
		assert.Nil(t, err)
		mockUserData.AssertExpectations(t)
		mockHashService.AssertExpectations(t)
	})

	t.Run("error hashing password", func(t *testing.T) {
		mockHashService.On("HashPassword", input.Password).Return("", errors.New("hash error"))

		err := userService.Create(input)
		assert.NotNil(t, err)
		mockHashService.AssertExpectations(t)
	})

	t.Run("error inserting user", func(t *testing.T) {
		mockHashService.On("HashPassword", input.Password).Return("hashed_password", nil)
		mockUserData.On("Insert", mock.Anything).Return(errors.New("insert error"))

		err := userService.Create(input)
		assert.NotNil(t, err)
		mockUserData.AssertExpectations(t)
	})
}

func TestUserServiceLogin(t *testing.T) {
	userService, mockUserData, mockHashService, _ := setup()

	email := "johndoe@example.com"
	password := "plain_password"
	hashedPassword := "hashed_password"

	t.Run("success", func(t *testing.T) {
		mockUserData.On("SelectByEmail", email).Return(&user.Core{
			ID:       1,
			Email:    email,
			Password: hashedPassword,
		}, nil)
		mockHashService.On("CheckPasswordHash", hashedPassword, password).Return(true)

		data, token, err := userService.Login(email, password)
		assert.NotNil(t, data)
		assert.NotEmpty(t, token)
		assert.Nil(t, err)
		mockUserData.AssertExpectations(t)
		mockHashService.AssertExpectations(t)
	})

	t.Run("user not found", func(t *testing.T) {
		mockUserData.On("SelectByEmail", email).Return(nil, errors.New("user not found"))

		data, token, err := userService.Login(email, password)
		assert.Nil(t, data)
		assert.Empty(t, token)
		assert.NotNil(t, err)
		mockUserData.AssertExpectations(t)
	})

	t.Run("invalid password", func(t *testing.T) {
		mockUserData.On("SelectByEmail", email).Return(&user.Core{
			ID:       1,
			Email:    email,
			Password: hashedPassword,
		}, nil)
		mockHashService.On("CheckPasswordHash", hashedPassword, password).Return(false)

		data, token, err := userService.Login(email, password)
		assert.Nil(t, data)
		assert.Empty(t, token)
		assert.NotNil(t, err)
		mockUserData.AssertExpectations(t)
		mockHashService.AssertExpectations(t)
	})

	// t.Run("token creation failure", func(t *testing.T) {
	// 	mockUserData.On("SelectByEmail", email).Return(&user.Core{
	// 		ID:       1,
	// 		Email:    email,
	// 		Password: hashedPassword,
	// 	}, nil)
	// 	mockHashService.On("CheckPasswordHash", hashedPassword, password).Return(true)
	// 	middlewares.CreateToken = func(userID int, role string) (string, error) {
	// 		return "", errors.New("token error")
	// 	}

	// 	data, token, err := userService.Login(email, password)
	// 	assert.Nil(t, data)
	// 	assert.Empty(t, token)
	// 	assert.NotNil(t, err)
	// })
}

func TestUserServiceGetProfile(t *testing.T) {
	userService, mockUserData, _, _ := setup()

	id := uint(1)

	t.Run("success", func(t *testing.T) {
		mockUserData.On("SelectById", id).Return(&user.Core{
			ID:       id,
			FullName: "John Doe",
			Email:    "johndoe@example.com",
		}, nil)

		data, err := userService.GetProfile(id)
		assert.NotNil(t, data)
		assert.Nil(t, err)
		mockUserData.AssertExpectations(t)
	})

	t.Run("invalid ID", func(t *testing.T) {
		data, err := userService.GetProfile(0)
		assert.Nil(t, data)
		assert.NotNil(t, err)
	})

	t.Run("user not found", func(t *testing.T) {
		mockUserData.On("SelectById", id).Return(nil, errors.New("user not found"))

		data, err := userService.GetProfile(id)
		assert.Nil(t, data)
		assert.NotNil(t, err)
		mockUserData.AssertExpectations(t)
	})
}

func TestUserServiceUpdateById(t *testing.T) {
	userService, mockUserData, mockHashService, mockHelperService := setup()

	id := uint(1)
	input := user.Core{
		FullName:       "Jane Doe",
		Email:          "janedoe@example.com",
		Password:       "plain_password",
		ProfilePicture: "profile_picture.jpg",
	}
	file := &bytes.Buffer{}
	handlerFilename := "profile_picture.jpg"

	t.Run("success with file", func(t *testing.T) {
		mockHashService.On("HashPassword", input.Password).Return("hashed_password", nil)
		mockHelperService.On("UploadProfilePicture", file, mock.Anything).Return("profile_picture.jpg", nil)
		mockUserData.On("PutById", id, mock.Anything).Return(nil)

		photoFileName, err := userService.UpdateById(id, input, file, handlerFilename)
		assert.Equal(t, "profile_picture.jpg", photoFileName)
		assert.Nil(t, err)
		mockUserData.AssertExpectations(t)
		mockHashService.AssertExpectations(t)
		mockHelperService.AssertExpectations(t)
	})

	t.Run("success without file", func(t *testing.T) {
		input.ProfilePicture = ""
		file = nil
		handlerFilename = ""

		mockHashService.On("HashPassword", input.Password).Return("hashed_password", nil)
		mockUserData.On("PutById", id, mock.Anything).Return(nil)

		photoFileName, err := userService.UpdateById(id, input, file, handlerFilename)
		assert.Empty(t, photoFileName)
		assert.Nil(t, err)
		mockUserData.AssertExpectations(t)
		mockHashService.AssertExpectations(t)
	})

	t.Run("invalid ID", func(t *testing.T) {
		photoFileName, err := userService.UpdateById(0, input, file, handlerFilename)
		assert.Empty(t, photoFileName)
		assert.NotNil(t, err)
	})

	t.Run("hash password error", func(t *testing.T) {
		mockHashService.On("HashPassword", input.Password).Return("", errors.New("hash error"))

		photoFileName, err := userService.UpdateById(id, input, file, handlerFilename)
		assert.Empty(t, photoFileName)
		assert.NotNil(t, err)
		mockHashService.AssertExpectations(t)
	})

	t.Run("upload profile picture error", func(t *testing.T) {
		mockHashService.On("HashPassword", input.Password).Return("hashed_password", nil)
		mockHelperService.On("UploadProfilePicture", file, mock.Anything).Return("", errors.New("upload error"))
		photoFileName, err := userService.UpdateById(id, input, file, handlerFilename)
		assert.Empty(t, photoFileName)
		assert.NotNil(t, err)
		mockHashService.AssertExpectations(t)
		mockHelperService.AssertExpectations(t)
	})

	t.Run("put by id error", func(t *testing.T) {
		mockHashService.On("HashPassword", input.Password).Return("hashed_password", nil)
		mockUserData.On("PutById", id, mock.Anything).Return(errors.New("put by id error"))

		photoFileName, err := userService.UpdateById(id, input, file, handlerFilename)
		assert.Empty(t, photoFileName)
		assert.NotNil(t, err)
		mockUserData.AssertExpectations(t)
	})
}

func TestUserServiceDelete(t *testing.T) {
	userService, mockUserData, _, _ := setup()

	id := uint(1)

	t.Run("success", func(t *testing.T) {
		mockUserData.On("Delete", id).Return(nil)

		err := userService.Delete(id)
		assert.Nil(t, err)
		mockUserData.AssertExpectations(t)
	})

	t.Run("invalid ID", func(t *testing.T) {
		err := userService.Delete(0)
		assert.NotNil(t, err)
	})

	t.Run("delete error", func(t *testing.T) {
		mockUserData.On("Delete", id).Return(errors.New("delete error"))

		err := userService.Delete(id)
		assert.NotNil(t, err)
		mockUserData.AssertExpectations(t)
	})
}
