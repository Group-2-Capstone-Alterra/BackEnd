package handler

import (
	"PetPalApp/app/middlewares"
	"PetPalApp/features/user"
	"PetPalApp/utils/encrypts"
	"PetPalApp/utils/responses"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

const (
	bindError = "error bind"
)

type UserHandler struct {
	userService user.UserService
	hashService encrypts.HashInterface
}

func New(us user.UserService, hash encrypts.HashInterface) *UserHandler {
	return &UserHandler{
		userService: us,
		hashService: hash,
	}
}

func (uh *UserHandler) Register(c echo.Context) error {
	newUser := UserRequest{}
	errBind := c.Bind(&newUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(bindError+errBind.Error(), nil))
	}

	inputCore := RequestToCore(newUser)
	errInsert := uh.userService.Create(inputCore)
	if errInsert != nil {
		if strings.Contains(errInsert.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Please fill in all required fields. Full Name, Email, and Password cannot be blank.", nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Email already exists. Please try another email address.", nil))
	}
	return c.JSON(http.StatusCreated, responses.JSONWebResponse("Registration successful! You can now log in to your account.", nil))
}

func (uh *UserHandler) Login(c echo.Context) error {
	var reqLoginData = LoginRequest{}
	errBind := c.Bind(&reqLoginData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(bindError+errBind.Error(), nil))
	}
	result, token, err := uh.userService.Login(reqLoginData.Email, reqLoginData.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(err.Error(), nil))
	}
	result.Token = token
	var resultResponse = ResponseLogin(result)
	return c.JSON(http.StatusOK, responses.JSONWebResponse("Login successful! You are now logged in.", resultResponse))
}

func (uh *UserHandler) Profile(c echo.Context) error {
	idToken, _, _ := middlewares.ExtractTokenUserId(c) // extract id user from jwt token
	userData, err := uh.userService.GetProfile(uint(idToken))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Unable to access profile information. Please contact our support team.", nil))
	}
	userResponse := ResponseProfile(*userData)
	return c.JSON(http.StatusOK, responses.JSONWebResponse("Profile information loaded successfully.", userResponse))
}

func (uh *UserHandler) UpdateUserById(c echo.Context) error {
	idToken, _, _ := middlewares.ExtractTokenUserId(c)
	updatedUser := UserRequest{}
	errBind := c.Bind(&updatedUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(bindError+errBind.Error(), nil))
	}

	var file multipart.File
	var handler *multipart.FileHeader
	var err error

	file, handler, err = c.Request().FormFile("profile_picture")
	if err != nil {
		if err != http.ErrMissingFile {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Unable to upload photo: " + err.Error(),
			})
		}
		// Handle the case where no file was uploaded
		file = nil
		handler = nil
	} else {
		defer file.Close()
	}

	inputCore := RequestToCore(updatedUser)

	var filename string
	if handler != nil {
		filename = handler.Filename
	}

	_, errUpdate := uh.userService.UpdateById(uint(idToken), inputCore, file, filename)
	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Unable to save changes. Please try again or contact our support team.", errUpdate))
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse("Profile information updated successfully.", nil))
}

func (uh *UserHandler) Delete(c echo.Context) error {
	idToken, _, _ := middlewares.ExtractTokenUserId(c)
	err := uh.userService.Delete(uint(idToken))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Unable to delete account. Please contact our support team.", nil))
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse("Your account has been deleted. Thank you for using our service.", nil))
}
