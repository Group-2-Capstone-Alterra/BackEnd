package service

import (
	"PetPalApp/app/middlewares"
	"PetPalApp/features/user"
	"PetPalApp/utils/encrypts"
	"PetPalApp/utils/helper"
	"errors"
	"fmt"
	"io"
	"log"
	"time"
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

	if input.Password != "" {
		result, errHash := u.hashService.HashPassword(input.Password)
		if errHash != nil {
			return "", errHash
		}
		input.Password = result
	}
	defaultPhoto := "https://air-bnb.s3.ap-southeast-2.amazonaws.com/profilepicture/default.jpg"
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

func (u *userService) GetProfile(id uint) (data *user.Core, err error) {
	if id <= 0 {
		return nil, errors.New("[validation] id not valid")
	}
	return u.userData.SelectById(id)
}

func (u *userService) UpdateById(id uint, input user.Core, file io.Reader, handlerFilename string) (string, error) {
	if id <= 0 {
		return "", errors.New("id not valid")
	}

	if input.Password != "" {
		result, errHash := u.hashService.HashPassword(input.Password)
		if errHash != nil {
			return "", errHash
		}
		input.Password = result
	}
	log.Println("[Service - UpdateById]")

	log.Println("[Service - UpdateById] file ", file)
	log.Println("[Service - UpdateById] handlerFilename ", handlerFilename)
	if file != nil && handlerFilename != "" {
		timestamp := time.Now().Unix()
		fileName := fmt.Sprintf("%d_%s", timestamp, handlerFilename)
		log.Println("[Service - UpdateById] fileName ", fileName)
		photoFileName, errPhoto := u.helperService.UploadProfilePicture(file, fileName)
		if errPhoto != nil {
			return "", errPhoto
		}
		input.ProfilePicture = photoFileName
		log.Println("[Service - UpdateById] input.ProfilePicture ", input.ProfilePicture)
	}

	err := u.userData.PutById(id, input)
	if err != nil {
		return "", err
	}
	return input.ProfilePicture, nil
}

func (u *userService) Delete(id uint) error {
	if id <= 0 {
		return errors.New("id not valid")
	}
	return u.userData.Delete(id)
}
