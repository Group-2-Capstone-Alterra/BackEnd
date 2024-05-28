package service

import (
	"PetPalApp/app/middlewares"
	"PetPalApp/features/user"
	"PetPalApp/utils/encrypts"
	"PetPalApp/utils/helper"
	"errors"
)

type userService struct {
	userData      user.DataInterface
	hashService   encrypts.HashInterface
	helperService helper.HelperInterface
}

func New(ud user.DataInterface, hash encrypts.HashInterface, helper helper.HelperInterface) user.ServiceInterface {
	return &userService{
		userData:      ud,
		hashService:   hash,
		helperService: helper,
	}
}

func (u *userService) Create(input user.Core) (string, error) {
	if input.FullName == "" || input.Email == "" || input.Password == "" {
		return "", errors.New("[validation] nama/email/password tidak boleh kosong")
	}
	if input.Password != input.RetypePassword {
		return "", errors.New("[validation] password dan ketik ulang password berbeda")
	}
	if input.Password != "" {
		result, errHash := u.hashService.HashPassword(input.Password)
		if errHash != nil {
			return "", errHash
		}
		input.Password = result
	}
	defaultPhoto := "https://air-bnb.s3.ap-southeast-2.amazonaws.com/fotoprofile/default-pp.jpg"
	input.ProfilePicture = defaultPhoto
	err := u.userData.Insert(input)
	if err != nil {
		return "", err
	}
	return input.ProfilePicture, nil
}

func (u *userService) Login(email string, password string) (data *user.Core, token string, err error) {
	data, err = u.userData.SelectByEmail(email)
	if err != nil {
		return nil, "", err
	}
	isLoginValid := u.hashService.CheckPasswordHash(data.Password, password)
	if !isLoginValid {
		return nil, "", errors.New("[validation] password tidak sesuai")
	}
	token, errJWT := middlewares.CreateToken(int(data.ID))
	if errJWT != nil {
		return nil, "", errJWT
	}
	return data, token, nil
}
