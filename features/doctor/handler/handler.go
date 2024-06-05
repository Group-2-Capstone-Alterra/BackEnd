package handler

import (
	"PetPalApp/app/middlewares"
	"PetPalApp/features/doctor"
	"PetPalApp/utils/responses"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type DoctorHandler struct {
	doctorService doctor.DoctorService
	// availDaysDoctData
}

func New(ds doctor.DoctorService) *DoctorHandler {
	return &DoctorHandler{
		doctorService: ds,
	}
}

func (dh *DoctorHandler) AddDoctor(c echo.Context) error {
	adminID, _, _ := middlewares.ExtractTokenUserId(c)
	if adminID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("Unauthorized", nil))
	}
	log.Println("[Handler Doctor - AddDoctor] adminID ", adminID)

	req := AddDoctorRequest{}
	errBind := c.Bind(&req)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Error binding doctor data: "+errBind.Error(), nil))
	}
	log.Println("[Handler Doctor - AddDoctor] req ", req)
	inputCore := RequestToCore(req)
	inputCore.AdminID = uint(adminID)
	log.Println("[Handler Doctor - AddDoctor] inputCore.AdminID ", inputCore.AdminID)
	log.Println("[Handler Doctor - AddDoctor] inputCore ", inputCore)

	errAdd := dh.doctorService.AddDoctor(inputCore)
	if errAdd != nil {
		if strings.Contains(errAdd.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Tambah dokter gagal: "+errAdd.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Tambah dokter gagal: "+errAdd.Error(), nil))
	}

	return c.JSON(http.StatusCreated, responses.JSONWebResponse("Tambah dokter berhasil. Dokter ID: "+strconv.Itoa(int(inputCore.ID)), nil))
}

func (dh *DoctorHandler) ProfileDoctor(c echo.Context) error {
	log.Println("[handler]")

	idToken, _, _ := middlewares.ExtractTokenUserId(c) // extract id user from jwt token
	log.Println("idtoken:", idToken)

	doctorDetails, errDoctorDetails := dh.doctorService.GetDoctorByIdAdmin(uint(idToken))
	if errDoctorDetails != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error read data", nil))
	}

	availDaysDoct, errAvailDaysDoct := dh.doctorService.GetAvailDoctorByIdDoctor(doctorDetails.ID)
	if errAvailDaysDoct != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error read data", nil))
	}
	log.Println("[Handler Doctor - ProfileDoctor] availDaysDoct ", availDaysDoct)

	doctorResponse := GormToCore(*doctorDetails, *availDaysDoct)
	return c.JSON(http.StatusOK, responses.JSONWebResponse("success get detail doctor profile", doctorResponse))
}

func (dh *DoctorHandler) UpdateProfile(c echo.Context) error {
	adminID, _, _ := middlewares.ExtractTokenUserId(c)
	if adminID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("Unauthorized", nil))
	}
	log.Println("[Handler Doctor - AddDoctor] adminID ", adminID)

	req := AddDoctorRequest{}
	errBind := c.Bind(&req)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Error binding doctor data: "+errBind.Error(), nil))
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

	log.Println("[Handler Doctor - AddDoctor] req ", req)
	inputCore := RequestToCore(req)
	// inputCore.AdminID = uint(adminID)
	// log.Println("[Handler Doctor - AddDoctor] inputCore.AdminID ", inputCore.AdminID)
	log.Println("[Handler Doctor - AddDoctor] inputCore ", inputCore)

	var filename string
	if handler != nil {
		filename = handler.Filename
	}

	_, errUpdate := dh.doctorService.UpdateByIdAdmin(uint(adminID), inputCore, file, filename)
	if errUpdate != nil {
		// Handle error from userService.UpdateById
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error update data", err))
	}
	// Return success response
	return c.JSON(http.StatusOK, responses.JSONWebResponse("success update data", err))
}
