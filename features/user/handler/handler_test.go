package handler

import (
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"PetPalApp/features/user"
	"PetPalApp/mocks"
)

func TestUserHandlerRegister(t *testing.T) {
	// Create a new instance of the UserService mock
	userServiceMock := &mocks.UserService{}
	userServiceMock.On("Create", mock.Anything).Return(nil)

	// Create a new instance of the HashInterface mock
	hashServiceMock := &mocks.HashInterface{}
	hashServiceMock.On("HashPassword", mock.Anything).Return("hashed_password", nil)

	// Create a new instance of the UserHandler
	userHandler := New(userServiceMock, hashServiceMock)

	// Create a new echo context
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Create a new user request
	userReq := UserRequest{
		FullName: "John Doe",
		Email:    "johndoe@example.com",
		Password: "hash_password",
	}

	// Bind the user request to the context
	c.Bind(&userReq)

	// Call the Register method
	err := userHandler.Register(c)

	// Assert that the error is nil
	assert.Nil(t, err)

	// Assert that the UserService Create method was called
	userServiceMock.AssertCalled(t, "Create", mock.Anything)

	// Assert that the HashInterface HashPassword method was called
	hashServiceMock.AssertCalled(t, "HashPassword", mock.Anything)
}

func TestUserHandlerLogin(t *testing.T) {
	// Create a new instance of the UserService mock
	userServiceMock := &mocks.UserService{}
	userServiceMock.On("Login", mock.Anything, mock.Anything).Return(&user.Core{}, "token", nil)

	// Create a new instance of the HashInterface mock
	hashServiceMock := &mocks.HashInterface{}
	hashServiceMock.On("CheckPasswordHash", mock.Anything, mock.Anything).Return(true)

	// Create a new instance of the UserHandler
	userHandler := New(userServiceMock, hashServiceMock)

	// Create a new echo context
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Create a new login request
	loginReq := LoginRequest{
		Email:    "johndoe@example.com",
		Password: "password",
	}

	// Set the request body
	req.Body = ioutil.NopCloser(strings.NewReader(`{"email":"johndoe@example.com","password":"password"}`))

	// Bind the login request to the context
	c.Bind(&loginReq)

	// Call the Login method
	err := userHandler.Login(c)

	// Assert that the error is nil
	assert.Nil(t, err)

	// Assert that the UserService Login method was called
	userServiceMock.AssertCalled(t, "Login", mock.Anything, mock.Anything)

	// Assert that the HashInterface CheckPasswordHash method was called
	hashServiceMock.AssertCalled(t, "CheckPasswordHash", mock.Anything, mock.Anything)
}

func TestUserHandler_Profile(t *testing.T) {
	// Create a new instance of the UserService mock
	userServiceMock := &mocks.UserService{}
	userServiceMock.On("GetProfile", mock.Anything).Return(&user.Core{}, nil)

	// Create a new instance of the UserHandler
	userHandler := New(userServiceMock, nil)

	// Create a new echo context
	e := echo.New()
	c := e.NewContext(nil, nil)

	// Set the user ID in the context
	c.Set("user_id", uint(1))

	// Call the Profile method
	err := userHandler.Profile(c)

	// Assert that the error is nil
	assert.Nil(t, err)

	// Assert that the UserService GetProfile method was called
	userServiceMock.AssertCalled(t, "GetProfile", mock.Anything)
}

func TestUserHandler_UpdateUserById(t *testing.T) {
	// Create a new instance of the UserService mock
	userServiceMock := &mocks.UserService{}
	userServiceMock.On("UpdateById", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	// Create a new instance of the UserHandler
	userHandler := New(userServiceMock, nil)

	// Create a new echo context
	e := echo.New()
	c := e.NewContext(nil, nil)

	// Set the user ID in the context
	c.Set("user_id", uint(1))

	// Create a new user request
	userReq := UserRequest{
		FullName: "John Doe",
		Email:    "johndoe@example.com",
	}

	// Bind the user request to the context
	c.Bind(&userReq)

	// Call the UpdateUserById method
	err := userHandler.UpdateUserById(c)

	// Assert that the erroris nil
	assert.Nil(t, err)

	// Assert that the UserService UpdateById method was called
	userServiceMock.AssertCalled(t, "UpdateById", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}

func TestUserHandler_Delete(t *testing.T) {
	// Create a new instance of the UserService mock
	userServiceMock := &mocks.UserService{}
	userServiceMock.On("Delete", mock.Anything).Return(nil)

	// Create a new instance of the UserHandler
	userHandler := New(userServiceMock, nil)

	// Create a new echo context
	e := echo.New()
	c := e.NewContext(nil, nil)

	// Set the user ID in the context
	c.Set("user_id", uint(1))

	// Call the Delete method
	err := userHandler.Delete(c)

	// Assert that the error is nil
	assert.Nil(t, err)

	// Assert that the UserService Delete method was called
	userServiceMock.AssertCalled(t, "Delete", mock.Anything)
}
