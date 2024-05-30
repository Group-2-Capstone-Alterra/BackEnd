package service

import (
	"PetPalApp/app/middlewares"
	"PetPalApp/features/admin"
	"PetPalApp/utils/encrypts"
	"errors"
)

type AdminService struct {
	AdminModel admin.AdminModel
	hashService encrypts.HashInterface
}

func New(am admin.AdminModel, hash encrypts.HashInterface) admin.AdminService {
	return &AdminService{
		AdminModel: am,
		hashService: hash,
	}
}

func (as *AdminService) Register(admin admin.Core) error {
	if admin.FullName == "" || admin.Email == "" || admin.NumberPhone == "" ||  admin.Address == "" || admin.Password == "" || 
	admin.KetikUlangPassword == "" {
		return errors.New("[validation] Fullname/email/numberphone/address/password tidak boleh kosong")
	}

	if admin.Password != admin.KetikUlangPassword {
		return errors.New("[validation] password dan konfirmasi password tidak cocok")
	}

	var errHash error
	if admin.Password, errHash = as.hashService.HashPassword(admin.Password); errHash != nil {
		return errHash
	}

	if admin.KetikUlangPassword, errHash = as.hashService.HashPassword(admin.KetikUlangPassword); errHash != nil {
		return errHash
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

	token, errJWT := middlewares.CreateToken(int(data.ID))
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

func (as *AdminService) Update(adminid uint, updateData admin.Core) error {
	if updateData.FullName == "" && updateData.Email == "" && updateData.NumberPhone == "" && updateData.Address == "" && updateData.ProfilePicture == "" {
		return errors.New("[validation] Tidak ada data yang diupdate")
	}

	return as.AdminModel.Update(adminid, updateData)
}

