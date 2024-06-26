package handler

import (
	"PetPalApp/app/middlewares"
	"PetPalApp/features/admin"
	"PetPalApp/utils/helper"
	"PetPalApp/utils/responses"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	adminService admin.AdminService
	helper       helper.HelperInterface
}

func New(as admin.AdminService, helper helper.HelperInterface) *AdminHandler {
	return &AdminHandler{
		adminService: as,
		helper:       helper,
	}
}

func (ah *AdminHandler) Register(c echo.Context) error {
	newAdmin := AdminRequest{}
	errBind := c.Bind(&newAdmin)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error binding data: "+errBind.Error(), nil))
	}

	dataAdmin := admin.Core{
		FullName: newAdmin.FullName,
		Email:    newAdmin.Email,
		Password: newAdmin.Password,
	}

	errInsert := ah.adminService.Register(dataAdmin)
	if errInsert != nil {
		if strings.Contains(errInsert.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("registration failed: "+errInsert.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("resgistration failed: "+errInsert.Error(), nil))
	}

	return c.JSON(http.StatusCreated, responses.JSONWebResponse("registration succesfull", nil))
}

func (ah *AdminHandler) Login(c echo.Context) error {
	loginReq := LoginRequest{}
	errBind := c.Bind(&loginReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error bind data: "+errBind.Error(), nil))
	}

	result, token, err := ah.adminService.Login(loginReq.Email, loginReq.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("login failed: "+err.Error(), nil))
	}
	result.Token = token
	LoginResponse := LoginResponse {
		ID:       result.ID,
		Role:     result.Role,
		FullName: result.FullName,
		Email:    result.Email,
		Token:    result.Token,
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse("login successfull", LoginResponse))
}

func (ah *AdminHandler) GetProfile(c echo.Context) error {
	adminID, _, _ := middlewares.ExtractTokenUserId(c)
	if adminID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("unauthorized", nil))
	}

	profile, err := ah.adminService.GetProfile(uint(adminID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("get profile failed: "+err.Error(), nil))
	}

	adminResponse := AdminResponse{
		ID:       		profile.ID,
		FullName: 		profile.FullName,
		Email:    		profile.Email,
		NumberPhone: 	profile.NumberPhone,
		Role:     		profile.Role,
		Address:  		profile.Address,
		ProfilePicture: profile.ProfilePicture,
		Coordinate:     profile.Coordinate,
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse("get profile successfull", adminResponse))
}

func (ah *AdminHandler) Delete(c echo.Context) error {
	adminID, _, _ := middlewares.ExtractTokenUserId(c)
	if adminID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("unauthorized", nil))
	}

	errDelete := ah.adminService.Delete(uint(adminID))
	if errDelete != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("delete admin failed: "+errDelete.Error(), nil))
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse("delete admin successfull", nil))
}

func (ah *AdminHandler) Update(c echo.Context) error {
	adminID, _, _ := middlewares.ExtractTokenUserId(c)
	if adminID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("unauthorized", nil))
	}

	updateReq := AdminRequest{}
	errBind := c.Bind(&updateReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error binding data: "+errBind.Error(), nil))
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
		file = nil
		handler = nil
	} else {
		defer file.Close()
	}

	updateData := admin.Core{
		FullName:       updateReq.FullName,
		Email:          updateReq.Email,
		Password:       updateReq.Password,
		Coordinate:     updateReq.Coordinate,
		NumberPhone:    updateReq.NumberPhone,
		Address:        updateReq.Address,
		ProfilePicture: updateReq.ProfilePicture,
	}

	var filename string
	if handler != nil {
		filename = handler.Filename
	}

	errUpdate := ah.adminService.Update(uint(adminID), updateData, file, filename)
	if errUpdate != nil {
		if strings.Contains(errUpdate.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("update failed: "+errUpdate.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("update failed: "+errUpdate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse("update successful", nil))
}

func (ah *AdminHandler) GetAllClinic(c echo.Context) error {
	log.Println("[HANDLER]")

	page := c.QueryParam("page")
	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		pageInt = 1
	}
	log.Println("page:", pageInt)
	offset := (pageInt - 1) * 10

	sortStr := c.QueryParam("sort")

	idToken, _, _ := middlewares.ExtractTokenUserId(c) // extract id user from jwt token
	log.Println("idtoken:", idToken)

	log.Println("[HANDLER - result]")
	result, errResult := ah.adminService.GetAllClinic(uint(idToken), uint(offset), sortStr)
	if errResult != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error read data", nil))
	}

	// clinicSort := ah.helper.SortClinicsByDistance(uint(idToken), result)
	log.Println("[HANDLER - result]", result)

	return c.JSON(http.StatusOK, responses.JSONWebResponse("success read data", result))
}

func (ah *AdminHandler) GetClinicByID(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("ID must be a positive integer", idConv))
	}

	result, errResult := ah.adminService.GetClinic(uint(idConv))
	if errResult != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error read data", nil))
	}

	// clinicSort := ah.helper.SortClinicsByDistance(uint(idToken), result)
	log.Println("[HANDLER - result]", result)

	return c.JSON(http.StatusOK, responses.JSONWebResponse("success read data", result))
}
