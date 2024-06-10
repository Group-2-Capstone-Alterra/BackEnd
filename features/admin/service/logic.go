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
	}

	existingAdmin, err := as.AdminModel.AdminById(adminid)
	if err != nil {
		return err
	}

	if file == nil && updateData.ProfilePicture == "" {
		updateData.ProfilePicture = existingAdmin.ProfilePicture
	}

	return as.AdminModel.Update(adminid, updateData)
}

func (as *AdminService) GetAllClinic(userid uint, offset uint, sortStr string) ([]clinic.Core, error) {
	log.Println("[Service]")
	// log.Println("[Service] sortStr", sortStr)
	var allAdmin []admin.Core
	var errAllAdmin error
	if userid == 0 {
		allAdmin, errAllAdmin = as.AdminModel.SelectAllAdmin()
	} else {
		allAdmin, errAllAdmin = as.AdminModel.SelectAllAdminWithCoor()
	}
	if errAllAdmin != nil {
		return nil, errAllAdmin
	}

	var allClinic []clinic.Core
	for _, v := range allAdmin {
		dataAdmin, _ := as.AdminModel.AdminById(v.ID)
		doctorDetail, errDoctorDetail := as.DoctorModel.SelectByAdminId(v.ID)
		if errDoctorDetail != nil {
			return nil, errDoctorDetail
		}
		doctorAvailDay, errDoctorAvailDay := as.DoctorModel.SelectAvailDayById(doctorDetail.ID)
		if errDoctorAvailDay != nil {
			return nil, errDoctorAvailDay
		}

		serviceDoctor, errServiceDoctor := as.DoctorModel.SelectServiceById(doctorDetail.ID)
		if errServiceDoctor != nil {
			return nil, errServiceDoctor
		}
		allClinic = append(allClinic, clinic.Core{
			ID:                v.ID,
			ClinicName:        v.FullName,
			ClinicPicture:     dataAdmin.ProfilePicture,
			Open:              *doctorAvailDay,
			Service:           *serviceDoctor,
			Veterinary:        doctorDetail.FullName,
			VeterinaryPicture: doctorDetail.ProfilePicture,
			Location:          v.Address,
			Coordinate:        v.Coordinate,
		})
	}
	if userid == 0 {
		return allClinic, nil
	} else {
		clinicSort := as.helper.SortClinicsByDistance(userid, allClinic)
		// log.Println("[Service - Admin] allClinic", allClinic)
		return clinicSort, nil
	}
}

func (as *AdminService) GetClinic(id uint) (clinic.Core, error) {
	log.Println("[Service]")
	// log.Println("[Service] sortStr", sortStr)

	dataAdmin, _ := as.AdminModel.AdminById(id)

	doctorDetail, _ := as.DoctorModel.SelectByAdminId(dataAdmin.ID)

	doctorAvailDay, _ := as.DoctorModel.SelectAvailDayById(doctorDetail.ID)

	serviceDoctor, _ := as.DoctorModel.SelectServiceById(doctorDetail.ID)

	result := clinic.Core{
		ID:                dataAdmin.ID,
		ClinicName:        dataAdmin.FullName,
		ClinicPicture:     dataAdmin.ProfilePicture,
		Open:              *doctorAvailDay,
		Service:           *serviceDoctor,
		IDDoctor:          doctorDetail.ID,
		Veterinary:        doctorDetail.FullName,
		VeterinaryPicture: doctorDetail.ProfilePicture,
		About:             doctorDetail.About,
		Price:             doctorDetail.Price,
		Location:          dataAdmin.Address,
		Coordinate:        dataAdmin.Coordinate,
	}

	return result, nil
}
