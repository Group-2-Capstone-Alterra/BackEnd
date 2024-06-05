package service

import (
	"PetPalApp/features/availdaydoctor"
	"PetPalApp/features/doctor"
	"PetPalApp/utils/helper"
	"errors"
	"fmt"
	"io"
	"log"
	"time"
)

type DoctorService struct {
	DoctorModel doctor.DoctorModel
	helper      helper.HelperInterface
}

func New(dm doctor.DoctorModel, helper helper.HelperInterface) doctor.DoctorService {
	return &DoctorService{
		DoctorModel: dm,
		helper:      helper,
	}
}

func (ds *DoctorService) AddDoctor(doctor doctor.Core, file io.Reader, handlerFilename string) (string, error) {
	//valisadminhavedoctor
	isAdminHaveDoct, _ := ds.DoctorModel.SelectByAdminId(doctor.AdminID)
	if isAdminHaveDoct.ID != 0 {
		return "", errors.New("Anda sudah mempunyai dokter !")
	} else {
		if doctor.FullName == "" {
			return "", errors.New("[validation] Fullname/specialization tidak boleh kosong")
		}

		timestamp := time.Now().Unix()
		fileName := fmt.Sprintf("%d_%s", timestamp, handlerFilename)
		photoFileName, errPhoto := ds.helper.UploadDoctorPicture(file, fileName)
		if errPhoto != nil {
			return "", errPhoto
		}
		doctor.ProfilePicture = photoFileName

		err := ds.DoctorModel.AddDoctor(doctor)
		if err != nil {
			return "", err
		}
		return doctor.ProfilePicture, nil
	}

}

func (ds *DoctorService) GetDoctorByIdAdmin(adminID uint) (*doctor.Core, error) {
	return ds.DoctorModel.SelectByAdminId(adminID)
}

func (ds *DoctorService) GetAvailDoctorByIdDoctor(doctorID uint) (*availdaydoctor.Core, error) {
	availDayCore, _ := ds.DoctorModel.SelectAvailDayById(doctorID)
	log.Println("[Serive - GetAvailDoctorByIdDoctor] availDayCore", availDayCore)

	return availDayCore, nil
}

func (ds *DoctorService) UpdateByIdAdmin(AdminId uint, input doctor.Core, file io.Reader, handlerFilename string) (string, error) {
	if AdminId <= 0 {
		return "", errors.New("id admin not valid")
	}

	if file != nil && handlerFilename != "" {
		timestamp := time.Now().Unix()
		fileName := fmt.Sprintf("%d_%s", timestamp, handlerFilename)
		photoFileName, errPhoto := ds.helper.UploadDoctorPicture(file, fileName)
		if errPhoto != nil {
			return "", errPhoto
		}
		input.ProfilePicture = photoFileName
	}

	err := ds.DoctorModel.PutByIdAdmin(AdminId, input)
	if err != nil {
		return "", err
	}
	return input.ProfilePicture, nil
}
