package handler

import (
	"PetPalApp/app/middlewares"
	"PetPalApp/features/admin"
	"PetPalApp/utils/responses"
	"net/http"
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
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Error binding data: "+errBind.Error(), nil))
	}

	dataAdmin := admin.Core{
		FullName:           newAdmin.FullName,
		Email:              newAdmin.Email,
		NumberPhone:        newAdmin.NumberPhone,
		Address:            newAdmin.Address,
		Password:           newAdmin.Password,
		KetikUlangPassword: newAdmin.KetikUlangPassword,
	}

	errInsert := ah.adminService.Register(dataAdmin)
	if errInsert != nil {
		if strings.Contains(errInsert.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Registrasi gagal: "+errInsert.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Registrasi gagal: "+errInsert.Error(), nil))
	}

	return c.JSON(http.StatusCreated, responses.JSONWebResponse("Registrasi berhasil", nil))
}

func (ah *AdminHandler) Login(c echo.Context) error {
	loginReq := LoginRequest{}
	errBind := c.Bind(&loginReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error bind data: "+errBind.Error(), nil))
	}

	_, token, err := ah.adminService.Login(loginReq.Email, loginReq.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("login gagal: "+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse("login berhasil", echo.Map{"token": token}))
}

func (ah *AdminHandler) GetProfile(c echo.Context) error {
	adminID := middlewares.ExtractTokenUserId(c)
	if adminID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("Unauthorized", nil))
	}

	profile, err := ah.adminService.GetProfile(uint(adminID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("gagal mengambil profil: "+err.Error(), nil))
	}

	adminResponse := AdminResponse{
		FullName: profile.FullName,
		Email: profile.Email,
		Address: profile.Address,
		NumberPhone: profile.NumberPhone,
		ProfilePicture: profile.ProfilePicture,
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse("berhasil mengambil profil", adminResponse))
}
