package handler

import (
	"PetPalApp/app/middlewares"
	"PetPalApp/features/doctor"
	"PetPalApp/utils/responses"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type DoctorHandler struct {
    doctorService doctor.DoctorService
}

func New(ds doctor.DoctorService) *DoctorHandler {
    return &DoctorHandler{
        doctorService: ds,
    }
}

func (dh *DoctorHandler) AddDoctor(c echo.Context) error {
    adminID := middlewares.ExtractTokenUserId(c)
    if adminID == 0 {
        return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("Unauthorized", nil))
    }

    newDoctor := DoctorRequest{}
    errBind := c.Bind(&newDoctor)
    if errBind != nil {
        return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Error binding data: "+errBind.Error(), nil))
    }

    doctor := doctor.Core{
        FullName:       newDoctor.FullName,
        Email:          newDoctor.Email,
        Specialization: newDoctor.Specialization,
    }

    errAdd := dh.doctorService.AddDoctor(doctor)
    if errAdd != nil {
        if strings.Contains(errAdd.Error(), "validation") {
            return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Tambah dokter gagal: "+errAdd.Error(), nil))
        }
        return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Tambah dokter gagal: "+errAdd.Error(), nil))
    }

    return c.JSON(http.StatusCreated, responses.JSONWebResponse("Tambah dokter berhasil", nil))
}
