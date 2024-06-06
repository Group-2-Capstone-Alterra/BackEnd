package service

import (
	"PetPalApp/app/middlewares"
	"PetPalApp/features/admin"
	"PetPalApp/features/clinic"
	"PetPalApp/features/doctor"
	"PetPalApp/utils/encrypts"
	"PetPalApp/utils/helper"
	"errors"
	"fmt"
	"io"
	"log"
	"time"
)

type AdminService struct {
	AdminModel  admin.AdminModel
	DoctorModel doctor.DoctorModel
	hashService encrypts.HashInterface
	helper      helper.HelperInterface
}

func New(am admin.AdminModel, hash encrypts.HashInterface, DoctorModel doctor.DoctorModel, helper helper.HelperInterface) admin.AdminService {
	return &AdminService{
		AdminModel:  am,
		hashService: hash,
		DoctorModel: DoctorModel,
		helper:      helper,
	}
}

func (as *AdminService) Register(admin admin.Core) error {

	if admin.FullName == "" || admin.Email == "" || admin.Password == "" {
		return errors.New("[validation] Fullname/email/password tidak boleh kosong")
	}

	var errHash error
	if admin.Password, errHash = as.hashService.HashPassword(admin.Password); errHash != nil {
		return errHash
	}

	if admin.ProfilePicture == "" {
		admin.ProfilePicture = "https://air-bnb.s3.ap-southeast-2.amazonaws.com/profilepicture/default.jpg"
	}

	return as.AdminModel.Register(admin)
}

func (as *AdminService) Login(email string, password string) (data *admin.Core, token string, err error) {
	data, err = as.AdminModel.AdminByEmail(email)
	if err != nil {
		return nil, "", err
	}

	isLoginValid := as.hashService.CheckPasswordHash(data.Password, password)
	if !isLoginValid {
		return nil, "", errors.New("email atau password tidak sesuai")
	}

	token, errJWT := middlewares.CreateToken(int(data.ID), data.Role)
	if errJWT != nil {
		return nil, "", errJWT
	}
	return data, token, nil
}

func (as *AdminService) GetProfile(adminid uint) (*admin.Core, error) {
	profile, err := as.AdminModel.AdminById(adminid)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (as *AdminService) Delete(adminid uint) error {
	err := as.AdminModel.Delete(adminid)
	if err != nil {
		return err
	}
	return nil
}

func (as *AdminService) Update(adminid uint, updateData admin.Core, file io.Reader, handlerFilename string) error {
	if updateData.FullName == "" && updateData.Email == "" && updateData.NumberPhone == "" && 
	updateData.Address == "" && updateData.Coordinate == "" && updateData.Password == "" && file == nil {
		return errors.New("[validation] Tidak ada data yang diupdate")
	}

	if updateData.Password != "" {
		result, errHash := as.hashService.HashPassword(updateData.Password)
		if errHash != nil {
			return errHash
		}
		updateData.Password = result
	}

	if file != nil && handlerFilename != "" {
		timestamp := time.Now().Unix()
		fileName := fmt.Sprintf("%d_%s", timestamp, handlerFilename)
		photoFileName, errPhoto := as.helper.UploadAdminPicture(file, fileName)
		if errPhoto != nil {
			return errPhoto
		}
		updateData.ProfilePicture = photoFileName
	} else if updateData.ProfilePicture == "" {
		updateData.ProfilePicture = "https://air-bnb.s3.ap-southeast-2.amazonaws.com/profilepicture/default.jpg"
	}

	return as.AdminModel.Update(adminid, updateData)
}



func (as *AdminService) GetAllClinic(userid uint, offset uint, sortStr string) ([]clinic.Core, error) {
	log.Println("[Service]")
	// log.Println("[Service] sortStr", sortStr)
	// log.Println("[Service] userid", userid)
	allDoctor, errAllDoctor := as.DoctorModel.SelectAllDoctor()
	if errAllDoctor != nil {
		return nil, errAllDoctor
	}
	// log.Println("[Service - Admin] allDoctor", allDoctor)

	var allClinic []clinic.Core
	for _, v := range allDoctor {
		// log.Println("[Service] value", v)
		// log.Println("[Service] value", v.AdminID)
		adminDetail, errAdminDetail := as.AdminModel.AdminById(v.AdminID)
		if errAdminDetail != nil {
			return nil, errAdminDetail
		}
		doctorAvailDay, errDoctorAvailDay := as.DoctorModel.SelectAvailDayById(v.ID)
		if errDoctorAvailDay != nil {
			return nil, errDoctorAvailDay
		}

		serviceDoctor, errServiceDoctor := as.DoctorModel.SelectServiceById(v.ID)
		if errServiceDoctor != nil {
			return nil, errServiceDoctor
		}
		allClinic = append(allClinic, clinic.Core{
			ID:         adminDetail.ID,
			ClinicName: adminDetail.FullName,
			Open:       *doctorAvailDay,
			Service:    *serviceDoctor,
			Veterinary: v.FullName,
			Location:   adminDetail.Address,
			Coordinate: adminDetail.Coordinate,
		})
	}
	clinicSort := as.helper.SortClinicsByDistance(userid, allClinic)
	// log.Println("[Service - Admin] clinicSort", clinicSort)
	// log.Println("[Service - Admin] allClinic", allClinic)
	return clinicSort, nil
}
