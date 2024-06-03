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
		Password:           newAdmin.Password,
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

	//mapping
	adminResponse := AdminResponse{
		FullName:       profile.FullName,
		Email:          profile.Email,
		Address:        profile.Address,
		NumberPhone:    profile.NumberPhone,
		ProfilePicture: profile.ProfilePicture,
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse("berhasil mengambil profil", adminResponse))
}

func (ah *AdminHandler) Delete(c echo.Context) error {
	adminID := middlewares.ExtractTokenUserId(c)
	if adminID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("Unauthorized", nil))
	}

	errDelete := ah.adminService.Delete(uint(adminID))
	if errDelete != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("hapus akun gagal: "+errDelete.Error(), nil))
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse("hapus akun berhasil", nil))
}

func (ah *AdminHandler) Update(c echo.Context) error {
	adminID := middlewares.ExtractTokenUserId(c)
	if adminID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("Unauthorized", nil))
	}

	updateReq := AdminRequest{}
	errBind := c.Bind(&updateReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Error binding data: "+errBind.Error(), nil))
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
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Update gagal: "+errUpdate.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Update gagal: "+errUpdate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse("Update berhasil", nil))
}
