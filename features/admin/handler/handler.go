package handler

import (
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
		FullName: newAdmin.FullName,
		Email: newAdmin.Email,
		NumberPhone: newAdmin.NumberPhone,
		Address: newAdmin.Address,
		Password: newAdmin.Password,
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