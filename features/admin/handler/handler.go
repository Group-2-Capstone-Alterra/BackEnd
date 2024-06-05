package handler

import (
	"PetPalApp/app/middlewares"
	"PetPalApp/features/admin"
	"PetPalApp/utils/responses"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	adminService admin.AdminService
}

func New(as admin.AdminService) *AdminHandler {
	return &AdminHandler{
		adminService: as,
	}
}

func (ah *AdminHandler) Register(c echo.Context) error {
	newAdmin := AdminRequest{}
	errBind := c.Bind(&newAdmin)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error binding data: "+errBind.Error(), nil))
	}

	dataAdmin := admin.Core{
		FullName:           newAdmin.FullName,
		Email:              newAdmin.Email,
		Password:           newAdmin.Password,
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
	var LoginResponse = ResponseLogin(result)
	return c.JSON(http.StatusOK, responses.JSONWebResponse("login successfull", LoginResponse))
}

func (ah *AdminHandler) GetProfile(c echo.Context) error {
	adminID := middlewares.ExtractTokenUserId(c)
	if adminID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("unauthorized", nil))
	}

	profile, err := ah.adminService.GetProfile(uint(adminID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("get profile failed: "+err.Error(), nil))
	}

	adminResponse := ResponseProfile(*profile)
	return c.JSON(http.StatusOK, responses.JSONWebResponse("get profile successfull", adminResponse))
}

func (ah *AdminHandler) Delete(c echo.Context) error {
	adminID := middlewares.ExtractTokenUserId(c)
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
	adminID := middlewares.ExtractTokenUserId(c)
	if adminID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("unauthorized", nil))
	}

	updateReq := AdminRequest{}
	errBind := c.Bind(&updateReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error binding data: "+errBind.Error(), nil))
	}

	updateData := admin.Core{
		FullName:       updateReq.FullName,
		Email:          updateReq.Email,
		NumberPhone:    updateReq.NumberPhone,
		Address:        updateReq.Address,
		ProfilePicture: updateReq.ProfilePicture,
	}

	errUpdate := ah.adminService.Update(uint(adminID), updateData)
	if errUpdate != nil {
		if strings.Contains(errUpdate.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("update failed: "+errUpdate.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("update failed: "+errUpdate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse("update successfull", nil))
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

	idToken := middlewares.ExtractTokenUserId(c) // extract id user from jwt token
	log.Println("idtoken:", idToken)

	log.Println("[HANDLER - result]")
	result, errResult := ah.adminService.GetAllClinic(uint(idToken), uint(offset), sortStr)
	if errResult != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error read data", nil))
	}

	// var allAdmin []handler.AllClinicResponse
	// for _, v := range result {
	// 	allAdmin = append(allAdmin, handler.ResponseAllClinic(v))
	// }

	return c.JSON(http.StatusOK, responses.JSONWebResponse("success read data", result))
}
